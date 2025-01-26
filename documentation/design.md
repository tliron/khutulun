[⇐ to main site](https://khutulun.org)

Design Choices
==============

Dual-Paradigm
-------------

Khutulun is intentionally *both* data-driven (a.k.a. [declarative](https://en.wikipedia.org/wiki/Declarative_programming)) *and* event-driven (a.k.a. [imperative](https://en.wikipedia.org/wiki/Imperative_programming)). Let's go over the why and how.

#### Can I Please Just Engineer?

Our starting point is to acknowledge the failings of fanatically declarative approaches. They are often reductively opinionated, only clumsily extensible, and come with automagical orchestration pipelines that can make some hard tasks very simple, but unfortunately can make the simplest tasks hard.

For example, as an engineer you might have a straightforward solution to an orchestration challenge, a simple script that tweaks a few configuration parameters, calls a few commands, etc. But if your orchestrator is designs to be anti-code, then your experience is going to be the opposite of straightforward as you fight your way through "not my use case" declarative schemas, and "my way or the highway" pipelines. Can I please just run my script?

#### I Disagree with Your Opinion

If you're a systems architect, you would appreciate that Khutulun aims to get out of your way by allowing you to inject code that realizes *your* opinion. Do you not agree with the design of our built-in networking implementation? Then tweak it or roll your own. You provide the technology and logic while Khutulun manages the boilerplate of coordination and delivery.

And if you are a sysadmin, you would appreciate that our built-in implementations embrace and promote a norm by which the end result is something that you could have put together yourself. We want you to be able to interact with individual cloud machines and programs using your everyday tools, again without Khutulun getting in your way. In other words, our goal is comprehensibility, because if experts can't make heads or tails of the cloud, then you'll be at the mercy of the orchestrator. Khutulun wants to be a helping hand, not an obstacle, and definitely not a [single point of failure](https://en.wikipedia.org/wiki/Single_point_of_failure).

#### The Limits of Intent

Another problem with fanatically declarative approaches it that they want the data to be your [single source of truth](https://en.wikipedia.org/wiki/Single_source_of_truth), a.k.a. a [digital twin](https://en.wikipedia.org/wiki/Digital_twin).

In reality, however, clouds are untamable, often deviating from your intent as soon as you deploy your workloads. In other words, Day 1 is already Day 2. From changes that happen due to hardware failures and network fluctuations, to self-healing and self-scaling software, it's only where and what you deploy that is "true", *not* your intent.

Khutulun doesn't even try to achieve impossible truth enforcement, because such attempts are invariably unreliable and inefficient. It doesn't waste resources by continuously polling the cloud in order to pull in data that would already be out-of-date. There is no endless reconciliation loop. Instead, Khutulun provides tools to update and validate data only when triggered to do so by events that *you* define.
