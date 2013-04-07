GOPATH :=`pwd`
SRC := ${GOPATH}/src

all:
	@echo ${GOPATH}
	go build msgstory

clean:
	go clean


install:


build:	msgstory.go
	go build ${SRC}/msgstory.go

test:
	go test