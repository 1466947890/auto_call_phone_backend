# syntax=docker/dockerfile:1
# 第一阶段：编译阶段
FROM 192.168.31.103:80/hub/library/golang:1.26.3 AS builder

# 1. 设置环境变量
# CGO_ENABLED=0 保证静态编译，在 alpine 运行不依赖 glibc
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /go/src/app

# 2. 先拷贝依赖清单
COPY go.mod go.sum ./

# 3. 【加速 1】挂载 Go 模块缓存，避免重复下载
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# 4. 拷贝源代码
COPY . .

# 5. 【加速 2】挂载编译缓存 + 模块缓存进行构建
# -ldflags="-s -w" 用于进一步压缩二进制体积
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -ldflags="-s -w" -o /go/src/app/main main.go


# 第二阶段：运行阶段
# 使用轻量级 alpine 作为运行环境
FROM alpine:latest

# 安装 bash 和基础证书（如果你的 build.sh 用到 bash）
RUN apk add --no-cache bash ca-certificates

WORKDIR /app

# 从 builder 阶段拷贝编译好的二进制文件
COPY --from=builder /go/src/app/main /app/main

# 拷贝你的脚本目录（保持原有逻辑）
COPY --from=builder /go/src/app/script /app/script

# 给脚本执行权限
RUN chmod +x /app/script/build.sh

EXPOSE 8000

# 启动命令
CMD ["/bin/bash", "/app/script/build.sh"]