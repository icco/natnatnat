# natnatnat

[![Build Status](https://travis-ci.org/icco/natnatnat.svg?branch=master)](https://travis-ci.org/icco/natnatnat)

The next iteration in Nat's content management system. Previous versions include:

 * [tumble.io](http://github.com/icco/tumble)
 * [pseudoweb.net](http://github.com/icco/pseudoweb)

Docs: [godoc.org/github.com/icco/natnatnat](https://godoc.org/github.com/icco/natnatnat)

## Install

 0. Install Go
 1. `go get code.google.com/p/xsrftoken`
 2. `go get github.com/pilu/traffic`
 3. `curl https://sdk.cloud.google.com | bash`
 4. `gcloud components update app`
 5. `make run`

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
 * https://developers.google.com/appengine/docs/domain
 * https://developers.google.com/appengine/docs/go/users/reference

### IRC Convos

```
[21:41:33]  <icco>   Question: How are people dealing with posting content from their site to social networks? Manually? Synchronously? Or some nice asynchronous setup?
[21:41:57]  <icco>   I'm trying to come up with a good way to do it asynchronously, but haven't got a good design yet
[21:42:20]  <bear>   aaronpk_ yea, but GPLv3 *blech*
[21:42:21]  <gRegor`>  icco: I use http://brid.gy to publish to Twitter
[21:42:25]  <Loqi>   bear: tantek left you a message 4 days ago: http://indiewebcamp.com/irc/2014-10-05#t1412539342956 ;)
[21:43:01]  <gRegor`>  icco: See https://www.brid.gy/about#publishing for more info
[21:43:14]  <gRegor`>  Several of us use it so can help with questions or problems.
[21:43:16]  <icco>   gRegor`, ah cool. I remember that being mentioned, but I thought it was only for webmentions, not for posting. thanks!
[21:43:29]  <gRegor`>  icco: And snarfed1 is the one who runs it :)
[21:43:51]  <icco>   That's what I thought.
[21:44:01] icco  has a bad habit of trying to reinvent the wheel
[21:44:03]  <gRegor`>  icco: It's pretty cool. You basically send a webmention to the bridgy publish endpoint and voila, it does the magic
[21:44:31]  <snarfed1>   aww thanks gRegor`!
[21:44:41]  <icco>   Ah, and here's the source! https://github.com/snarfed/bridgy
[21:44:57]  <gRegor`>  bridgy publish does have a form field you can paste your URL into to manually publish, too.
[21:45:10]  <icco>   Hosted on GAE, why am I not surprised :p
[21:45:30] snarfed1  has a hammer
[21:45:53]  <icco>   :)
[21:46:14]  <gRegor`>  async is recommended. My UI definitely hangs sometimes while waiting for the bridgy publish response.
[21:46:26]  <icco>   snarfed1, bug me if you guys ever go over that 5GB limit, I know a guy.
[21:46:28]  <gRegor`>  I suspect that's more twitter api slowness than bridgy
[21:47:10]  <snarfed1>   icco: 5gb limit?
[21:47:29]  <icco>   gRegor`, yeah. I want to be async, I've drawn up like three designs for doing it myself though and they all seem wrong or strange
[21:47:50]  <icco>   snarfed1, oh in your readme you mention you backup your data to GCS and there's a 5gb free quota
[21:47:55]  <snarfed1>   gRegor`: re latency, also the fetch of your own site, at least for publish :P
[21:48:22]  <snarfed1>   icco: ahhh got it. heh. thanks! 5G comfortably gets me snapshot backups for over a month, which is fine, but i'll definitely let you know
[21:48:33]  <gRegor`>  Nonsense. My site is always lightning fast! ;)
[21:48:47]  <icco>   :)
[21:49:16]  <KevinMarks_>  Does bridgy publish post to blogger?
[21:49:44]  <kylewm>   KevinMarks_: twitter and facebook only right now
[21:49:47]  <snarfed1>   KevinMarks: not yet, but it's a great feature request
[21:50:27]  <snarfed1>   KevinMarks_: https://github.com/snarfed/bridgy/issues/276
[21:50:45] snarfed1  runs off to change a diaper :P
[21:51:17]  <kylewm>   KevinMarks_: would you want publishing comments or publishing posts?
[21:51:27]  <KevinMarks_>  It does reflect connects back to blogger, but I'd like to send it a posse post
[21:54:14]  <KevinMarks_>  Eg sending http://www.kevinmarks.com/twitterhatespeech.html to http://epeus.blogspot.com/2014/10/how-did-twitter-become-hate-speech-wing.html
[21:55:00]  <danlyke>  iccio: I insert into a table for "statusupdates", and then record numbers from that table into "update_facebook", "update_twitter", "update_flutterby" etc., along with a timestamp and a delay factor. I then have a queue runner that runs every 5 minutes and attempts to push those various updates out, deleting the records from the update_* table as those posts succeed.
[21:57:16]  <kylewm>   icco: i use the RedisQueue python library to send and receive webmentions and do POSSE stuff
[21:57:59]  <icco>   kylewm, interesting, so you just write into a queue when you save the post, and then read from the queue, create a short link and post?
[21:58:15]  <icco>   (that last post meaning share to other sites)
[21:58:29]  <snarfed1>   KevinMarks_: good request. definitely feel free to add to that issue or create a new one
[21:59:13]  <kylewm>   icco: not that exciting, i do basically the same thing danlyke described just using a queue worker and redis instead of a db table
[21:59:49]  <icco>   danlyke, neat, do you use a short link or any sort of truncator on your content?
[22:00:14]  <icco>   kylewm, do you keep any historical record, or just clear the queue once the worker has come through?
[22:01:04]  <danlyke>  I used to, before Twitter started auto-shortening. I do a little bit of prep work, if it's longer than 140 characters and the Twitter shortener wouldn't make it shorter, I crop at the first space before 118 (or the https?://) and insert a link to the full entry on my site.
[22:03:30]  <icco>   ah, twitter does autoshortening on API requests as well? I did not know that, assumed only did it on website posts for some reason
[22:03:51]  <danlyke>  I also explicitly extract links and send them to Facebook as a "FEEDLINK" rather than a "STATUS", so Facebook grabs the preview (makes posting pictures to Facebook actually work nicer for the end user than the native Facebook UI).
[22:04:03]  <icco>   ooh, neat tip
[22:04:07]  <icco>   thanks
[22:04:12]  <danlyke>  yep. 22(?) characters (usually I'm conservative on this).
[22:04:34]  <icco>   I had always heard 20, but I haven't tested in a while.
[22:04:40]  <danlyke>  I use the PHP utility "fbcmd" to do the Facebook posting. A bit of a PITA because Facebook is getting away from authorizing non web apps, so I have to re-auth regularly.
[22:05:00]  <danlyke>  Twitter's URL shortener used to be shorter, but then they went to https:// on everything. That added at least a char.
[22:06:12]  <icco>   ah
[22:07:30]  <snarfed1>   re twitter link lengths: https://dev.twitter.com/docs/counting-characters , https://dev.twitter.com/rest/reference/get/help/configuration#highlighter_933136
[22:08:03]  <icco>   Cool, thanks for the tips peoples. Gonna go give this all a try.
```
