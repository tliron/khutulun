Khutulun - FAQ
==============

Why is it called "Khutulun"?
----------------------------

[Khutulun](https://en.wikipedia.org/wiki/Khutulun) (Mongolian: Хотулун) was a Mongolian warrior princess and daughter of Kaidu Khan.

She was likely the inspiration for *Turandot*, the protagonist of Carlo Gozzi's [*commedia dell'arte* play](https://en.wikipedia.org/wiki/turandot_(Gozzi)), which in turn inspired [Giacomo Puccini](https://en.wikipedia.org/wiki/Giacomo_Puccini)'s [opera of the same name](https://en.wikipedia.org/wiki/Turandot).

And [Puccini](https://puccini.cloud/) is Khutulun's TOSCA frontend.

See, it all makes perfect sense.

How do I pronounce "Khutulun"?
------------------------------

1. International level: "koo-too-LOON"
2. Cosmopolitan level: "khoo-too-LOON" ("kh" like the "ch" in "Johann Sebastian Bach")
3. Expert level: "Хотулун" ([video](https://www.youtube.com/watch?v=uP0BagZ-ZCE&t=58s))

What's wrong with Kubernetes?
-----------------------------

### Containers

First, the requirement that almost *all* workloads live in Docker (OCI) containers—and be packaged in container images hosted in a container image repository—is costly. Whatever specific advantages containers have in some use cases, in other situations they present an unecessary hurdle. Indeed, running code on the host is often preferrable and for some use cases is required. Developers can attest to the pain of having to package code, upload it to an image repository (which they have to stand up and maintain), and patiently wait for Kubernetes to pull it. In production environments
this roundabout is non-determenistic and can be unreliable.

Khutulun doesn't require containers. If your use case doesn't need them, don't use them. Just run refreshingly normal bare processes, which can be installed or deployed using whatever mechanism makes sense to you. Khutulun also supports alternative containerization technologies, such as [nspawn](https://www.freedesktop.org/software/systemd/man/systemd-nspawn.html), which does not require an image repository. Also, if you enjoy the isolation that containers offer, consider Wasm as an alternative. Virtual machines work, too, without any specific networking requirements. Kubernetes does have virtual machine add-ons, such as [KubeVirt](https://kubevirt.io/), but they are very awkward because they must participate in container networking. Which leads us to the next point.

### Networking

Second, networking is poorly thought out, with especially costly consequences. Kubernetes pods must be connected to an internal container network, which then needs to be connected to the other nodes in the cluster as well as everything outside the cluster. This is called *container networking*, and it's an astoundingly complex self-inflicted challenge. It requires a full network management apparatus (DHCP, DNS, NTP), and the deployment of complex "fabrics" to enable
interoperability with the "actual" host networks. Users go so far as to rely on in-kernel eBPF optimizations to amortize overhead costs of this layer. Just as astoundedly, Kubernetes started with only IPv4 support (in 2014!), so IPAM is also often needed. To make sure your clusters don't have overlapping IPv4 subnets, you will have to orchestrate those, as well.

So, just as with containers themselves, whatever advantages container networking has in some use cases, the cost is absurdly high for *all* users.

With Khutulun, we leave all the above up to you. If you're willing to pay the price, we give you container networking. If you can't or don't want to, we let you use the host networks. Supporting various networking solutions is indeed a central goal for Khutulun. Because it's built on graphs, with relationships being first-class citizens (see the next point), Khutulun provides a strong basis for modeling complex network topologies (including container networking). So, let's talk
about graphs.

### Relationships

Third point: There are no systemic relationships between entities in Kubernetes. The only obvious one is "ownership", to support garbage collection of children when an owner is deleted, and it's handled as a second-class citizen via metadata. This is a very unfortunate omission from Kubernetes's data model, because graphs are everywhere you look in the cloud. From modeling networking infrastructure (as noted above), through coordination between microservices (service meshes), to logical inter-dependencies for scalability and healing policies, all the way down to the cluster machine infrastructure itself.

Khutulun *starts* with graphs. Relationships (graph edges) get as much attention as nodes (graph vertexes). Moreover, relationships are dynamic, meaning that they are formed, dissolved, and reformed imperatively. An event, e.g. a node failure, can cause the whole graph to be updated on demand. If you store your data in a graph database (see next point), you can made good use of its ability to process complex path-following queries and gain a true topological view of your
deployments.

### Data Storage

Fourth, Kubernetes's choice of using just one storage engine, in particular [etcd](https://etcd.io/), was shortsighted. The idea of not having to deploy a separate storage service, and instead letting the cluster's nodes distribute the data among themselves, might seem convenient at first. But any attempt to treat etcd as a serious database ends in disappointment. It sorely lacks capabilities for orchestrating data in large clouds, features that are readibly
and reliably available in mature database solutions.

With Khutulun we intend to support all the popular databases, including those services hosted by public cloud providers. Graph databases are especially useful for large-scale deployments with complex topologies.

### Declarative Hell

Finally, Kubernetes is far too declarative. Extending it with code requires writing an "operator" (custom controller), creating at least one custom resource, and participating in the declarative dance with other controllers. Instead of being event driven, there is a constant hum of polling and updating. Just standing up a vanilla Kuberentes cluster, with no running pods, will show non-negligible CPU usage. Developers can attest to the pain of having to wait for controllers to kindly respond to updates, and when a chain of operators is involved, it can be prohibitively slow
and finicky.

As we point out in the Khutulun README, when everything is declarative some hard things become easy, but many easy things become very hard indeed. Khutulun is dual-paradigm, both declarative and imperative. Wherever possible, it provides a hook for you to insert your own custom code. You can respond to cloud and application events immediately instead of hoping that a criticial (possibly third-party) controller will do what it's supposed to.

### Summary: Is Kubernetes Fixable?

It's not easy to imagine Kubernetes growing out of its current set of problematic opinions. In a way, it's a victim of its own success: Now that it's being used in production environments, it can't afford to break critical deployments by fundamentally changing how it operates. It will surely continue to improve (e.g. there are efforts to allow for plugabble storage), but the basic design is what it is.

Khutulun is born from the necessity of going back to the drawing board, learning from successes and mistakes of the past decade, and building a system that's flexible enough that it can evolve with the times and needs.
