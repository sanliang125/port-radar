.PHONY: build run clean

# 构建
build:
	go mod tidy
	go build -ldflags="-s -w" -o port-radar .

# 运行
run:
	go run main.go

# Windows 构建
build-win:
	go mod tidy
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o port-radar.exe .

# Linux 构建
build-linux:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o port-radar .

# 清理
clean:
	rm -f port-radar port-radar.exe
	rm -rf data/
