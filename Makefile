all: local

PID = tmp/server.pid

GOAPP=../go_appengine/goapp

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
	$(GOAPP) build # We do this for build checking
	$(GOAPP) serve & echo $$! > $(PID)

clean:
	rm -f natnatnat
	rm -f $(PID)

deploy:
	git push
	$(GOAPP) get -v github.com/icco/natnatnat
	$(GOAPP) deploy -application=natwelch-writing

update:
	$(GOAPP) get -u -v github.com/icco/natnatnat/...
	$(GOAPP) get -u -v ...

.PHONY: serve restart kill clean deploy
