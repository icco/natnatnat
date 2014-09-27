GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w

.PHONY: run deploy

all: run

css:
	scss --trace -t compressed public/scss/style.scss public/css/style.css

run: css *.go
	goapp get -v github.com/icco/natnatnat
	gcloud preview app run . --project=natwelch-writing

deploy:
	gcloud preview app deploy . --project=natwelch-writing
