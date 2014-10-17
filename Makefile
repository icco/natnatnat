.PHONY: run deploy

all: run

css:
	scss --trace -t compressed public/scss/style.scss public/css/style.css

run: css *.go
	gcloud preview app run . --project=natwelch-writing

deploy:
	gcloud preview app deploy . --project=natwelch-writing

update:
	goapp get -v -u ...
