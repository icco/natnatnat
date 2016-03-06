---

layout: post
title: Promoting Reliability in Web Software Companies
location: Chester, CA
time: 18:29:28

---

Building a website is easy. Building a website that makes money is a little harder. Keeping your website online and available as you grow can be very hard. I have heard many reasons for this, including:

 - the company gained users faster than they could handle
 - the developers are too junior
 - the company needs an SRE
 - there is too much technical debt
 - we don't know how to scale
 - _Technology X_ doesn't work for our use case anymore.

The root problem isn't these things. Instead it is a lack of understanding of and promoting of production health inside a company.

What do I mean by "Production Health"? When I say production, I am referring to the software, the infrastructure it runs on, and its dependencies, that your customers see when they visit your website and interact with your product. This can be as simple as a few lines of Ruby running on a Heroku account, a PHP script running on a laptop in your parents' closet, or hundreds of thousands of lines of code from many languages running on thousands of computers in datacenters scattered around the world.

My general rule of thumb is: if you need a piece of software to make money, or if people who aren't software developers use that software, that software is in production.

As for health, let's stick to the dictionary definition: "the state of being free from illness or injury."

What does a healthy production look like? When you go to a doctor, they usually compare you to some ideal version of health. You get a physical so a doctor can tell you what you need to improve, and what is fine.

The first level is: "can our customers use our product?" Often, many people stop there. If the product is a simple web page or a small art project, you probably care about its health like you would a house plant. It's still green. I water it on occasion. People smile when they visit and see it. If it dies, I don't care too much. I'll just compost the current one and go buy another or plant some new seeds.

But let's assume you want to get past this level of "it looks fine" to actually knowing how healthy something is. If you're a human, you often will start weighing yourself, watching what you eat, going to the gym and measuring your progress. With a website, you can do this as well with a single word: Monitoring.

