# dockerfile 编译
FROM golang:1.17-alpine AS builder

# 设置固定的项目路径
ENV APPPATH /go/src/fwds

# 为我们的镜像设置必要的环境变量 GOPROXY为go mod 代理
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai

# 移动到工作目录:
WORKDIR $APPPATH

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .
# 复制配置文件
COPY config ./config


# 将我们的代码编译成二进制可执行文件 fwds
RUN go build -o fwds .

# 创建一个小镜像
# Final stage
FROM debian:stretch-slim

# 设置固定的项目路径 (需要在设定一次)
ENV APPPATH /go/src/fwds

# 移动到工作目录:
WORKDIR /app

# 从builder镜像中把 /build 拷贝到当前目录
COPY --from=builder $APPPATH/fwds    /app
COPY --from=builder $APPPATH/config   /app/config
COPY --from=builder $APPPATH/storage   /app/storage

#创建日志文件夹
RUN mkdir -p /data/logs/

# 暴露端口
EXPOSE 8080

# 需要运行的命令
CMD ["/app/fwds", "-e", "env.yaml"]

# 命令示例:
# 1. build image: docker build -t fwds:v1 -f Dockerfile .
# 2. start: docker run --rm -it -p 8080:8080 fwds:v1
# 3. test: curl -i http://localhost:8080/health