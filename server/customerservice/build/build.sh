go build -o customerservice ../cmd/customerservice/main.go
docker build . -t customerservice:latest
docker tag customerservice:latest registry.heroku.com/go-server/customerservice
docker push registry.heroku.com/go-server/customerservice