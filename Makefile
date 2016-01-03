all: local

PID = tmp/server.pid

GOAPP=../go_appengine/goapp
DEVAPPSERVER=../go_appengine/dev_appserver.py

local: clean assets build
	$(DEVAPPSERVER) --log_level=debug --clear_datastore=true app.yaml

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
	rm -rf node_modules
	-npm install
	-$(GOAPP) get -d -u -v ...

test: update build
	$(GOAPP) test

new:
	./longform/new_post.sh

publish:
	./longform/publish.sh

drafts:
	@ls longform/drafts/* | grep '-'

.PHONY: local assets clean deploy update build test drafts publish new
