all: local

PID = tmp/server.pid

GOAPP=../go_appengine/goapp

local: clean assets build
	$(GOAPP) serve

build:
	$(GOAPP) build

assets:
	./node_modules/webpack/bin/webpack.js -p

clean:
	rm -f natnatnat
	rm -f $(PID)

deploy:
	git push
	$(GOAPP) deploy -application=natwelch-writing -version=$(shell date +%Y%m%d-%H%M)

update:
	-$(GOAPP) get -u -v ...

test: update build
	$(GOAPP) test

.PHONY: local assets clean deploy update build test
