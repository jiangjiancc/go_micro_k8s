GOCMD=go
GOBUILD= $(GOCMD) build

all: build

build: clear build-apline docker-build docker-push

clean: clear

# linux cmd
clear:
	rm -rf web main

build-apline:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD)  -o web main.go

docker-build:
	docker build -t 846602483/go-micro-web:kubernetes .

docker-push:
	docker push 846602483/go-micro-web:kubernetes
