set GOARCH=amd64
set GOOS=linux
go build
set GOOS=windows
docker-compose build --no-cache && docker-compose up -d --force-recreate

