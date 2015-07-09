# natnatnat

[![Build Status](https://travis-ci.org/icco/natnatnat.svg?branch=master)](https://travis-ci.org/icco/natnatnat)

The next iteration in Nat's content management system. Previous versions include:

 * [tumble.io](http://github.com/icco/tumble)
 * [pseudoweb.net](http://github.com/icco/pseudoweb)

Docs: [godoc.org/github.com/icco/natnatnat](https://godoc.org/github.com/icco/natnatnat)

## Install

 0. Install Go
 1. `goapp get -v -u ...`
 2. `curl https://sdk.cloud.google.com | bash`
 3. `gcloud components update app`
 4. `make`

## Design

This site is hosted at <http://writing.natwelch.com>. For now, it will use Google App Engine to auth me as an Admin and allow me to post new content. The eventual goal will be to switch to indie auth.

### Database

So this system focuses around one gigantic database table for storing all entries.

 > Table Name: entries
 >  - id: unique id (integer...)
 >  - title: optional string
 >  - content: required string, markdown
 >  - datetime: required datetime
 >  - created: required datetime
 >  - modified: required datetime
 >  - tags: optional comma seperated list of tags
 >  - meta: json hash of extra data

This is out of date... TODO: Copy down all tables.

### Routes

TODO: Update.

 * `/` - Welcome page. Contains list of five most recent posts and an about page.
 * `/mention` - http://indiewebcamp.com/webmention
 * `/post/new` - Special admin only page to create a new post
 * `/post/:id` - View an individual post and its related webmentions
 * `/feed.atom` - Atom feed of content
 * `/tags/:tag` - List of all posts with this tag

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

## Git DB Idea

 * Creation: git log --diff-filter=A --follow --format=%aD -- main.go
 * Modified: git log --format=%aD -1 main.go
 * ID: filename
 * Tags: Hashtags

## Draft Post Idea

TODO: fill out
