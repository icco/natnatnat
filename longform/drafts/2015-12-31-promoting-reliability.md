---

layout: post
title: Promoting Reliability in Web Software Companies
location: Chester, CA
time: 18:29:28

---

Building a website is easy. Building a website that makes money is a little harder. Keeping your website up as you grow it can be very hard. I've heard many reasons for this. Sometimes it's because the company gained users faster than they could handle, because the developers are too junior, because the company needs an SRE, because there is too much technical debt, or one of many other excuses. After a conversation held at RC, I realized it is usually none of these things. Instead it is a lack of understanding and promoting production health inside the company.

What do I mean by "Production Health"? When I say production, I am referring to the software, and the infrastructure it runs on, that your customers see when they visit your website and interact with your product. This can be as simple as a few lines of Ruby running on a Heroku account, a PHP script running on a laptop in your parent's closet, or hundreds of thousands of lines of code from many languages running on thousands of computers in datacenters scattered around the world.

My general rule of thumb is, if you need it to make money, or if people who aren't software developers use it, it is production.

As for health, let's stick to the dictionary definition: "the state of being free from illness or injury."

What does a healthy production look like? When you go to a doctor, they usually compare you to some ideal version of health right? You get a physical so a doctor can tell you what you need to improve, and what is fine.

The first level is "can our customers use our product?" Often many people stop there. If the product is a simple web page or a small art project, you probably care about its health like you would a house plant. It's still green, I water it on occasion, people smile when they visit and see it. If it dies, I don't care too much, I'll just go buy another.

But let's assume you want to get past this level of "it looks fine" to actually knowing how healthy something is. If you're a human, you often will start weighing yourself, watching what you eat, going to the gym and measuring your progress. With a website, you can do this as well, with a single word: Monitoring.

 gif here?

What do you monitor you ask? Well that's actually a complicated question, because it depends on what you consider important, often it's not just one thing.

## RESEARCH:

https://www.technologyreview.com/s/540056/the-seemingly-unfixable-crack-in-the-internets-backbone/

https://www.youtube.com/watch?v=y1OJlJ9COg0

https://web.stanford.edu/class/msande91si/www-spr04/readings/week1/InternetWhitepaper.htm

http://www.site-reliability-engineering.info/2014/04/what-is-site-reliability-engineering.html

https://zulip.com/#narrow/stream/programming/topic/writing.20bug-free.20code

https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/32583.pdf

https://www.blacksintechnology.net/google-site-reliability-engineering-reasoning-about-systems/

http://www.oreilly.com/pub/e/2712

http://sigops.org/sosp/sosp15/current/2015-Monterey/008-tang-online.pdf

https://www.youtube.com/watch?v=4fFDFbi3toc

https://plus.google.com/+MarcBevand/posts/fBfCsaXReH5


It's often hard to convince stake holders in a business of the classic software maintance rule of thumb that [Software development is 80% maintanence](https://programmers.stackexchange.com/questions/47991/where-is-this-from-originally-80-of-the-lifetime-cost-of-a-piece-of-softwa).

Hey, this is a fantastic thread, lots of amazing discussion in here. I started writing this, and it turned into a bit of a rant... sorry about that... Everything below is my opinion based on my experience, but it jives with a lot of what’s above


I’ve worked in both environments of optimizing for mean time to failure, and mean time to recovery. I very much prefer the world of MTTR, but often the constraints of your world require you to focus on both. For example: a cloud Virtual Machine hosting company. If a VM goes down, you are having a customer outage. So your VM uptime should be used as a measurement of MTTF. But then, the other thing is, you only have so much capacity to serve customers, so when there is a failure, while you do want to be able to debug, you also want to make sure that capacity is available for customers to create VMs on. You could measure something % capacity unavailable for VM creation and not serving VMs.

This touches on what a few people have said, but I don’t think was emphasized enough, monitoring is the key to promoting reliability within your company. Having easy to read graphs that tell you high level system status and reliability is the first step to a healthy production environment. The second step is providing graphs for debugging that you can drill into. For example, let’s say I have a graph of global failures per 1000 VM hours. I should be able to take that grab and then look at every data center, and then from a datacenter’s failures / 1000 VM hours, drill down into the list of machines, maybe sorted by recent failures, and see commonalities between machines (for example versions deployed to machines, average VM age, load, etc)

Often how this conversation goes down at Google (or at least how it went down on the three teams I worked with during my time as an SRE) is first try to measure ~everything. Get basic data collection pipelines in place. Make sure logs and historical metrics about machines are being saved and are query-able (sp?). Next step is to figure out "What is important". Often this is by looking at what’s important to the business, and creating a number (often called a SLI or Service Level Indicator) that you can graph. Usually focusing on five is a good place for a largish service. These SLIs determine: are we having an outage, do we need to page someone, and is our overall system healthy. Once you start having SLIs, you can go to your product team and start talking about where we want the system to be. When we say four nines of uptime, which SLI are we talking about? Are we there now? Do we want to promise the business that we will have four nines for all of these SLIs? Do we want to promise the customer that? Setting these goals (we want 0.00001 VM failures per 1000 VM hours) is often called setting Service Level Objectives (SLOs) and creating a legal agreement by which to give to customers (we will meet our SLO 99.99% of the time) is often called an SLA or Service Level Agreement.

By making these metrics available to everyone, and making sure everyone is aware of the expectations of where we want to be, and allow for people to dig down and research where the data is coming from (I see this spike, by drilling down into the data, I can see what caused the spike and fix it), it makes reliability something of a goal for the whole team, and also empowers developers across the company to find out what is making their system unreliable. Also, by putting internal and external business goals around reliability, it helps remind everyone that reliability is everyones problem.

Note that most of these examples focused on a MTTF focused SLI, but having a mixture of MTTR and MTTF SLIs is the best bet for a healthy production environment.

As for Code Freezes at Google: it was usually viewed as an Ops smell, unless the business would be put at serious risk. Analytics Frontend and Wallet for example have large code freezes around the holidays, as do parts of Core Infrastructure. Most other teams were recommended to limit their code freezes only to times when there would be severe lack of people availability (AKA, don’t deploy past 3pm, don’t deploy on Fridays, the day before holidays, weekends and actual holidays).

