GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=Yam

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
	make bundle
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	make bundle
	./$(BINARY_NAME)
bundle:
	zgok build -e Yam -z static/ -o Yam
deps:
	dep ensure
	$(GOGET) github.com/srtkkou/zgok/...
