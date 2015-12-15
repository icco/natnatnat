---

layout: post
title: DDoS and You
location: NYC
time: 21:25:38

---

Do you hear horror stories about [China attacking services you use](http://arstechnica.com/security/2015/03/massive-denial-of-service-attack-on-github-tied-to-chinese-government/) or about [4Chan taking out services](http://www.wired.com/2011/12/anonymous-101-part-deux/3/) with their [Low Orbit Ion Cannon](https://en.wikipedia.org/wiki/Low_Orbit_Ion_Cannon)? After hearing stories like that, do you think, "Wait. What?" or, "How does this even work?" or, "Why can random people take down other people's websites?" or especially, "How can I protect myself?" Then this is the article for you.

I'm here to attempt to explain the world of denial-of-service attacks, and to offer some strategies for survival in this crazy internet world.

## What is a DDoS?

A DDoS is a Distributed [Denial-of-Service attack](https://en.wikipedia.org/wiki/Denial-of-service_attack). They happen all the time on the internet, wars played out by computers trying to make other computers inaccessible or overloaded.

The [Digital Attack Map](http://www.digitalattackmap.com/), built collaboratively by Google Ideas and Arbor Networks, displays a snapshot of internet attack activity at any one time. Particularly interesting is the [gallery](http://www.digitalattackmap.com/gallery/), which shows a bunch of exciting days on the internet.

One of my recent favorites is from April 16, 2014, described as, "Volumetric attacks targeting Poland with sustained levels of over 100 Gbps." I haven't taken the time to figure out why this attack happened (because often such things don't make the news, nor does it matter), but it's interesting to know that so much data is being thrown around. For reference, there are 8 gigabits in a gigabyte, so 100 gigabits per second is 12.5 gigabytes of data. A 720p Blu-ray movie is approximately 6.25 gigabytes in size (to make the math easy, since movies are often in the 4- to 10-gigabyte range), so someone was pushing two entire movies every second to computers based in Poland.

A quick tip: Google responds to search queries like, "50 Gigabits in gigabytes," for fast lookups of stuff like this.

[![April 16](http://pseudoweb.net/images/2015/ddos/april16.png)](http://www.digitalattackmap.com/#anim=1&color=0&country=PL&time=16176&view=map)

The Digital Attack Map actually has a fantastic [Understanding DDoS page](http://www.digitalattackmap.com/understanding-ddos/). It includes a few videos on how to use the site, what each part of the site means, and what DDoS is.

The key point is that these attacks can come from any type of network connection. When I talk about defenses, this will be important. But first let's talk about the main types of attacks using the names that the Attack Map uses.

 - TCP Connection Attacks

  > [TCP](https://en.wikipedia.org/wiki/Transmission_Control_Protocol) is the networking protocol of the internet. It provides reliable, ordered and error-checked delivery of data in the form of packets (unlike [UDP](https://en.wikipedia.org/wiki/User_Datagram_Protocol)). However, TCP connections can be made by attackers to never close, and computers (such as load balancers, HTTP servers and routers) have a limited number of connections they can keep open. So if someone can take and hold the connections your computer has available to connect to others, others won't be able to connect to you.

 - Volumetric Attacks

 > A volumetric attack (relating to volume, or how much stuff a three-dimensional object can contain) is what it sounds like. Your network connection is like a pipe ([joke](https://en.wikipedia.org/wiki/Series_of_tubes)) that can only transport a certain amount of data at once. Also, your computer can only process a finite amount of data at once. So if someone starts sending lots of bits to your computer, a couple of different things can happen. Either responses will start to slow down as your computer takes more and more time to process the large requests, or the network connection will slow down because the bandwidth between your server and the internet is diminished by traffic congestion.

 - Fragmentation Attacks

  > Remember how I said TCP "provides reliable, ordered, and error-checked delivery of data"? Well, an attacker can purposely send bad data. Some examples are [SYN Floods](https://en.wikipedia.org/wiki/SYN_flood), [PING Floods](https://en.wikipedia.org/wiki/Ping_of_death), and [Teardrop Attacks](https://www.juniper.net/techpubs/software/junos-es/junos-es92/junos-es-swconfig-security/understanding-teardrop-attacks.html), among [many others](https://tools.ietf.org/html/rfc1858).

 - Application Attacks

 > Application attacks are interesting because they are hard to detect. They look like normal user traffic but target a specific part of an application to bring a server to its knees. For example, imagine a search engine that has a URL that, when hit, performs uncached lookups that are very CPU intensive. An attacker finds this and sends thousands of requests to this URL, which causes the servers to use up all of their CPU resources.

Not all of these attacks are necessarily malicious. Some readers may remember the term "[Slashdotted](https://en.wikipedia.org/wiki/Slashdot_effect)", which referred to a situation when a website was featured on [Slashdot](http://slashdot.org/) and the traffic directed to the site took it offline. We still see this effect from time to time when sites unexpectedly get featured on sites like [Hacker News](https://news.ycombinator.com/) or [Reddit](https://www.reddit.com/).

## Why do I care?

Engineering is a never-ending problem of cost-benefit analysis. With no constraints, an engineering team can prepare for a large set of possibilities of failure given enough imagination, time and money. But in reality, every system has different reliability requirements. For example, my personal website does not need to be reliable as gmail.com, which does not need to be as reliable as a plane's fly-by-wire system.

Imagine your site is down for an hour. Now a day. Now a week. Will this hurt your livelihood? Will it cost you money? Will people die?

If yes, that's a good thing to know, and you should be prepared to make investments to counteract bad outcomes. As the Digital Attack Map website mentions, your attackers can buy a lot of sustained attack power for $125.

## How can I know?

Monitoring! Monitoring is the answer to all problems... or something. Basically, you should have some sort of setup that measures if your system is available and then notifies you if it's not. This advice is incredibly ambiguous because monitoring is a relatively hard problem. For example, for my personal sites, I have a [cron](https://en.wikipedia.org/wiki/Cron) job that emails me if it can't get a [HTTP status code 200](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#2xx_Success) from some of my sites. But it only checks once an hour. So if I have intermittent unavailability, I'll probably never find out. And it uses a static list of sites, so I have a bunch of sites I am not monitoring at all, because I build them and forget to add them.

Obviously, this type of ad-hoc monitoring and alerting strategy is not a valid solution for large systems. There are many systems for monitoring, ranging from the simple, like [Pingdom](https://www.pingdom.com/) or a script you run every hour, to the complex. Some systems I've played with or used in this category include:

 - [Keen IO](https://keen.io/)
 - [Nagios](http://www.nagios.org/)
 - [Monit](https://mmonit.com/monit/)
 - [Prometheus](https://developers.soundcloud.com/blog/prometheus-monitoring-at-soundcloud)
 - [Borgmon](https://www.reddit.com/r/IAmA/comments/177267/we_are_the_google_site_reliability_team_we_make/c82y43e)
 - [statsd](https://github.com/etsy/statsd)
 - [Cacti](http://www.cacti.net/)
 - Many, many others

Once you are monitoring system availability, you can set up alerting. Many companies and teams have different theories about how and when to alert real people if monitoring tools determine that things are broken. [Rob Ewaschuk](http://rob.infinitepigeons.org/) has a nice summary of [Google SRE](http://www.site-reliability-engineering.info/2014/04/what-is-site-reliability-engineering.html)'s alerting philosophy in a doc titled, "[My Philosophy on Alerting](https://docs.google.com/document/d/199PqyG3UsyXlwieHaqbGiWVa8eMWi8zzAn0YfcApr8Q/edit)".

You can also do some load testing to find out how your server can handle an attack. I haven't tested these tools, but there are many options out there to provide insight into how things might hold up.

 - [Apache's ab](https://httpd.apache.org/docs/2.4/programs/ab.html)
 - [bees with machine-guns](https://github.com/newsapps/beeswithmachineguns)
 - [traffgen](http://netsniff-ng.org/)

## What can I do to help me prepare for attacks?

As I mentioned earlier in the article, getting prepared for attacks is a rabbit hole. Because you know very little about who will attack you, when they will attack you and with what force, you can sink a lot of money and time into preparation and still end up down or broken. So I'm going to try to suggest some avenues for preparation and protection, and I'll provide some possible metrics for evaluating how much preparation is enough for you.

I'm going to focus mainly on HTTP server protection, because that's what I know the most about, but in theory this could apply to any sort of TCP/IP connection.

First off, if you're using a [PaaS](https://en.wikipedia.org/wiki/Platform_as_a_service) like [Heroku](https://www.heroku.com/) or [App Engine](https://cloud.google.com/appengine/), your considerations will be different, because your PaaS provider will probably do most of this work for you. If you believe you are going to have a lot of traffic (a spike of 100K QPS, for example), then it might be worthwhile to reach out to your PaaS provider and give them a heads-up. Each provider deals with this differently, but it never hurts to be in contact with your dependencies so that they can prepare for the worst as well.

### Load balancers

If you're hosting things yourself (on a personal server, Colo, EC2, GCE, DO, etc.), I usually suggest setting up a load balancer, which is a system that balances load from a single point to many.

![lbwithapps](https://s3.amazonaws.com/f.cl.ly/items/130j2V1a1o2b1d3f0t28/Screen%20Shot%202015-05-15%20at%2023.25.16%20.png)

A load balancer is also a first line of defense. It can decide quickly whether or not to allow a connection to be passed to the application server based on: 

- The type of connection (Not an HTTP connection? Goodbye!) 
- Where the connection is coming from (Are you coming from a known attacker? Adios!) 
- If the connection is malformed in a way meant to hurt us (Are you a fragmented packet? Dropped!)

Two software load balancers that I often use are:

 - [HAProxy](http://www.haproxy.org/) - [How To](https://serversforhackers.com/load-balancing-with-haproxy)
 - [Nginx](http://nginx.org/) - [How To](http://nginx.org/en/docs/http/load_balancing.html)

If you're using GCE or EC2, Google and Amazon both have load balancers to sell you. These tools are often worth the money because you can take advantage of their scale, and they usually perform better than running your own load balancer on a virtual machine.

 - [AWS Elastic Load Balancing](http://aws.amazon.com/elasticloadbalancing/) for EC2
 - [Google Compute Engine Load Balancing](https://cloud.google.com/compute/docs/load-balancing/)

I mentioned in the pitch for load balancers that you can drop bad packets, and you can and should do that on all of your machines with [iptables](https://en.wikipedia.org/wiki/Iptables) or another software firewall. nixCraft has [a short article on how to set up iptables for protection](http://www.cyberciti.biz/tips/linux-iptables-10-how-to-block-common-attack.html) and a longer article about [configuring Linux server security](http://www.cyberciti.biz/tips/linux-unix-bsd-nginx-webserver-security.html).

### CDN

A CDN or [Content Delivery Network](https://en.wikipedia.org/wiki/Content_delivery_network) is another tool you can use to protect yourself. It takes the load off of your server for static assets: things like images, movies, CSS and Javascript. Instead, you upload your files to the CDN, which serves the files for you, usually in a way that is fast and distributed (for example, by putting one copy of each of your files in five geographically separate data centers).

There are a few CDN providers that are popular:

 - [CloudFlare](https://www.cloudflare.com/)
 - [MaxCDN](https://www.maxcdn.com/)
 - [CloudFront](http://aws.amazon.com/cloudfront/)
 - [Akamai](http://www.akamai.com/html/solutions/network-operator-solutions.html)

Another (often cheaper) way to do this is simply to use a blob storage service. You probably won't get the same level of geographic diversity, but at least your servers will be doing less work of pushing static files. Two popular options are:

 - [GCS](https://cloud.google.com/storage/)
 - [S3](http://aws.amazon.com/s3/)

### Automation

At some point, you'll start getting a lot of traffic. You may have a load balancer set up, but a load balancer is only useful if you can quickly and horizontally scale the number of app servers behind it. The solution to this is automation.

![progression](https://s3.amazonaws.com/f.cl.ly/items/3Y22032z210x3u2V3X1v/Screen%20Shot%202015-05-21%20at%2010.56.59%20.png)

Automation requires a few blog posts on its own, so I'll try to keep this section brief. When I say automation, I'm referring to automating processes that take a lot of time. In the web hosting world, a process that often takes a lot of time is setting up and starting a new app.

The simplest form of automation for this process is to write down the steps needed (with exact commands) to create a new app server. Then, when you need to turn up a new app server, you follow your directions. This is more of a playbook than true automation, but it's a great start.

Once you have everything written down, if you find yourself either iterating on your environment or wanting to turn up machines often, you may want to turn your instructions into code. There are a lot of tools to do this, including:

 - [Puppet](https://puppetlabs.com/)
 - [Chef](https://www.chef.io/chef/)
 - [SaltStack](http://saltstack.com/)
 - [Ansible](http://www.ansible.com/)
 - [Fog](http://fog.io/)
 - [Vagrant](https://www.vagrantup.com/)
 - So many others!

My preferred approach for small projects right now is to write a small Ruby script using Fog, but that's not ideal. People are constantly trying to improve this area, because everyone has different opinions on how it should be done (which is probably why there are so many solutions to the problem).

### Playbooks

I mentioned playbooks briefly above. The idea is that you should write down everything you do to your production environment, especially if you ever plan on doing it again. Things I like to have written up include how to:

 - Create a database snapshot
 - Restore a database snapshot
 - Turn up a new app server
 - Create read-only database slaves and point the apps at them
 - Create common useful SQL queries
 - Set up the dev environment
 - Respond to every alert received

### Backups

Speaking of backups, do them. If you care about your data, I usually advocate at least hourly backups for the past week, and daily backups for the past month. These backups should be written to somewhere that is not your database server. I usually gzip, encrypt and upload my backups to GCS. If you care less, then just do daily backups for the past month.

All code should be versioned. You should never be running code in production that isn't committed somewhere in a canonical place.

Take hard-drive snapshots every time you make a significant change to the system. I usually have the last step of my turn-up automation be a snapshot, and then the first step to let you boot a VM from a saved snapshot instead of building from scratch.

## Acting in the face of an attack

Lets say you find out your site is down. How should you respond?

 1. Evaluate
 2. Stop bleeding
 3. Find and fix root cause
 4. Write postmortem, prevent class of problems

Note that the goal is to stop bleeding before finding and fixing the root cause. We do this because we usually want to minimize downtime. For example, in many attacks, a common first response is to increase capacity (for example, turn up 10 more app servers) and then investigate where the attack is coming from so you can try to block it.

## Further reading and research

TODO: Add links

 - Google SRE
 - Etsy Code as Craft
 - Network Security Books
 - Microservices Book
 - CIDR

Thanks to [David](https://dmpatierno.com/), [Alex](http://alexbaldwin.com/) and __ for the proofreading and edits!
