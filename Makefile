SHELL := /bin/bash
export GOPATH:=$(HOME)/.gopath:$(PWD)

build: 
	@[ -d bin ] || mkdir bin
	( go build -o bin/geo-zipdb src/geo-zipdb.go )

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/geozipdb src/geozipdb.go

install-deps:
	go get -u github.com/golang/lint/golint
	go get github.com/franela/goblin

format:
	( gofmt -s -w src/*.go src/geozipdb/*.go test/*/*.go )

lint:
	@( golint src/... && golint test/... && golint tools/... && golint examples )

qtest:
	@( cd test/unit && go test -cover )

test:
	@( go vet src/geozipdb/*.go && go vet src/*.go && cd test/unit && go test -cover )
	@( make lint )

watch:
	./watcher.js

run:
	( go run src/geozipdb.go )

start:
	( make build )
	./bin/geozipdb &

status:
	@( echo "implement a socket client that will request status..." )

ping:
	@( echo "implement a socket client that will request a ping..." )

edit:
	vi -O3 src/*/*.go test/unit/*.go src/*.go

.PHONY: format test qtest watch run
