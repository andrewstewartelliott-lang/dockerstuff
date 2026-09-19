go build .
docker build -t gin-info-api:latest .
docker tag gin-info-api:latest andrewstewartelliott/gin-info-api:latest
docker push andrewstewartelliott/gin-info-api:latest