GOPATH=/tmp/natnatnat

.PHONY: run deploy

all: run

css:
	scss --trace -t compressed public/scss/style.scss public/css/style.css

run: clean css *.go
	goapp get -v github.com/icco/natnatnat
	gcloud preview app run app.yaml --project=natwelch-writing

deploy:
	git push
	gcloud preview app deploy app.yaml --project=natwelch-writing

deploy_alt:
	goapp deploy -application=natwelch-writing

update:
	cd $(GOPATH)/src/github.com/icco/natnatnat/; git pull
	goapp get -v -u ...

clean:
	rm -rf /tmp/natnatnat
	mkdir -p /tmp/natnatnat
