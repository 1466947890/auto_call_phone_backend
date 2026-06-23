FROM golang:1.26.3

LABEL author="北稚"

# 设置环境变量，国内构建必带
ENV GOPROXY=https://goproxy.cn,direct

ENV GO111MODULE=on

WORKDIR /go/src/

# --- 核心改进：先处理依赖 ---

# 1. 只拷贝依赖描述文件
COPY go.mod go.sum ./

# 2. 下载依赖（这一层只要 mod 文件不变，就会被 Docker 永久缓存）
# 使用 download 比 tidy 在构建镜像时更标准
RUN go mod download

# --- 处理代码 ---

# 3. 此时再拷贝剩下的所有源代码
COPY . .

# 4. 执行编译
# 注意：确保 /go/src/app 目录存在，或者直接输出到当前目录
RUN mkdir -p /go/src/app && go build -o /go/src/app/main main.go

EXPOSE 8000

# 5. 启动脚本
# 确保你的 script/build.sh 内部调用的路径和编译出的 main 路径一致
RUN chmod +x /go/src/script/build.sh
CMD ["/bin/bash", "/go/src/script/build.sh"]