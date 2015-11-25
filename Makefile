all: local

PID = tmp/server.pid

GOAPP=../go_appengine/goapp

local: clean
	$(GOAPP) build # We do this for build checking
	$(GOAPP) serve

clean:
	rm -f natnatnat
	rm -f $(PID)

deploy:
	git push
	$(GOAPP) deploy -application=natwelch-writing

update:
	$(GOAPP) get -u -v ...

.PHONY: serve restart kill clean deploy
