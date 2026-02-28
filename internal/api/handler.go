package api

import (
	"database/sql"
	"net/http"
	"os/exec"
	"port-radar/internal/database"
	"port-radar/internal/models"
	"port-radar/internal/scanner"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置 API 路由
func SetupRoutes(r *gin.RouterGroup, db *sql.DB) {
	r.GET("/ports", listPorts(db))
	r.GET("/marks", listMarks(db))
	r.POST("/marks", saveMark(db))
	r.DELETE("/marks", deleteMark(db))
	r.GET("/check/:port", checkPort(db))
	r.POST("/kill/:pid", killProcess)

	// Docker 相关路由
	r.GET("/docker/stats", getDockerStats)
	r.GET("/docker/containers", listContainers)
	r.POST("/docker/:id/:action", dockerAction)
}

// listPorts 列出所有端口
func listPorts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ports, err := scanner.ScanPorts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 获取所有标记
		marks, err := database.GetAllAppMarks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 获取Docker容器信息
		containers, _ := scanner.GetContainers()

		// 合并端口信息和标记、容器信息
		result := make([]models.PortWithMark, 0, len(ports))
		for _, port := range ports {
			pwm := models.PortWithMark{
				PortInfo: port,
			}
			key := port.Protocol + ":" + strconv.Itoa(port.Port)
			if mark, exists := marks[key]; exists {
				pwm.AppMark = mark
			}

			// 关联Docker容器
			if containers != nil {
				container := scanner.GetContainerByPort(port.Port, port.Protocol, containers)
				pwm.Container = container
			}

			result = append(result, pwm)
		}

		c.JSON(http.StatusOK, result)
	}
}

// listMarks 列出所有标记
func listMarks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		marks, err := database.GetAllAppMarks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		result := make([]*models.AppMark, 0, len(marks))
		for _, mark := range marks {
			result = append(result, mark)
		}
		c.JSON(http.StatusOK, result)
	}
}

// saveMark 保存标记
func saveMark(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Port        int    `json:"port" binding:"required"`
			Protocol    string `json:"protocol"`
			AppName     string `json:"appName" binding:"required"`
			Description string `json:"description"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Protocol == "" {
			req.Protocol = "tcp"
		}

		err := database.SaveAppMark(db, req.Port, req.Protocol, req.AppName, req.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

// deleteMark 删除标记
func deleteMark(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Port     int    `json:"port" binding:"required"`
			Protocol string `json:"protocol"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Protocol == "" {
			req.Protocol = "tcp"
		}

		err := database.DeleteAppMark(db, req.Port, req.Protocol)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

// checkPort 检查端口是否被占用
func checkPort(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		portStr := c.Param("port")
		port, err := strconv.Atoi(portStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid port"})
			return
		}

		ports, err := scanner.ScanPorts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var result *models.PortWithMark
		for _, p := range ports {
			if p.Port == port {
				mark, _ := database.GetAppMark(db, p.Port, p.Protocol)
				result = &models.PortWithMark{
					PortInfo: p,
					AppMark:  mark,
				}
				break
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"occupied": result != nil,
			"info":     result,
		})
	}
}

// killProcess 终止进程
func killProcess(c *gin.Context) {
	pidStr := c.Param("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pid"})
		return
	}

	if pid == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot kill process with pid 0"})
		return
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid))
	default:
		cmd = exec.Command("kill", "-9", strconv.Itoa(pid))
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"output": string(output),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "process killed",
		"pid":     pid,
	})
}

// getDockerStats 获取Docker统计信息
func getDockerStats(c *gin.Context) {
	stats := scanner.GetDockerStats()
	c.JSON(http.StatusOK, stats)
}

// listContainers 列出所有容器
func listContainers(c *gin.Context) {
	containers, err := scanner.GetContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, containers)
}

// dockerAction 执行Docker容器操作
func dockerAction(c *gin.Context) {
	containerID := c.Param("id")
	action := c.Param("action")

	// 验证操作类型
	validActions := map[string]bool{
		"start":   true,
		"stop":    true,
		"restart": true,
		"remove":  true,
	}

	if !validActions[action] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid action: " + action})
		return
	}

	output, err := scanner.DockerAction(containerID, action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"output": output,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "container " + action + "ed",
		"containerId": containerID,
		"action":      action,
	})
}
