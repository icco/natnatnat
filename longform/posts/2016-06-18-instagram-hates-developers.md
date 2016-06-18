---

layout: post
title: Instagram hates developers
location: Brooklyn, NY
time: 18:53:45

---

I love the internet. The internet is a place where people can create amazing content and share it among their peers. They can [post a video of themselves singing](https://www.youtube.com/watch?v=eQOFRZ1wNLw) and get a recording contact. They can create [art inspired by their favorite art](http://behindinfinity.deviantart.com/art/Death-Note-This-Is-Heaven-52682456) and share it with millions of people. They can [share fresh ideas](https://dribbble.com/shots?sort=recent) with their industry. Or whatever they want. Because in general, the internet is a place to allow for people to share content openly.

But lately, Instagram has decided it hates developers. Many sites on the internet let their users use web APIs to get machine parsable information about themselves and their followers. Some example sites that do this are [Twitter](https://dev.twitter.com/overview/documentation), [Youtube](https://developers.google.com/youtube/), [Flickr](https://www.flickr.com/services/api/), [Dribbble](http://developer.dribbble.com/v1/), [DeviantArt](https://www.deviantart.com/developers/) and many others. Instagram has begun limiting the information you can get out of them.

As of [June 1st, 2016, Instagram has decided](http://developers.instagram.com/post/133424514006/instagram-platform-update) as a developer you can no longer get this information about yourself or other users. This was first [announced in November](http://thenextweb.com/dd/2015/11/17/instagram-limits-developer-api-access-with-new-app-review-process/).

As far as I can tell, Instagram doesn't want anyone that isn't a business to build apps against their APIs. I mean I knew Facebook hated developers, but [even Facebook](https://developers.facebook.com/) lets you get some data from them and build apps that post content to their site.

Instagram has never respected the web, you still can't see the content you like on the site, but now, unless you are a business, any content you put in, you can't get it back out in any way besides the Instagram app.

I figured, "oh, Instagram was created by a smart small group of devs, maybe this is all just them covering their ass, and really they are approving any app who applies."

This is not the case. Let's look at their flow for submitting an application for review.

First off, you need to be a company:

![need to be a company screenshot](https://s3.amazonaws.com/f.cl.ly/items/0w243s3u2f0U4041342s/Screen%20Shot%202016-06-18%20at%2014.21.48%20.png)

Second off, there are all these options for apps to review:

![example app screenshot](https://s3.amazonaws.com/f.cl.ly/items/2F2X2x2n302K1X1U2o39/Screen%20Shot%202016-06-18%20at%2014.22.31%20.png)

Don't let that fool you, only three are actually valid.

 - My app allows people to login with Instagram and share their own content.
 - My product helps brands and advertisers understand, manage their audience and media rights.
 - My product helps broadcasters and publishers discover content, get digital rights to media, and share media with proper attribution.

It is kind of nice UX that Instagram will actually not let you submit an application for review if it doesn't meet one of these three criteria. So at least there is that.

Anyway, why am I so angry? Beyond coming from an age of the internet where all content was available to everyone, I just want a copy of my content.

When I post or like something on the internet, I have a cron job that runs on a server, grabs a link to the content and [the OEmbed link](http://oembed.com/). It caches that in a gigantic JSON file (because plain text is best database), and then uploads it to http://mood.natwelch.com. Then I have a site I can scroll through content I like. If I want to comment on it, I can click through to the original work.

This use case, and describing this use case as a business, have both been rejected. So I've resorted to writing a ranty blog post that will be ignored by Instagram. I wouldn't even care if my friends weren't uploading beautiful content to Instagram, I'd go somewhere friendlier to developers, like Flickr or 500px. But alas here we are. I want to participate in the community. I post photos to Instagram. I comment on my friends artwork. I share it out. I'd like to my favorite content on Instagram that my friends are creating, but I can't get to it, so I can't even share it with you.

Look, Instagram, I know you've given up. You don't have to do anything anymore, everyone uploads content to you, and the only improvements you've made in the last few years are launching Android support and change your logo. Oh, right, and you added support for non-square photos and videos. But at least give us access to our content (your `public_content` permission scope) without having to beg you for it.

Sorry for the rant, I'm going to go back to building side projects on platforms that care about me, or at least let me try.

/Nat
