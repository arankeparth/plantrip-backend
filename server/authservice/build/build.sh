go build -o authservice ../cmd/authservice/main.go
docker build . -t authservice:latest
docker tag authservice:latest registry.heroku.com/go-server/authservice
docker push registry.heroku.com/go-server/authservice