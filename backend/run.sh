cd /www/wwwroot/source.mangolang.com
docker pull 27.25.157.49:5000/public/appsource:prod
docker tag 27.25.157.49:5000/public/appsource:prod appsource:prod
docker-compose down && docker-compose up -d
#"insecure-registries": ["192.168.42.133:5000"]
# docker pull 27.25.157.49:5000/mango/mango:dev