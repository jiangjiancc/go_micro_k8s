GOCMD=go
GOBUILD= $(GOCMD) build

all: build

build: clear build-apline docker-build docker-push

clean: clear

# linux cmd
clear:
	rm -rf client main

build-apline:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD)  -o client main.go

docker-build:
	docker build -t 846602483/go-micro-client:kubernetes .

docker-push:
	docker push 846602483/go-micro-client:kubernetes
