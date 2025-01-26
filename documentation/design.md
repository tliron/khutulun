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

What's Wrong with Kubernetes?
-----------------------------

### Containers

First, the requirement that almost *all* workloads live in Docker (OCI) containers—and be packaged in container images hosted in a container image repository—is costly. Whatever specific advantages containers have in some use cases, in other situations they present an unnecessary hurdle. Indeed, running bare code on the host is often preferable and for some use cases is required. Developers can attest to the pain of having to package code in container images, uploading it to an image repository (which they have to stand up and maintain), and patiently wait for Kubernetes to pull it. This roundabout is non-deterministic and can be unreliable.

Khutulun doesn't require containers. If your use case doesn't need them, don't use them. Just run refreshingly normal bare processes, which can be installed or deployed using whatever mechanism makes sense to you. Khutulun does support containers if you need them, including alternative containerization technologies, such as [nspawn](https://www.freedesktop.org/software/systemd/man/systemd-nspawn.html), which does not require an image repository. Note that if you appreciate the isolation that containers offer, you can consider Wasm as an alternative. Virtual machines work, too, without any special networking requirements. Note that Kubernetes supports virtual machines, too, using add-ons like [KubeVirt](https://kubevirt.io), but they are very awkward because they must participate in container networking. Which leads us to the next point.

### Networking

Second, networking is poorly thought out, with especially costly consequences. Kubernetes pods must be connected to an internal container network, which then needs to be connected to the other nodes in the cluster as well as everything outside the cluster. This is called *container networking*, and it's an astoundingly complex self-inflicted challenge. It requires a full network management apparatus (DHCP, DNS, NTP), and the deployment of complex "fabrics" to enable interoperability with the "actual" host networks. Users go so far as to rely on in-kernel eBPF optimizations to amortize overhead costs of this layer. Just as astoundingly, Kubernetes started with only IPv4 support (in 2014!), so IPAM is also often needed. To make sure your clusters don't have overlapping IPv4 subnets, you will have to orchestrate those, as well.

So, just as with containers themselves, whatever advantages container networking has in some use cases, the cost is absurdly high for *all* users.

With Khutulun, we leave all the above up to you. If you're willing to pay the price, we give you container networking. If you can't or don't want to, we let you use the boring old host networks. Supporting various networking solutions is indeed a central goal for Khutulun, and because it's built on graphs, with relationships being first-class citizens (see the next point), Khutulun has a strong basis for modeling complex network topologies—including container networking.

### Relationships

Third point: There are no systemic relationships between entities in Kubernetes. The only obvious one is "ownership", to support garbage collection of children when an owner is deleted, and it's handled as a second-class citizen via metadata. This is a very unfortunate omission from Kubernetes's data model, because graphs are everywhere you look in the cloud. From modeling networking infrastructure (as noted above), through coordination between microservices (service meshes), to logical inter-dependencies for scalability and healing policies, all the way down to the cluster machine infrastructure itself.

Khutulun *starts* with graphs. Relationships (graph edges) get as much attention as nodes (graph vertexes). Moreover, relationships are dynamic, meaning that they are formed, dissolved, and reformed imperatively. An event, e.g. a node failure, can cause the whole graph to be updated on demand. If you store your data in a graph database (see next point), you can made good use of its ability to process complex path-following queries and gain a true topological view of your
deployments.

### Data Storage

Fourth, Kubernetes's choice of using just one storage engine, in particular [etcd](https://etcd.io), was shortsighted. The idea of not having to deploy a separate storage service and instead letting the cluster's nodes distribute the data among themselves might seem convenient at first. But any attempt to treat etcd as a serious database ends in disappointment. It sorely lacks capabilities for orchestrating data in large clouds, features that are readily and reliably available in mature database solutions.

With Khutulun we intend to support many popular, mature databases, including those services hosted by public cloud providers. Graph databases are especially useful for large-scale deployments with complex topologies.

### Declarative Hell

Finally, Kubernetes is far too declarative. Extending it with code requires writing an "operator" (custom controller), creating at least one custom resource and participating in the declarative dance with other controllers. Kubernetes is "event driven" only behind the scenes. Indeed, Just standing up a vanilla Kubernetes cluster with no running pods will show non-negligible CPU usage as there is a constant hum of polling and updating. Developers can attest to the pain of having to wait for controllers to kindly respond to updates, and when a chain of operators is involved it can be prohibitively slow and finicky.

As we point out above in "Design Choices", when everything is declarative some hard things become easy, but many easy things become very hard indeed. Khutulun is dual-paradigm, both declarative and imperative. Wherever possible, it provides a hook for you to insert your own custom code.

### Is Kubernetes Really Unfixable?

It's not easy to imagine Kubernetes growing out of its current set of problematic opinions. In a way, it's a victim of its own success: Now that it's being used in production environments, it can't afford to break critical deployments by fundamentally changing how it operates. It will surely continue to improve (e.g. there are efforts to allow for plugabble storage), but the basic design is what it is.

Khutulun is born from the necessity of going back to the drawing board, learning from successes and mistakes of the past decade, and building a system that's flexible enough that it can evolve with the times and needs.
