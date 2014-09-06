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

run: *.go
	gcloud preview app run . --project=natwelch-writing

deploy:
	gcloud preview app deploy . --project=natwelch-writing
