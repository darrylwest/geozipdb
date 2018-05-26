SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)
IMAGE=ebay-local/geozipdb:latest

build: 
	@[ -d bin ] || mkdir bin
	( go build -o bin/geozipdb src/main.go )

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/geozipdb src/main.go

container:
	( make build-linux && cd docker && docker build -t $(IMAGE) . )

install-deps:
	go get -u github.com/golang/lint/golint
	go get github.com/franela/goblin
	go get github.com/darrylwest/cassava-logger/logger
	go get github.com/julienschmidt/httprouter

format:
	( gofmt -s -w src/*.go src/geozipdb/*.go test/*/*.go )

lint:
	@( golint src/... && golint test/... )

qtest:
	@( cd test/unit && go test -cover )

test:
	@( go vet src/geozipdb/*.go && go vet src/*.go && cd test/unit && go test -cover )
	@( make lint )

watch:
	./watcher.js

run:
	( go run src/main.go )

start:
	( make build )
	./bin/geozipdb &

start-container:
	docker run --name geozipdb --detach -p 4539:4539 darrylwest/geozipdb:latest

status:
	curl http://localhost:4540/v1/zipdb/status
	curl http://localhost:4541/v1/zipdb/status
	curl http://localhost:4542/v1/zipdb/status

edit:
	( gofmt -s -w src/*.go src/geozipdb/*.go test/*/*.go )
	vi -O2 src/*/*.go test/unit/*.go src/*.go

.PHONY: format test qtest watch run
