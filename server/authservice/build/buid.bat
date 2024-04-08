go build -o authservice server/authservice/cmd/main.go
docker build . -t authservice:latest
docker tag authservice:latest registry.heroku.com/go-server/authservice:2
docker push registry.heroku.com/go-server/authservice:2