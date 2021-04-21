# Burgers API

## Task - <a href='https://www.notion.so/Backend-Dev-Task-3eaa4227e5c144c582743d40b372cbf3'>link</a>

## A sample app is deployed <a href=''>here</a>

## Tech stack:
- Go (Golang)
- PostgreSQL
- Docker/docker-compose

## How to start the project:

### Docker way  (preferred):
```shell
docker-compose build # build the project
docker-compose up # start it
```

### Non docker way (annoying):
- Install PostgreSQL - <a href='https://www.postgresql.org/download/'>Download from here</a>, and follow some setup guide
- Install Go :)
- Change the params in the .env file, according to your local environment
- On linux:
```shell
# env stuff
export GO111MODULE=on
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

# run tests
go test ./domain/repository

# build and start
go mod download
go build
./burger-api
```

- On windows:
```shell
echo "Start WSL and go for the linux/docker version :) :) :)"
```

