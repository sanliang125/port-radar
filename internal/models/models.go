package models

import "time"

// PortInfo 端口占用信息
type PortInfo struct {
	Port        int       `json:"port"`
	Protocol    string    `json:"protocol"`    // tcp/udp
	PID         int       `json:"pid"`
	ProcessName string    `json:"processName"` // 系统进程名
	LocalAddr   string    `json:"localAddr"`   // 本地地址
	State       string    `json:"state"`       // 连接状态
}

// AppMark 应用标记
type AppMark struct {
	ID          int       `json:"id"`
	Port        int       `json:"port"`
	Protocol    string    `json:"protocol"`
	AppName     string    `json:"appName"`     // 用户标记的应用名
	Description string    `json:"description"` // 描述
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ContainerInfo Docker容器信息
type ContainerInfo struct {
	ID      string   `json:"id"`      // 容器ID（短ID）
	Name    string   `json:"name"`    // 容器名
	Image   string   `json:"image"`   // 镜像名
	State   string   `json:"state"`   // 运行状态: running, exited, paused
	Ports   []PortMapping `json:"ports"` // 端口映射列表
}

// PortMapping 端口映射
type PortMapping struct {
	HostPort    int    `json:"hostPort"`    // 宿主机端口
	HostIP      string `json:"hostIP"`      // 宿主机IP
	ContainerPort int  `json:"containerPort"` // 容器端口
	Protocol    string `json:"protocol"`    // 协议 tcp/udp
}

// PortWithMark 端口信息带标记
type PortWithMark struct {
	PortInfo
	AppMark     *AppMark        `json:"appMark,omitempty"`
	Container   *ContainerInfo  `json:"container,omitempty"` // 关联的容器信息
}

// DockerStats Docker统计信息
type DockerStats struct {
	Available    bool `json:"available"`    // Docker是否可用
	TotalContainers int `json:"totalContainers"` // 总容器数
	RunningContainers int `json:"runningContainers"` // 运行中容器数
}
