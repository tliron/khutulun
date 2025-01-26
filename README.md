*Work in progress, not ready for general use. See also [previous work](https://github.com/tliron/go-khutulun).*

Khutulun
========

A dual-paradigm (declarative *and* imperative) orchestrator for clouds.

<p align="center">
    <img src="https://khutulun.org/assets/media/logo.png" width="120" height="179" alt="The seal of Kaidu, Khutulun's father" />
</p>

Its goal is to provide a platform for managing a diversity of compute and networking technologies distributed across clusters of computers, from Day 0 to Day 2, and to make it easy to add support for new technology needs and trends as they emerge.

Out of the box we provide implementations for the following types of workloads:

* Bare processes: self-contained or installable executables and scripts (Khutulun *does not* force you to use containers)
* Containers and pods: via [Podman](https://podman.io/), [Docker](https://www.docker.com), or [nspawn](https://www.freedesktop.org/software/systemd/man/systemd-nspawn.html)
* Pristine containers: via [Distrobox](https://distrobox.privatedns.org) (on top of Podman or Docker)
* Virtualized containers: via [Kata Containers](https://katacontainers.io)
* Virtual machines: via [libvirt](https://libvirt.org)

For networking, we currently have:

* [OVN](https://www.ovn.org)
* Straightforward TCP/UDP port reservation/discovery with [Firewalld](https://firewalld.org) protection (simple is often best)

Standards
---------

Khutulun rolls out the red carpet for two standards:

[TOSCA](https://docs.oasis-open.org/tosca/TOSCA/v2.0/TOSCA-v2.0.html) is a YAML-based language for designing cloud services. It's object-oriented, extensible, and has a packaging format, CSAR, which can bundle Wasm implementation blobs together with the design specifications. Khutulun natively supports TOSCA and CSAR via the [Puccini](https://puccini.cloud) frontend, which is developed in tandem with Khutulun. Included in Khutulun are TOSCA profiles for all supported workloads and networking types. The developers of Khutulun are longtime active participants in the TOSCA standards committee and the ecosystem.

[Wasm](https://webassembly.org) is our best-supported choice for implementing orchestration logic. We love that Wasm is natively sandboxed and runs well on the many kinds of CPU architectures found in clouds. One blob for everywhere. Even if you don't use Wasm for your own implementations, all the built-in ones do and will "just work" in your clouds. Floria comes with an SDK to make it easy to develop your own Wasm plugins in Rust, but any programming language that targets the Wasm Component Model would also work.

Technology and Design
---------------------

Khutulun's data system is [Floria](https://floria.khutulun.org), which handles storing and querying the orchestration data as well as running code on it. It is our manifestation of the dual-paradign philosophy. Floria is purposely developed as a separate project in order to ensure its integrity and coherence.

At it simplest, Khutulun comprises:

* A distributed *in situ* event bus for Floria, which can respond to local events by executing plugins locally
* A repository of useful plugins for Floria (Wasm blobs)
* A repository of TOSCA profiles for designing cloud services that make use of these plugins

Documentation
-------------

* [Design Choices](https://khutulun.org/documentation/design)
* [Frequently Asked Questions](https://khutulun.org/documentation/faq)

License
-------

Like much of the Rust ecosystem, licensed under your choice of either of

* [Apache License, Version 2.0](https://github.com/tliron/khutulun/blob/main/LICENSE-APACHE)
* [MIT license](https://github.com/tliron/khutulun/blob/main/LICENSE-MIT)

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in the work by you, as defined in the Apache-2.0 license, shall be dual licensed as above, without any additional terms or conditions.

The logo was made by [Bahatur](https://commons.wikimedia.org/wiki/User:Bahatur) and was copied from [here](https://en.wikipedia.org/wiki/File:Kaidu%27s_tamga.png). It is available for use under the [Creative Commons Attribution-ShareAlike 3.0 Unported license](https://creativecommons.org/licenses/by-sa/3.0/).
