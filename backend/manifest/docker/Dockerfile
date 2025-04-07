# ------------------------------
# 构建阶段
# ------------------------------

# ------------------------------
# 构建 libzsign.so
# ------------------------------
FROM alpine:3.18 AS zsign-builder

RUN apk add --no-cache \
    git \
    build-base \
    cmake \
    musl-dev \
    openssl-dev \
    minizip-dev \
    zlib-dev

COPY ./manifest/docker/zsign /zsign

WORKDIR /zsign/build/linux
RUN make clean && make CC=gcc CXX=g++


# ------------------------------
# 构建 go运行文件（使用Golang官方镜像指定1.23版本）
# ------------------------------
FROM golang:1.23-alpine AS go-builder

# 安装C依赖库和编译工具
RUN apk add --no-cache \
    musl-dev \
    minizip-dev \
    openssl-dev \
    zlib-dev \
    gcc \
    g++

# 复制 libzsign 头文件和库文件
COPY --from=zsign-builder /zsign/src /usr/local/include/zsign
COPY --from=zsign-builder /zsign/bin/libzsign.so /usr/local/lib/

# 配置动态库查找路径（确保运行时能找到）
RUN ln -s /usr/local/lib/libzsign.so /usr/lib/libzsign.so

WORKDIR /app
COPY . .

# 设置CGO编译参数
RUN CGO_ENABLED=1 \
    CC=gcc \
    CXX=g++ \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_CFLAGS="-I/usr/local/include/zsign" \
    CGO_LDFLAGS="-L/usr/local/lib -lzsign -lminizip" \
    go build -o bin/linux_amd64/main main.go


# ------------------------------
# 运行阶段
# ------------------------------
FROM alpine:3.18

ENV TZ=Asia/Shanghai
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

RUN apk add --no-cache \
    musl \
    openssl \
    libstdc++ \
    minizip \
    zlib

WORKDIR /app

COPY --from=go-builder /app/bin/linux_amd64/main /usr/bin/main
COPY --from=zsign-builder /zsign/bin/libzsign.so /usr/local/lib/

# 确保运行时动态库能被找到
ENV LD_LIBRARY_PATH=/usr/local/lib

RUN chmod +x /usr/bin/main

EXPOSE 8808

CMD ["main"]