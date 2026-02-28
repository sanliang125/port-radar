# 构建阶段
FROM golang:1.21-alpine AS builder

# 安装编译依赖（SQLite 需要 CGO）
RUN apk add --no-cache gcc musl-dev

WORKDIR /build

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源码并编译
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o port-radar .

# 运行阶段
FROM alpine:3.19

# 安装运行时依赖（添加 iproute2 for ss 命令，net-tools for netstat）
RUN apk add --no-cache ca-certificates tzdata iproute2 net-tools docker-cli

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /build/port-radar .

# 创建数据目录
RUN mkdir -p /app/data

# 暴露端口
EXPOSE 8099

# 设置环境变量
ENV TZ=Asia/Shanghai

# 运行程序
CMD ["./port-radar"]
