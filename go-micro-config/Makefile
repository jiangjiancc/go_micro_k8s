GOCMD=go
GOBUILD= $(GOCMD) build

all: build

build: clear build-apline docker-build docker-push

clean: clear

# linux cmd
clear:
	rm -rf main

build-apline:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD)  main.go

docker-build:
	docker build -t 846602483/go-micro-config:kubernetes .

docker-push:
	docker push 846602483/go-micro-config:kubernetes
