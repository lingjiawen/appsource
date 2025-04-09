.PHONY: all clean docker-prod zsign-builder go-builder

BUILD_DIR := ./build
ZSIGN_LIB := ${BUILD_DIR}/zsign/libzsign.so
GO_BIN := ${BUILD_DIR}/go/main

# 完整构建流程
all: zsign-builder go-builder docker-prod

# 清理所有构建产物
clean:
	rm -rf ${BUILD_DIR}

# 构建运行镜像（直接从本地缓存复制）
docker-prod:
	docker buildx build --platform linux/amd64 \
		-f ./manifest/docker/Dockerfile-runtime \
		-t 27.25.157.49:5000/public/appsource:prod \
		--push .

# ------------------------------
# 独立构建步骤（按需触发）
# ------------------------------

# 构建 zsign 库（仅在文件不存在时触发）
zsign-builder: ${ZSIGN_LIB}
${ZSIGN_LIB}:
	@echo "Building zsign..."
	@mkdir -p ${BUILD_DIR}/zsign
	docker buildx build --target export \
		--platform linux/amd64 \
		-o type=local,dest=${BUILD_DIR}/zsign/ \
		-f ./manifest/docker/Dockerfile-zsign .
	@tar -xzf ${BUILD_DIR}/zsign/zsign-output.tar.gz -C ${BUILD_DIR}/zsign
	@rm ${BUILD_DIR}/zsign/zsign-output.tar.gz

# 构建 Go 二进制（依赖 zsign，但仅在 Go 代码变化时触发）
go-builder: ${GO_BIN}
${GO_BIN}: ${ZSIGN_LIB}
	@echo "Building Go binary..."
	@mkdir -p ${BUILD_DIR}/go
	docker buildx build --target export \
		--platform linux/amd64 \
		-o type=local,dest=${BUILD_DIR}/go \
		-f ./manifest/docker/Dockerfile-go .