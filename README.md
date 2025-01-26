*Work in progress, not ready for general use. See also [previous work](https://github.com/tliron/go-khutulun).*

Khutulun
========

An orchestrator for clouds.

Its goal is to provide an unopinionated platform for managing a diversity of compute and networking technologies distributed across clusters of computers. And to make it easy to add support for new technology needs and trends as they emerge.

Out of the box we provide implementations for the following types of workloads:

* Bare processes: self-contained or installable executables and scripts (Khutulun *does not* force you to use containers)
* Containers and pods: via [Podman](https://podman.io/), [Docker](https://www.docker.com/), or [nspawn](https://www.freedesktop.org/software/systemd/man/systemd-nspawn.html)
* Pristine containers: via [Distrobox](https://distrobox.privatedns.org/) (on top of Podman or Docker)
* Virtualized containers: via [Kata Containers](https://katacontainers.io/)
* Virtual machines: via [libvirt](https://libvirt.org/)

For networking, we currently have:

* [OVN](https://www.ovn.org/)
* Straightforward TCP/UDP port reservation/discovery with [Firewalld](https://firewalld.org/) protection (simple is often best)

Khutulun's data system is [Floria](https://github.com/tliron/rust-floria). It is purposely developed as a separate project in order to ensure its integrity and coherence.

Bias
----

We roll out the red carpet for two standards:

[Wasm](https://webassembly.org/) is our best-supported choice for implementing orchestration logic. We love that Wasm is natively sandboxed and runs well on the many kinds of CPU architectures found in clouds. One blob for everywhere. Even if you don't use Wasm for your own implementations, all the built-in ones do and will "just work" in your clouds. Floria comes with an SDK to make it easy to develop your own Wasm extensions in Rust, but any programming language that targets Wasm would also work.

[TOSCA](https://www.oasis-open.org/committees/tosca/) is a YAML-based language for designing cloud services. It's object-oriented, extensible, and has a packaging format, CSAR, which can bundle Wasm implementation blobs together with the design specifications. Khutulun natively supports TOSCA and CSAR via the [Puccini](https://puccini.cloud/) library, which is developed in tandem with Khutulun. All our included examples are based on TOSCA, and we provide TOSCA profiles for all supported workloads and networking types. Even if you don't choose to use TOSCA for your own workflow, you will indirectly benefit from our efforts to support all of its powerful features. As an OASIS open standard, anyone can contribute to its evolution. The developers of Khutulun are longtime active participants in the TOSCA standards committee and the ecosystem.

Design Choices
--------------

Khutulun is intentionally dual-paradigm, both data-driven (declarative) and event-driven (imperative).

#### Can I Please Just Engineer?

This starting point is a reaction to the failings of fanatically declarative approaches. They are often reductively opinionated, only clumsily extensible, and come with automagical orchestration pipelines that can make some hard tasks very simple, but unfortunately can make the simplest tasks hard.

For example, as an engineer you might have a straightforward solution to an orchestration challenge, a simple script that tweaks a few configuration parameters, calls a few commands, etc. But if your orchestrator is anti-code, then your experience is going to be the opposite of straightforward, as you fight your way through "not my use case" declarative schemas, and "my way or the highway" pipelines. Just let me run my script!

#### I Disagree with Your Opinion

If you're a systems architect, you would appreciate that Khutulun aims to get out of your way by allowing you to inject code that realizes *your* opinion. Do you not agree with the design of our built-in networking implementation? Then tweak it or roll your own. You provide the technology and logic while Khutulun manages the boilerplate of coordination and delivery.

And if you are a sysadmin, you would appreciate that our built-in implementations embrace and promote a norm by which the end result is something that you could have put together yourself. We want you to be able to interact with individual cloud machines and programs using your everyday tools, again without Khutulun getting in your way. In other words, our goal is comprehensibility, because if experts can't make heads or tails of the cloud, then you'll be at the mercy of the orchestrator. Khutulun wants to be a helping hand, not an obstacle.

#### The Limits of Intent

Another problem with fanatically declarative approaches it that they want the data to be your "single source of truth". In reality, however, clouds are untamable, often deviating from your intent as soon as you deploy your workloads. In other words, Day 1 is already Day 2. From changes that happen due to hardware failures and network fluctuations, to self-healing and self-scaling software, it's only where and what you deploy that is "true", *not* your intent.

Khutulun doesn't even try to achieve impossible truth enforcement, because such attempts are invariably unreliable and inefficient. It doesn't waste resources by continuously polling the cloud in order to pull in data that would already be out-of-date. There is no endless reconciliation loop. Instead, Khutulun provides tools to update and validate data only when triggered to do so by an event.

FAQ
===

Please [read it](https://github.com/tliron/khutulun/blob/main/FAQ.md)!

License
-------

Like much of the Rust ecosystem, licensed under your choice of either of

* [Apache License, Version 2.0](https://github.com/tliron/khutulun/blob/main/LICENSE-APACHE)
* [MIT license](https://github.com/tliron/khutulun/blob/main/LICENSE-MIT)

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in the work by you, as defined in the Apache-2.0 license, shall be dual licensed as above, without any additional terms or conditions.