![wheee](http://cl.natw.me/fCos/d)

What do you monitor, you ask? Well that's actually a complicated question, because it depends on what you consider important, and often it's not just one thing. A classic trio of things to monitor for websites are request rate, request duration and request error rate. These three aren't the end-all-be-all of how to measure a service, but they're a decent place to start. But after you're monitoring those three basic things, it's time to start talking to your coworkers and figure out where to go from here.

## SL{I,O,A}

So you've got some services, you're measuring them, maybe you've even got a few graphs you're checking out on occasion. What is next?

First step is to approach your product owner. Sometimes it's you, sometimes it's a product manager, sometimes it's an executive of some sort, and sometimes it's a third party customer. You walk up to them and ask "how do I know if the product is working?" It often takes time, but you can get to the most important metric for the product. The thing that, no matter what, says "yes, we are still ticking". This metric is often not a binary thing, but rather a value. Something along a scale, and if it's perfect, it's at one value, and as it changes, the system gets in a worse and worse state. This is often called a Service Level Indicator, or SLI. A common path is to come up with a few. Some may be the metrics you already graphed, others may be ones you didn't even think of.

Graph those SLIs and sit on them for a few weeks.

Now come back to them. Are the graphs consistent? Do inconsistencies match the times when you suffered degraded service to your customers? If not, then sit down and try and find a better metric (or source for the data). If yes, then this SLI is probably a good candidate to be a Service Level Objective or Agreement!

SLOs and SLAs are used differently depending on the organization. Often an SLA is a legal agreement, while an SLO is just a goal, but they both have the same idea: X% of the time, we will perform at a certain level (our SLI will be at or better than a specific value).

You often hear people talk about this by saying a service has a certain number of nines. For instance, [Google Compute Engine says that](https://cloud.google.com/compute/sla) "the Covered Service will provide a Monthly Uptime Percentage to Customer of at least 99.95%". 99.95% is called three and a half nines of uptime over the last month. That equates to about twenty two minutes of downtime. So for twenty two minutes, every month, your SLI can be below your goal value. I've heard of some networking providers aim for six nines, which is about two and a half seconds of downtime per month, or about thirty seconds per year (also known as incredibly difficult).

As you can imagine, the engineering effort needed between allowing for thirty minutes of downtime per month is a lot lower than making sure you don't have more than thirty seconds of downtime every year. This scale of difference is important when you start having a discussion about SLAs. You now know how your system performs normally because you've been monitoring the SLI, and you have an idea of what is required to maintain that because of the work you and your coworkers have been doing since you started measuring the SLI.

If you find that you aren't meeting your goal, there are often a few questions you can ask:

 - What incidents happened that caused us to be below our goal?
 - Are our dependencies running at their SLA for this month?
 - Is our SLO higher than our dependencies SLOs?

If your dependency isn't meeting its SLO, there isn't much you can do, unless you run the dependency yourself (e.g. it is not owned and run by a third-party). And then you need to figure out why it is failing (more about that in a second).

If your dependency has a lower SLO than you, then you are asking for the impossible. For example, if you are running a service on [Amazon EC2](https://aws.amazon.com/ec2/sla/), and you are promising your system will be up for four nines of the time, you are promising more than your dependency's SLA. EC2 claims it will be up only for three and a half nines of the time, which makes it hard for your service to be any better than that.

## Postmortems and Outages

When a service you run has an outage, priority number one should be getting it back to a healthy state. After that, there are some things to think about:

 - What was the time-to-recovery?

This is the time between when the service went down (not from when you noticed it was down) and the time the system recovered. A good goal for production health is to minimize the mean time to recovery or MTTR. This often involves making things like rollbacks easy, or even automatic.

 - Did we notice that things were broken? If not, how can we fix that?

This is a big thing. Did our monitoring systems know things were broken? Did they tell someone (with an app like [pagerduty](https://www.pagerduty.com/) or with email or with a text message or something)? Did that someone respond to being alerted that things were broken? If you answered no to any of these questions, put an action item on someone to fix it.

 - How can we prevent this type of outage in the future?

This is often the most interesting part of the problem to me and can lead to total over-engineering or serious bike-shedding discussions if you're not careful. But if this type of outage happens twice, it's usually time to start making sure it doesn't happen again.

 - Should we write a postmortem?

The rule at Google, as I understood it, was if someone (anyone) asks for a postmortem, you need to write one. I like this rule. But whatever rule of thumb you create, you should make sure everyone involved with the service understands it.

Postmortems are a document where you write down what happened (usually with a timeline as specific as it can be, including links to commits, emails, etc), why those things happened (the root cause of the outage), and what you are going to do about it (action items). Often answering the first two questions is a great way to start coming up with action items.

Then send the doc around. Make sure everyone involved in the incident agrees with what the doc says. Also, make sure the document does not point blame, but rather states things as they happened. It doesn't matter if someone broke production, no one should get fired over these things or be given a talking to or whatever. The point is to find systemic issues in the system, not make yourself feel better by saying how much someone else sucks. Finally store it somewhere you can read about it in the future.

## Taking on some risk

If you've been monitoring your service for a while, and it's doing well, I've got some good news for you: it's time to take that error budget for a spin. When you've been above your goal for a while (lets say you had 100% uptime the last three months), you can use that as an excuse to take on risk. You shouldn't take on risk for no reason, but a little risk let's you try new or overly dangerous things that you might not have done if you didn't know how stable your system was.

## Culture

Next is the hard question, the one that brought me to write this article in the first place: How do you convince coworkers that production health and reliability are important?

I'll be up front and honest: I am not sure. I think this is one of the core problems with engineering cultures in software companies right now.

A lot of times when putting responsibility on people, directly asking them to be the person who responds first to an outage helps them understand and respect production. It also helps people understand what are the stable things in the system and what are the broken things. You want to get away from mentalities such as "operations can deal with that" or "I write features, I'm not a maintainer" or "I don't own that code any more, I don't care".

Keeping people involved with postmortems is another way to keep people interested. "The code broke this way, any suggestions on how we can fix that?"

Kripa Krishnan wrote a great article called ["Weathering the Unexpected"](https://queue.acm.org/detail.cfm?id=2371516) which talks about Google's DiRT (Disaster Recovery Testing). This event is a week long event where you plan on taking down parts of the system or trying new things you might only do in the case of a disaster. Common tasks you might want to go through with your team:

 - Can we recover from a backup?
 - What happens when the database goes away?
 - Can we survive a DDoS attack?
 - Can we start up the service in another data center?

Create a safe scenario to test these questions, or others your teammates might have, and then do them! Watching things blow up or survive is exciting, and can often bring a team together.

The thing to remember is that production health is everyone's responsibility. Even if you have a dedicated operations or site reliability engineering team, they can't be all-seeing and all-knowing. Making sure everyone feels comfortable changing production and responsible for production is key.

## More on this topic

 - [What is Site Reliability Engineering?](http://www.site-reliability-engineering.info/2014/04/what-is-site-reliability-engineering.html) - an interview with one of the fathers of SRE, Google's Ben Treynor.
 - [John Allspaw's Blog](http://www.kitchensoap.com/) - this blog is full of amazing articles on the topic of web operations. Highly recommend.
 - [Logging v. instrumentation](http://peter.bourgon.org/blog/2016/02/07/logging-v-instrumentation.html) - an article exploring how to get data about the health of your applications.
 - [A tweet on monitoring philosophy](https://twitter.com/LindsayofSF/status/692191001692237825) by @LindsayofSF
 - [Engineering Reliability into Web Sites: Google SRE](https://static.googleusercontent.com/media/research.google.com/en/pubs/archive/32583.pdf) - a slide deck on SRE by Alex Perry at Google.
 - [The Seemingly Unfixable Crack in the Internet’s Backbone](https://www.technologyreview.com/s/540056/the-seemingly-unfixable-crack-in-the-internets-backbone/) - a nice article that should scare you about how fragile the internet is.
 - [How does the internet work?](https://web.stanford.edu/class/msande91si/www-spr04/readings/week1/InternetWhitepaper.htm) - a useful whitepaper for understanding everything outside of your system.


