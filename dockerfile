# 不要 syntax 开头那行，直接 FROM
FROM 192.168.31.103:80/hub/library/golang:1.26.3 AS builder

# 1. 设置 Go 代理（国内构建核心）
ENV GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /go/src/app

# 2. 【关键优化】先只拷贝 mod 和 sum 文件
# 只要你不增加或删除第三方包，这两行产生的缓存层永远不会失效
COPY go.mod go.sum ./

# 3. 【加速关键】在此处下载依赖
# 只要 mod 文件不变，这一步在第二次构建时直接跳过（显示 CACHED）
RUN go mod download

# 4. 此时再拷贝剩下的所有代码（main.go 等）
COPY . .

# 5. 执行编译
# 因为依赖已经下好了，即使改了代码，这里也只是编译你修改的那部分，通常 10-30 秒。
RUN mkdir -p /go/src/app && go build -ldflags="-s -w" -o /go/src/app/main main.go


# 第二阶段：运行阶段 (让镜像从 800MB 缩小到 20MB)
FROM 192.168.31.103:80/hub/library/alpine:latest
RUN apk add --no-cache bash ca-certificates
WORKDIR /app
COPY --from=builder /go/src/app/main /app/main
COPY --from=builder /go/src/app/script /app/script
RUN chmod +x /app/script/build.sh
EXPOSE 8000
CMD ["/bin/bash", "/app/script/build.sh"]