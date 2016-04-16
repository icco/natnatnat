all: local

PID = tmp/server.pid

GOAPP=../go_appengine/goapp
DEVAPPSERVER=../go_appengine/dev_appserver.py
VERSION := $(shell date +%Y%m%d-%H%M)

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
	git tag -a $(VERSION) -m "Release version: $(VERSION)"
	git push && git push origin $(VERSION)
	$(GOAPP) deploy -application=natwelch-writing -version=$(VERSION)

update:
	rm -rf node_modules
	npm install
	../go_appengine/goapp get -d -v github.com/gorilla/feeds
	../go_appengine/goapp get -d -v github.com/gorilla/sessions
	../go_appengine/goapp get -d -v github.com/icco/xsrftoken
	../go_appengine/goapp get -d -v github.com/kennygrant/sanitize
	../go_appengine/goapp get -d -v github.com/pilu/traffic
	../go_appengine/goapp get -d -v github.com/russross/blackfriday
	../go_appengine/goapp get -d -v github.com/spf13/cast
	../go_appengine/goapp get -d -v github.com/spf13/hugo/parser
	../go_appengine/goapp get -d -v google.golang.org/appengine/search
	../go_appengine/goapp get -d -v google.golang.org/appengine/taskqueue
	../go_appengine/goapp get -d -v google.golang.org/appengine/user

test: update build
	$(GOAPP) test

new:
	./longform/new_post.sh

publish:
	./longform/publish.sh

drafts:
	@ls longform/drafts/* | grep '-'

.PHONY: local assets clean deploy update build test drafts publish new
