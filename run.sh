BASE_PATH=$(pwd)

cd cmd/http
env GOOS=linux GOARCH=amd64 go build
cd ../queue
env GOOS=linux GOARCH=amd64 go build
cd ../..
docker-compose down
docker-compose up -d
