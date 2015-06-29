all: local

PID = tmp/server.pid

GOPATH=/tmp/natnat

local: clean
	make restart
	fswatch -0 *.go */*.go sass/*.scss views/* | xargs -0 -n 1 -I {} make restart || make kill

kill:
	[ -f $(PID) ] && kill -9 `cat $(PID)` || true
	make clean

restart:
	make kill
	npm start
	make update
	goapp build # We do this for build checking
	goapp serve & echo $$! > $(PID)

clean:
	rm -f writing
	rm -f $(PID)
	rm -rf $(GOPATH)

deploy:
	git push
	goapp get -v github.com/icco/natnatnat
	gcloud preview app deploy app.yaml --project=natwelch-writing

deploy_alt:
	goapp deploy -application=natwelch-writing

update:
	-go get -u -v github.com/icco/natnatnat/...
	-goapp get -u -v ...

.PHONY: serve restart kill clean deploy deploy_alt
