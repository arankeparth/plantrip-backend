go build -o taskservice ../cmd/taskservice/main.go
docker build . -t taskservice:latest
docker tag taskservice:latest registry.heroku.com/go-server/taskservice
docker push registry.heroku.com/go-server/taskservice