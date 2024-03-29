*This is an early release. Some features are not yet fully implemented.*

Khutulun
========

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Latest Release](https://img.shields.io/github/release/tliron/khutulun.svg)](https://github.com/tliron/khutulun/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/tliron/khutulun)](https://goreportcard.com/report/github.com/tliron/khutulun)

A distributed declarative orchestrator for services that speaks
[TOSCA](https://www.oasis-open.org/committees/tosca/).

Khutulun is a straightforward, flexible alternative to [Kubernetes](https://kubernetes.io/),
[Nomad](https://www.nomadproject.io/), etc.

Its primary design goal is that the outcome of orchestration would be no different from what a
sysadmin or DevOps team would do themselves. Drifting away from this goal would mean losing control
over orchestration and unraveling the narrative.

In practice, this means that if you want to simply install and run a bare process on a machine,
Khutulun will do that for you. If you want straightforward networking based on reserved TCP ports,
Khutulun won't do anything more than keep track of those ports for you. More complex deployments using
containers, virtual machines, and virtual networks are also supported, but Khutulun's aim is to manage
complexity without getting in the way of simplicity.

Delegates
---------

Khutulun is modular and extensible. Resource types are handled by a cooperative ecosystem of
delegates, the main delegate types being for running compute workloads, for networking, and for
storage. Delegates can call on other delegates using the operator pattern. They can be implemented
as in-process plugins, system services, and even packaged into workloads on the cluster.

Some included resource types and delegates:

* Bare processes: self-contained or otherwise installable executables and scripts
* Containers using [Podman](https://podman.io/) or [Docker](https://www.docker.com/) or
  [systemd-nspawn](https://www.freedesktop.org/software/systemd/man/systemd-nspawn.html)
* Virtualized containers using [Kata Containers](https://katacontainers.io/)
* Pods of containers using Podman or Docker
* Pristine containers using [Distrobox](https://distrobox.privatedns.org/) (on top of Podman)
  or systemd-nspawn
* Virtual machines using [libvirt](https://libvirt.org/)
* TCP port reservation/discovery with support for exposure through [Firewalld](https://firewalld.org/)
* Local or networked directory storage

Plugins can optionally wrap executables in usermode systemd units. This provides a unified admin
experience as well as resilience in the case of failures and restarts.

Note that unlike Kubernetes Khutulun does not demand that every container (or pod) have its own IP
address in an internal network. If desired this feature could be implemented by a networking delegate.

Clusters
--------

Cluster formation is emergent and based on the
[SWIM gossip protocol](https://ieeexplore.ieee.org/document/1028914), with optional support for UDP
multicast for automatic mutual discovery. At the minimum you need just one "seed" host to bootstrap
a cluster, but because all hosts are "masters" the cluster can survive with as little as one arbitrary
host.

Khutulun doesn't distribute its management state among hosts. That's too much responsibility and
hard to design for scale. It's also a solved problem. So, instead Khutulun simply requires that all
hosts have access to the same shared filesystem. A simple NFS share can be enough even for large
clusters. Change coordination is handled via
[fileystem flock](https://man7.org/linux/man-pages/man2/flock.2.html).

What about setting up the cluster hosts? Bare metal tasks like partitioning drives, installing
operating systems, and configuring networking and other essential services? Or cloud tasks like
provisioning virtual machines, virtual storage, and virtual networks? Simply put, that work is out of
the scope of Khutulun. Use a dedicated infrastructure manager instead. Khutulun can interact with such
tools, for example to allow workloads to modify their own cluster, or to use a Khutulun cluster as
a dedicated "management cluster" that, well, manages the hardware of all other clusters.

Included are plugins for [Terraform](https://www.terraform.io/) and
[Ansible](https://www.ansible.com/) that make it easier to include Khutulun installation on your
infrastructure.

By the way, individual hosts do not have to be dedicated to Khutulun and its workloads. You can use
Khutulun to manage services across many machines without having to conceptualize them as a "cloud".

Get It
------

[![Download](assets/media/download.png "Download")](https://github.com/tliron/khutulun/releases)

FAQ
---

### What's wrong with Kubernetes?

Kubernetes is delightfully minimalistic, as orchestrators go, but still makes some potentially costly
decisions:

* The requirement that every pod have its own IP address demands complex container networking
  solutions and is a substantial obstacle for inter-cluster connectivity. For some use cases the cost
  of this requirement is too high to bear. ([Multus](https://github.com/k8snetworkplumbingwg/multus-cni)
  enables "side-loading" networking, but we still need a primary IP address on the Kubernetes
  control plane.) Distributed storage solutions must also be made to participate in this networking
  scheme. Thus developing for the "cloud native" environment is often a synonym for "workarounds for
  an overly opinionated architecture".
* Kubernetes is focused on one kind of activity: pods of Docker-style containers. This means that we
  also require a Docker-style container image repository (either external or internal to the cluster).
  That's a not-insignficant cost. And what if we don't need or want to use containers? Sometimes we want
  just bare processes, or virtual machines, or even other container technologies (e.g. systemd-nspawn).
  (Yes, [KubeVirt](https://kubevirt.io/) enables VMs on Kubernetes, but they have to awkwardly dress up
  as pods and participate in container networking.)
* Kubernetes's resource data model, often represented as YAML manifests, has no relational
  capabilities (except [ownership](https://kubernetes.io/docs/concepts/architecture/garbage-collection/)).
  But cloud workloads are all about relationships, e.g. service meshes. This lack of topological
  expressiveness is an obstacle to application and service modeling. (And, no,
  [Helm charts](https://helm.sh/) are not graphs. Though, as an alternative, check out
  [Turandot](https://turandot.puccini.cloud/), which brings TOSCA to Kubernetes.) Also, the Kubernetes
  data model isn't easily extensible: custom resource definitions require admin access to set up and
  custom resources do not behave exactly like built-in types.
* Relying on etcd for cluster state limits its usability. Etcd documents have a maximum size of ~1KB.
  If you need to share anything more substantial then you'll have to deploy your own solution.

If none of the above is a problem for your use case then by all means stick to Kubernetes. Otherwise,
try Khutulun, because it deliberately attempts to avoid these limitations.

### Why TOSCA?

TOSCA is an open standard with broad industry support. It is, as of version 2.0, a pure
object-oriented language that relies on "profiles", or type libraries, that in Khutulun can work
with delegates to provide specific implementations. Khutulun comes with its own TOSCA profile and
ecosystem of delegates. You are encouraged to add your own.

One of the hallmarks of TOSCA is that every service is a topological graph. Moreover, the edges
of the graph are first-class citizens. This killer feature supercharges your modeling power for the
cloud.

The developers of Khutulun are involved in the TOSCA community and committed to improving the
standard.

### Why support bare processes? Don't containers provide better isolation?

Yes, containers indeed provide better isolation and Khutulun supports them out the box via Podman,
Distrobox (on top of Podman), and systemd-nspawn.

But don't just jump on the bandwagon, ask yourself: Is isolation really what you need for your use
case? And do you understand and are willing to pay for what it costs? We are in the midst of an
architectural shift towards service composition and away from component isolation. Isolation is often
beneficial, and in some specific use cases even necessary, but if isolation technologies get in the way
of collaboration technologies then you're are shooting yourself in the foot. Most container
technologies require you to build ready-to-run container images and stand up container image registries
to store them, adding significant complexity to your development and deployment workflows. Also complex
is managing container networking across clusters. If you entirely own your cluster and workloads then
it might save you mountains of pain to simply use bare processes with bare networking.

Consider Distrobox as a Goldilocks solution: it provides pristine containers that provide only a
minimal operating system but no workloads, so you can run your workloads there instead of on the
bare host. Khutulun will handle the heavy lifting for you. The result may give you the best of both
worlds.

### Why not use a distributed key-value store like [etcd](https://etcd.io/) for management state?

What's wrong with just having a filesystem shared among all hosts? Seriously, why make things more
complicated than they have to be?

Note that etcd has strict limits on the size of documents (~1KB), which is an obstacle for sharing large,
useful binary artifacts. That means that if you need to share large, useful binary artifacts you will
need to deploy yet another storage system. Are we winning yet?

### Why is there no custom Khutulun cluster installer? Why recommend using Terraform and other tools instead?

Infrastructure management is a solved problem. Let's please not reinvent the wheel just for Khutulun
to have its own opinion.

### Why is it called "Khutulun"?

[Khutulun](https://en.wikipedia.org/wiki/Khutulun) (Mongolian: Хотулун) was a fabled Mongolian warrior,
daughter of Kublai Khan's cousin, Kaidu.

She was likely the inspiration for *Turandot*, the protagonist of Count Carlo Gozzi's *commedia dell'arte*
[play](https://en.wikipedia.org/wiki/turandot_(Gozzi)), which in turn inspired
[Giacomo Puccini](https://en.wikipedia.org/wiki/Giacomo_Puccini)'s
[opera of the same name](https://en.wikipedia.org/wiki/Turandot).

And [Puccini](https://puccini.cloud/) is the TOSCA processor that drives Khutulun.

### How do I pronounce "Khutulun"?

* International level: "KOO-too-loon"
* Cosmopolitan level: "CHOO-too-loon" ("ch" like in "Johann Sebastian Bach")
* Expert level: Modern Mongolian "Хотулун" ([video](https://www.youtube.com/watch?v=uP0BagZ-ZCE&t=58s))
