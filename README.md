# natnatnat

[![Build Status](https://travis-ci.org/icco/natnatnat.svg?branch=master)](https://travis-ci.org/icco/natnatnat) [![GoDoc](https://godoc.org/github.com/icco/natnatnat?status.svg)](https://godoc.org/github.com/icco/natnatnat)

The next iteration in Nat's content management system. Previous versions include:

 * [tumble.io](http://github.com/icco/tumble)
 * [pseudoweb.net](http://github.com/icco/pseudoweb)

## Install

These directions are for OSX and assume you have [homebrew](http://brew.sh/) installed.

 1. Download [the Google Go App Engine SDK](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).
 2. Extract the SDK into a folder, and assign the location of it to the variable in `GOAPP` in the Makefile
 3. Run `brew bundle`
 3. Run `make update`
 5. Run `make` to run locally
 6. Run `make deploy` to deploy to Google App Engine.

## Design

This site is hosted at <http://writing.natwelch.com>. For now, it will use Google App Engine to auth me as an Admin and allow me to post new content. The eventual goal will be to switch to indie auth.

### Database

This uses Google Datastore to store structs. This makes it so you can only add columns and never remove them. We also write simpler versions of objects into the search indexes to that we can do full text search.

 - https://cloud.google.com/appengine/docs/go/datastore/reference
 - https://cloud.google.com/appengine/docs/go/datastore/queries
 - https://cloud.google.com/appengine/docs/go/search/

### Routes

TODO: Update.

 * `/` -
 * `/about` -
 * `/admin/?` -
 * `/aliases` -
 * `/archive(s?)` -
 * `/archive/work` -
 * `/clean/work` -
 * `/day/:year/:month/:day/?` -
 * `/edit/:id/?` -
 * `/edit/:id/?` -
 * `/feed.atom` - Atom feed of content
 * `/feed.rss` - RSS feed of content
 * `/images/:year/:month/:file` -
 * `/link/work` -
 * `/links` -
 * `/longform.json` -
 * `/longform/work` -
 * `/md` -
 * `/mention` - http://indiewebcamp.com/webmention
 * `/page/:page/?` -
 * `/post/:id/?` - View an individual post and its related webmentions
 * `/post/:id/json` -
 * `/post/:id/md` -
 * `/post/?` -
 * `/post/new/?` - Special admin only page to create a new post
 * `/posts.json` -
 * `/posts.md.json` -
 * `/search/work` -
 * `/search` -
 * `/settings` -
 * `/sitemap.xml` -
 * `/stats` -
 * `/summary.atom` -
 * `/summary.rss` -
 * `/tags/:id/?` - List of all posts with this tag
 * `/tags/?` -
 * `/work/queue` -

### Visual Design

I wrote posts about some of the process in these posts while I was at RC: [#164](https://writing.natwelch.com/post/164), [#134](https://writing.natwelch.com/post/134) and [#94](https://writing.natwelch.com/post/94).

## TODO

 * http://docs.travis-ci.com/user/languages/go/
 * https://developers.google.com/appengine/docs/go/gettingstarted/helloworld
 * http://dev.mikamai.com/post/68453619468/building-web-apps-with-traffic-the-go-micro-framework
 * https://developers.google.com/appengine/docs/domain
 * https://developers.google.com/appengine/docs/go/users/reference

### IRC Convos

 * http://indiewebcamp.com/irc/2014-10-09/line/1412887317947
 * http://indiewebcamp.com/irc/2014-10-09/line/1412888630804
 * http://indiewebcamp.com/irc/2014-10-09/line/1412888100522
 * http://indiewebcamp.com/irc/2014-10-09/line/1412887481931
