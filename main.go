package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"port-radar/internal/api"
	"port-radar/internal/database"
	"strings"
)

//go:embed web/*
var webFS embed.FS

func main() {
	// 确保数据目录存在
	dataDir := "./data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// 初始化数据库
	dbPath := filepath.Join(dataDir, "portmanager.db")
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// 创建 Gin 路由
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// API 路由
	apiGroup := r.Group("/api")
	api.SetupRoutes(apiGroup, db)

	// 静态文件服务
	webContent, err := fs.Sub(webFS, "web")
	if err != nil {
		log.Fatalf("Failed to sub web FS: %v", err)
	}

	// 首页
	r.GET("/", func(c *gin.Context) {
		data, _ := fs.ReadFile(webContent, "index.html")
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	// 静态资源 (style.css, app.js)
	r.GET("/static/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		data, err := fs.ReadFile(webContent, filename)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		contentType := "text/plain"
		switch {
		case strings.HasSuffix(filename, ".css"):
			contentType = "text/css; charset=utf-8"
		case strings.HasSuffix(filename, ".js"):
			contentType = "application/javascript; charset=utf-8"
		}
		c.Data(http.StatusOK, contentType, data)
	})

	log.Println("🚀 Port Radar starting on http://localhost:8099")
	if err := r.Run(":8099"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
