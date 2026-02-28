# 📡 端口雷达

一个轻量级的端口管理工具，支持端口扫描、应用标记、Docker容器识别与管理。专为 NAS、服务器运维场景设计。

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)
![License](https://img.shields.io/badge/License-MIT-green?style=flat)

[English](README.md)

## 💡 为什么需要端口雷达？

在 NAS 上部署多个 Docker 容器时，你是否经常遇到端口冲突的问题？新服务启动失败，却发现是端口被占用了？

端口雷达让你一目了然地看到系统中所有已占用的端口，轻松规避冲突，告别"端口已被占用"的烦恼！

## ✨ 功能特性

### 端口管理
- 🔍 **端口扫描** - 实时扫描系统已占用端口（TCP/UDP）
- 🏷️ **应用标记** - 为端口添加自定义名称和描述，方便识别
- 📋 **快速筛选** - 按协议、标记状态、Docker容器筛选
- 🔎 **搜索功能** - 支持端口号、进程名、应用名搜索

### Docker 集成
- 🐳 **容器识别** - 自动关联端口与Docker容器
- 📊 **容器统计** - 显示运行中容器数量
- ⚡ **容器管理** - 支持启动、停止、重启、删除容器
- 🔗 **端口映射** - 清晰展示容器端口映射关系

### 用户体验
- 🌐 **国际化** - 支持中文/英文切换
- 📱 **响应式设计** - 适配桌面和移动设备
- 🎨 **现代UI** - 暗色主题，清爽美观
- ⚡ **轻量高效** - 单二进制文件，无额外依赖

## 📸 截图

![端口雷达主界面](docs/img/main.png)

## 🚀 快速开始

### 方式一：Docker 部署（推荐）

```bash
# 使用 docker-compose
git clone https://github.com/sanliang125/port-radar.git
cd port-radar
docker-compose up -d

# 或直接运行
docker run -d \
  --name port-radar \
  --network host \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -v ./data:/app/data \
  -e TZ=Asia/Shanghai \
  sanliang125/port-radar:latest
```

访问 `http://localhost:8099` 即可使用。

### 方式二：二进制运行

```bash
# 下载对应平台的二进制文件
# Linux / macOS
./port-radar

# Windows
port-radar.exe
```

### 方式三：源码编译

```bash
git clone https://github.com/sanliang125/port-radar.git
cd port-radar
go build -o port-radar .
./port-radar
```

## 📖 使用说明

### 端口列表

| 列名 | 说明 |
|------|------|
| Port | 端口号 |
| Protocol | 协议类型（TCP/UDP） |
| Process | 进程名称（Docker容器显示容器名） |
| PID | 进程ID（Docker端口显示🐳图标） |
| Local Address | 本地监听地址 |
| App Mark | 应用标记名称 |
| Actions | 操作按钮 |

### 操作按钮

- **Mark/Edit** - 添加或编辑应用标记
- **Unmark** - 移除应用标记
- **Kill** - 终止进程（非Docker端口）
- **Stop/Start** - 停止/启动容器（Docker端口）
- **Restart** - 重启容器（Docker端口）
- **Remove** - 删除容器（已停止的容器）

### 应用模板

添加标记时，可快速选择常用应用模板：

| 应用 | 默认端口 |
|------|---------|
| MySQL | 3306 |
| Redis | 6379 |
| PostgreSQL | 5432 |
| MongoDB | 27017 |
| Nginx | 80 |
| Docker | 2375 |
| SSH | 22 |
| ... | ... |

## ⚙️ 配置说明

### 环境变量

| 变量 | 默认值 | 说明 |
|------|--------|------|
| `PORT` | `8099` | Web服务端口 |
| `TZ` | `Asia/Shanghai` | 时区设置 |

### 数据存储

- 数据库文件：`./data/portmanager.db`（SQLite）
- 自动创建，无需手动配置

### Docker 部署要求

要使用 Docker 容器管理功能，需要：

1. 挂载 Docker Socket：`-v /var/run/docker.sock:/var/run/docker.sock:ro`
2. 使用 host 网络模式：`--network host`

## 🛠️ 技术栈

- **后端**: Go 1.21+, Gin, SQLite
- **前端**: 原生 JavaScript, CSS3
- **部署**: Docker, Alpine Linux

## 📁 项目结构

```
port-radar/
├── main.go                 # 程序入口
├── internal/
│   ├── api/
│   │   └── handler.go      # API 处理器
│   ├── database/
│   │   └── database.go     # 数据库操作
│   ├── models/
│   │   └── models.go       # 数据模型
│   └── scanner/
│       ├── scanner.go      # 端口扫描器
│       └── docker.go       # Docker 容器扫描
├── web/
│   ├── index.html          # 主页面
│   ├── app.js              # 前端逻辑
│   └── style.css           # 样式文件
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## 🔌 API 接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/ports` | 获取端口列表 |
| GET | `/api/marks` | 获取所有标记 |
| POST | `/api/marks` | 保存标记 |
| DELETE | `/api/marks` | 删除标记 |
| GET | `/api/check/:port` | 检查端口状态 |
| POST | `/api/kill/:pid` | 终止进程 |
| GET | `/api/docker/stats` | Docker 统计信息 |
| GET | `/api/docker/containers` | 容器列表 |
| POST | `/api/docker/:id/:action` | 容器操作 |

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 License

[MIT License](LICENSE)

## 🙏 致谢

- 感谢 [Gin](https://github.com/gin-gonic/gin) Web Framework
- 图标来自 [Emoji](https://emojipedia.org/)
