GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w

.PHONY: run deploy

all: natnatnat

clean:
	rm natnatnat

natnatnat: *.go
	go build

run: natnatnat
	./natnatnat

deploy:
	gcloud preview app deploy . --project=natwelch-writing
