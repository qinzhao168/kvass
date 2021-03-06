# Kvass

  [![Go Report Card](https://goreportcard.com/badge/github.com/tkestack/kvass)](https://goreportcard.com/report/github.com/tkestack/kvass)  [![Build](https://github.com/tkestack/kvass/workflows/Build/badge.svg?branch=master)]()   [![codecov](https://codecov.io/gh/tkestack/kvass/branch/master/graph/badge.svg)](https://codecov.io/gh/tkestack/kvass)

------

Kvass provides a solution for Prometheus sharding, which uses Sidecar to generate new config only use "static_configs" for Prometheus scraping according to targets assigned from Coordinator.

A Coordinator manage all shards  and assigned targets to each of them。
Thanos (or other storage solution) is used to provide a global data view。

![image-20201123224137790](./README.assets/image-20201123224137790.png)

* **Coordinator** loads origin config file and do all prometheus service discovery, for every target, Coordinator do "relabel_configs" and explore it's series scale and try assgin it to Sidecar according to Head Block Series of all Prometheus instance.
* **Sidecar** receives targets and generate new config file for prometheus only use "static_configs"。
------

## Feature

* Tens of millions series supported (thousands of k8s nodes)
* One configuration file
* Dynamic scaling
* Sharding according to the actual target load instead of label hash
* Multiple replicas supported

## Quick start 

clone kvass to local 

> git clone https://github.com/tkestack/kvass

install example (just an example with testing metrics)

> Kubectl create -f ./examples

you can found a Deployment named "metrics" with 6 Pod, each Pod will generate 10045 series (45 series from golang default metrics) metircs。

we will scrape metrics from them。

![image-20200916185943754](./README.assets/image-20200916185943754.png)

the max series each Prometheus Shard can scrape is a flag of Coordinator Pod.

in the example case we set to 30000.

> ```
> --shard.max-series=30000
> ```

now we have 6 target with 60000+ series  and each Shard can scrape 30000 series，so need 3 Shard to cover all targets.

Coordinator  automaticly change replicate of Prometheus Statefulset to 3 and assign targets to them.

![image-20200916190143119](./README.assets/image-20200916190143119.png)

only 20000+ series in prometheus_tsdb_head of one Shard

![image-20200917112924277](./README.assets/image-20200917112924277.png)

but we can get global data view use thanos-query

![image-20200917112711674](./README.assets/image-20200917112711674.png)

## Multiple replicas

Coordinator use label selector to select shards StatefulSets, every StatefulSet is a replica, Kvass puts together Pods with same index of different StatefulSet into one Shards Group.

> --shard.selector=app.kubernetes.io/name=prometheus

## Suggestion for flags

The memory useage of every Prometheus is associated with the max head series.

The recommended "max series" is 750000, set  Coordinator flag

> --shard.max-series=750000

The memory limit of Prometheu with 750000 max series is 8G.

## Build binary

> git clone https://github.com/tkestack/kvass
>
> cd kvass
>
> make

## License
Apache License 2.0, see [LICENSE](./LICENSE).

