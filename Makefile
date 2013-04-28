export GOPATH=/home/travis/gopath/jmadan/go-msgstory 

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w

TARG=go-msgstory

GOFILES=\
	main.go\
	register/register.go\
	authenticate/authenticate.go\
	circle/circle.go\
	user/user.go\
	message/message.go\


format:
	${GOFMT} -w ${GOFILES}

all:
	$(GOINSTALL)

