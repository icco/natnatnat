# natnatnat

The next iteration in Nat's content management system. Previous versions include:

 * [tumble.io](http://github.com/icco/tumble)
 * [pseudoweb.net](http://github.com/icco/pseudoweb)

## Design

This site will be hosted at <http://writing.natwelch.com>. For now, it will use Google App Engine to auth me as an Admin and allow me to post new content. The eventual goal will be to switch to indie auth.

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

The current thought is that tags will filter out what is shown on each domain. So when you visit [tumble.io](http://tumble.io), you'll actually just be getting a view of the RSS feed for the links tag. [pseudoweb.net](http://pseudoweb.net) will be a view of the longform tag.

### Routes

 * `/` - Welcome page. Contains list of five most recent posts and an about page.
 * `/mention` - http://indiewebcamp.com/webmention
 * `/post/new` - Special admin only page to create a new post
 * `/post/:id` - View an individual post and its related webmentions
 * `/feed.atom` - Atom feed of content

## TODO

 * http://docs.travis-ci.com/user/languages/go/
 * https://developers.google.com/appengine/docs/go/gettingstarted/helloworld
 * http://dev.mikamai.com/post/68453619468/building-web-apps-with-traffic-the-go-micro-framework
