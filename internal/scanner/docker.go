package scanner

import (
	"bufio"
	"bytes"
	"os/exec"
	"port-radar/internal/models"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// IsDockerAvailable 检查Docker是否可用
func IsDockerAvailable() bool {
	cmd := exec.Command("docker", "version")
	err := cmd.Run()
	return err == nil
}

// GetDockerStats 获取Docker统计信息
func GetDockerStats() models.DockerStats {
	stats := models.DockerStats{
		Available: false,
	}

	if !IsDockerAvailable() {
		return stats
	}

	stats.Available = true

	// 获取容器总数
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.ID}}")
	output, err := cmd.Output()
	if err == nil {
		stats.TotalContainers = len(strings.Split(strings.TrimSpace(string(output)), "\n"))
		if strings.TrimSpace(string(output)) == "" {
			stats.TotalContainers = 0
		}
	}

	// 获取运行中容器数
	cmd = exec.Command("docker", "ps", "--format", "{{.ID}}")
	output, err = cmd.Output()
	if err == nil {
		stats.RunningContainers = len(strings.Split(strings.TrimSpace(string(output)), "\n"))
		if strings.TrimSpace(string(output)) == "" {
			stats.RunningContainers = 0
		}
	}

	return stats
}

// GetContainers 获取所有Docker容器信息
func GetContainers() ([]models.ContainerInfo, error) {
	if !IsDockerAvailable() {
		return nil, nil
	}

	// 使用 docker ps -a 获取所有容器
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.ID}}|{{.Names}}|{{.Image}}|{{.State}}|{{.Ports}}")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Windows 可能需要 GBK 转换
	if runtime.GOOS == "windows" {
		output = gbkToUTF8(output)
	}

	var containers []models.ContainerInfo

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) < 4 {
			continue
		}

		container := models.ContainerInfo{
			ID:    parts[0],
			Name:  parts[1],
			Image: parts[2],
			State: parts[3],
			Ports: []models.PortMapping{},
		}

		// 解析端口映射
		if len(parts) >= 5 && parts[4] != "" {
			container.Ports = parsePortMappings(parts[4])
		}

		containers = append(containers, container)
	}

	return containers, nil
}

// parsePortMappings 解析端口映射字符串
// 格式示例: 0.0.0.0:8080->80/tcp, 0.0.0.0:3306->3306/tcp, [::]:443->443/tcp
func parsePortMappings(portsStr string) []models.PortMapping {
	var mappings []models.PortMapping

	// 分割多个端口映射
	portParts := strings.Split(portsStr, ", ")
	for _, part := range portParts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		mapping := parseSinglePortMapping(part)
		if mapping != nil {
			mappings = append(mappings, *mapping)
		}
	}

	return mappings
}

// parseSinglePortMapping 解析单个端口映射
func parseSinglePortMapping(s string) *models.PortMapping {
	// 匹配格式: [IP]:HostPort->ContainerPort/Protocol
	// 示例: 0.0.0.0:8080->80/tcp, [::]:443->443/tcp, 0.0.0.0:3306->3306/tcp

	// 正则匹配
	// IPv4: 0.0.0.0:8080->80/tcp
	// IPv6: [::]:443->443/tcp
	re := regexp.MustCompile(`^(?:\[([^\]]+)\]|([^:]+)):?(\d+)->(\d+)/(tcp|udp)$`)

	matches := re.FindStringSubmatch(s)
	if matches == nil {
		return nil
	}

	hostIP := matches[1]
	if hostIP == "" {
		hostIP = matches[2]
	}
	hostPort, _ := strconv.Atoi(matches[3])
	containerPort, _ := strconv.Atoi(matches[4])
	protocol := matches[5]

	return &models.PortMapping{
		HostIP:        hostIP,
		HostPort:      hostPort,
		ContainerPort: containerPort,
		Protocol:      protocol,
	}
}

// GetContainerByPort 根据端口查找容器
func GetContainerByPort(port int, protocol string, containers []models.ContainerInfo) *models.ContainerInfo {
	for _, c := range containers {
		for _, p := range c.Ports {
			if p.HostPort == port && (p.Protocol == protocol || protocol == "") {
				return &c
			}
		}
	}
	return nil
}

// DockerAction 执行Docker操作
func DockerAction(containerID, action string) (string, error) {
	var cmd *exec.Cmd

	switch action {
	case "stop":
		cmd = exec.Command("docker", "stop", containerID)
	case "start":
		cmd = exec.Command("docker", "start", containerID)
	case "restart":
		cmd = exec.Command("docker", "restart", containerID)
	case "remove":
		cmd = exec.Command("docker", "rm", "-f", containerID)
	default:
		return "", nil
	}

	output, err := cmd.CombinedOutput()
	return string(output), err
}
