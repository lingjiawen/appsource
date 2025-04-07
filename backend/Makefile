.PHONY: docker-prod all

# 使用 Docker Buildx 构建并推送镜像
docker-prod:
	docker buildx build --platform linux/amd64 --no-cache -f ./manifest/docker/Dockerfile -t 27.25.157.49:5000/public/appsource:prod --push .

# 一键构建全部
all: docker-prod
