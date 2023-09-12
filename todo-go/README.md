# ToDo application backend writen in golang

## generate swagger documentation

all API docks are published with swagger on address <http://localhost:8080/swagger/index.html>

Prepare api docs

``` bash
go install github.com/swaggo/swag/cmd/swag@latest 
swag init -g cmd/main/main.go -o ./docs
```

## Build application

### Build on local machine

Run following command in terminal

``` bash
go build cmd/main/main.go
```

### Build in Docker

``` bash
docker build -f docker/Dockerfile . 
```

### Build in docker-compose

``` bash
docker-compose build -d
```

## Requirements

* Postgresql database
* environment variables:

|Name|Required|DefaultValue|Description|
|---|---|---|---|
|PORT|no|8080|Listen port for WEB API

## Run Application

``` bash
./todo-go 
```
