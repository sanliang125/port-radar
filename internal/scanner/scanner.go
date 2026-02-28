package scanner

import (
	"bufio"
	"bytes"
	"os/exec"
	"port-radar/internal/models"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// ScanPorts 扫描端口占用
func ScanPorts() ([]models.PortInfo, error) {
	switch runtime.GOOS {
	case "linux", "darwin":
		return scanPortsUnix()
	case "windows":
		return scanPortsWindows()
	default:
		return scanPortsUnix()
	}
}

// scanPortsUnix Linux/Mac 系统使用 ss 或 netstat
func scanPortsUnix() ([]models.PortInfo, error) {
	var ports []models.PortInfo

	// 尝试使用 ss 命令（更快）
	cmd := exec.Command("ss", "-tulpn")
	output, err := cmd.Output()
	if err != nil {
		// 回退到 netstat
		cmd = exec.Command("netstat", "-tulpn")
		output, err = cmd.Output()
		if err != nil {
			return nil, err
		}
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "LISTEN") {
			port := parseSSLine(line)
			if port != nil {
				ports = append(ports, *port)
			}
		}
	}

	return ports, nil
}

// scanPortsWindows Windows 系统使用 netstat
func scanPortsWindows() ([]models.PortInfo, error) {
	var ports []models.PortInfo

	cmd := exec.Command("netstat", "-ano")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Windows 输出可能是 GBK 编码，转换为 UTF-8
	output = gbkToUTF8(output)

	// 使用 map 去重
	portMap := make(map[string]models.PortInfo)

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "LISTENING") {
			port := parseNetstatLine(line)
			if port != nil {
				key := port.Protocol + ":" + strconv.Itoa(port.Port)
				if _, exists := portMap[key]; !exists {
					portMap[key] = *port
				}
			}
		}
	}

	for _, p := range portMap {
		ports = append(ports, p)
	}

	return ports, nil
}

// parseSSLine 解析 ss 命令输出行
func parseSSLine(line string) *models.PortInfo {
	fields := strings.Fields(line)
	if len(fields) < 6 {
		return nil
	}

	// 解析本地地址
	localAddr := fields[3]
	parts := strings.Split(localAddr, ":")
	if len(parts) < 2 {
		return nil
	}

	port, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return nil
	}

	// 解析协议
	protocol := "tcp"
	if strings.Contains(fields[0], "udp") {
		protocol = "udp"
	}

	// 解析 PID 和进程名
	pid := 0
	processName := ""
	if len(fields) >= 7 {
		processInfo := fields[len(fields)-1]
		if strings.Contains(processInfo, ",") {
			pInfoParts := strings.Split(processInfo, ",")
			pid, _ = strconv.Atoi(pInfoParts[1])
			processName = strings.TrimPrefix(pInfoParts[0], "users:((")
			processName = strings.Split(processName, ",")[0]
		}
	}

	return &models.PortInfo{
		Port:        port,
		Protocol:    protocol,
		PID:         pid,
		ProcessName: processName,
		LocalAddr:   localAddr,
		State:       "LISTEN",
	}
}

// parseNetstatLine 解析 Windows netstat 输出行
func parseNetstatLine(line string) *models.PortInfo {
	fields := strings.Fields(line)
	if len(fields) < 5 {
		return nil
	}

	// 解析协议
	protocol := strings.ToLower(fields[0])
	if protocol == "tcp" {
		protocol = "tcp"
	} else if protocol == "udp" {
		protocol = "udp"
	}

	// 解析本地地址
	localAddr := fields[1]
	parts := strings.Split(localAddr, ":")
	if len(parts) < 2 {
		return nil
	}

	port, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return nil
	}

	// 解析 PID
	pid, _ := strconv.Atoi(fields[len(fields)-1])

	// 获取进程名（Windows 需要额外查询）
	processName := getWindowsProcessName(pid)

	return &models.PortInfo{
		Port:        port,
		Protocol:    protocol,
		PID:         pid,
		ProcessName: processName,
		LocalAddr:   localAddr,
		State:       "LISTENING",
	}
}

// getWindowsProcessName 获取 Windows 进程名
func getWindowsProcessName(pid int) string {
	if pid == 0 {
		return ""
	}
	cmd := exec.Command("tasklist", "/FI", "PID eq "+strconv.Itoa(pid), "/NH")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	// Windows 输出可能是 GBK 编码，转换为 UTF-8
	output = gbkToUTF8(output)
	line := strings.TrimSpace(string(output))
	if line == "" || strings.Contains(line, "INFO:") {
		return ""
	}
	fields := strings.Fields(line)
	if len(fields) > 0 {
		return fields[0]
	}
	return ""
}

// gbkToUTF8 将 GBK 编码转换为 UTF-8
func gbkToUTF8(data []byte) []byte {
	decoder := simplifiedchinese.GBK.NewDecoder()
	decoded, err := decoder.Bytes(data)
	if err != nil {
		// 如果转换失败，返回原始数据
		return data
	}
	return decoded
}


