GOPATH=/tmp/natnatnat

.PHONY: run deploy

all: run

css:
	scss --trace -t compressed public/scss/style.scss public/css/style.css

run: clean css *.go
	goapp get -v github.com/icco/natnatnat
	gcloud preview app run . --project=natwelch-writing

deploy:
	gcloud preview app deploy . --project=natwelch-writing

clean:
	rm -rf /tmp/natnatnat
	mkdir -p /tmp/natnatnat
