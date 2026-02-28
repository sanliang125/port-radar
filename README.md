# 📡 Port Radar

A lightweight port management tool with port scanning, application labeling, and Docker container integration. Designed for NAS and server administration scenarios.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)
![License](https://img.shields.io/badge/License-MIT-green?style=flat)

[中文文档](README_CN.md)

## 💡 Why Port Radar?

Running multiple Docker containers on your NAS? Tired of port conflicts causing deployment failures?

Port Radar helps you visualize all active ports at a glance, making it easy to identify available ports before deploying new services. Say goodbye to "port already in use" errors!

## ✨ Features

### Port Management
- 🔍 **Port Scanning** - Real-time scanning of occupied ports (TCP/UDP)
- 🏷️ **Application Labeling** - Add custom names and descriptions to ports for easy identification
- 📋 **Quick Filtering** - Filter by protocol, label status, or Docker containers
- 🔎 **Search** - Search by port number, process name, or application name

### Docker Integration
- 🐳 **Container Recognition** - Automatically associate ports with Docker containers
- 📊 **Container Statistics** - Display running container count
- ⚡ **Container Management** - Start, stop, restart, and remove containers
- 🔗 **Port Mapping** - Clearly display container port mappings

### User Experience
- 🌐 **Internationalization** - Support for English/Chinese
- 📱 **Responsive Design** - Adapts to desktop and mobile devices
- 🎨 **Modern UI** - Dark theme with clean aesthetics
- ⚡ **Lightweight** - Single binary with no external dependencies

## 📸 Screenshots

![Port Radar Main Interface](docs/img/main.png)

## 🚀 Quick Start

### Option 1: Docker Deployment (Recommended)

```bash
# Using docker-compose
git clone https://github.com/sanliang125/port-radar.git
cd port-radar
docker-compose up -d

# Or run directly
docker run -d \
  --name port-radar \
  --network host \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -v ./data:/app/data \
  -e TZ=Asia/Shanghai \
  sanliang125/port-radar:latest
```

Visit `http://localhost:8099` to access the application.

### Option 2: Binary Execution

```bash
# Download the binary for your platform
# Linux / macOS
./port-radar

# Windows
port-radar.exe
```

### Option 3: Build from Source

```bash
git clone https://github.com/sanliang125/port-radar.git
cd port-radar
go build -o port-radar .
./port-radar
```

## 📖 User Guide

### Port List Columns

| Column | Description |
|--------|-------------|
| Port | Port number |
| Protocol | Protocol type (TCP/UDP) |
| Process | Process name (Docker containers show container name) |
| PID | Process ID (Docker ports show 🐳 icon) |
| Local Address | Local listening address |
| App Mark | Application label name |
| Actions | Action buttons |

### Action Buttons

- **Mark/Edit** - Add or edit application label
- **Unmark** - Remove application label
- **Kill** - Terminate process (non-Docker ports)
- **Stop/Start** - Stop/start container (Docker ports)
- **Restart** - Restart container (Docker ports)
- **Remove** - Remove container (stopped containers)

### Application Templates

When adding a label, you can quickly select from common application templates:

| Application | Default Port |
|-------------|--------------|
| MySQL | 3306 |
| Redis | 6379 |
| PostgreSQL | 5432 |
| MongoDB | 27017 |
| Nginx | 80 |
| Docker | 2375 |
| SSH | 22 |
| ... | ... |

## ⚙️ Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8099` | Web service port |
| `TZ` | `Asia/Shanghai` | Timezone setting |

### Data Storage

- Database file: `./data/portmanager.db` (SQLite)
- Automatically created, no manual configuration needed

### Docker Deployment Requirements

To use Docker container management features:

1. Mount Docker Socket: `-v /var/run/docker.sock:/var/run/docker.sock:ro`
2. Use host network mode: `--network host`

## 🛠️ Tech Stack

- **Backend**: Go 1.21+, Gin, SQLite
- **Frontend**: Vanilla JavaScript, CSS3
- **Deployment**: Docker, Alpine Linux

## 📁 Project Structure

```
port-radar/
├── main.go                 # Application entry point
├── internal/
│   ├── api/
│   │   └── handler.go      # API handlers
│   ├── database/
│   │   └── database.go     # Database operations
│   ├── models/
│   │   └── models.go       # Data models
│   └── scanner/
│       ├── scanner.go      # Port scanner
│       └── docker.go       # Docker container scanner
├── web/
│   ├── index.html          # Main page
│   ├── app.js              # Frontend logic
│   └── style.css           # Styles
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## 🔌 API Reference

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/ports` | Get port list |
| GET | `/api/marks` | Get all labels |
| POST | `/api/marks` | Save label |
| DELETE | `/api/marks` | Delete label |
| GET | `/api/check/:port` | Check port status |
| POST | `/api/kill/:pid` | Terminate process |
| GET | `/api/docker/stats` | Docker statistics |
| GET | `/api/docker/containers` | Container list |
| POST | `/api/docker/:id/:action` | Container operations |

## 🤝 Contributing

Issues and Pull Requests are welcome!

## 📄 License

[MIT License](LICENSE)

## 🙏 Acknowledgments

- [Gin](https://github.com/gin-gonic/gin) Web Framework
- Icons from [Emoji](https://emojipedia.org/)
