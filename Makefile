all: local

PID=tmp/server.pid

local: clean
	make restart
	fswatch -0 *.go public/scss/*.scss | xargs -0 -n 1 -I {} make restart || make kill

kill:
	[ -f $(PID) ] && kill -9 `cat $(PID)` || true
	make clean

restart:
	make kill
	make css
	goapp build # We do this for build checking
	goapp serve & echo $$! > $(PID)

clean:
	rm -f writing
	rm -f $(PID)

css:
	scss --trace -t compressed public/scss/style.scss public/css/style.css

update:
	#cd $(GOPATH)/src/github.com/icco/natnatnat/ && git pull
	goapp get -u -v ...

deploy:
	git push
	goapp get -v github.com/icco/natnatnat
	gcloud preview app deploy app.yaml --project=natwelch-writing

deploy_alt:
	goapp deploy -application=natwelch-writing

.PHONY: serve restart kill clean deploy deploy_alt
