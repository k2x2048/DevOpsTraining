# Prometheus 101

Module: **Monitoring** </br>
Competency: is able to understand how to Prometheus works and how to integrate it into your projects] </br>
Duration: **2 days** </br>
Participant(s): **solo** </br>

## Introduction

In production software, things go wrong at lots of different levels, from servers running out of disk space, to denial of service attacks on the network, to misconfigured application releases. To keep your systems working, you need a monitoring and alerting solution which tracks metrics at all those levels over time so you can see trends emerging and you can take action before things get really bad. And that is what Prometheus gives you, It's a free open‑source monitoring toolkit.

## The mission

In today's challenge you will be asked to learn about every things you need to know to be able to work with prometheus in your projects. As usual when it comes to learning a new tool, you will have to discover what can the tool do for you by understanding some of the related concepts such as :

- [Data Model](https://prometheus.io/docs/concepts/data_model/)
- [Metric types](https://prometheus.io/docs/concepts/metric_types/)
- [Jobs and instances](https://prometheus.io/docs/concepts/jobs_instances/)

Once you've leaned about that you can proceed about the installation and configuration process. To start, I personnaly recommand a local install of the prometheus server using Docker. Once you have your Prometheus server up and running, you still have to learn about what we call "[Exporters](https://prometheus.io/docs/instrumenting/exporters/)". To start, I highly recommand you to install a NodeExporter on a either your local (VM or not) or a remote system. One other important thing with prometheus is it's [alerting system](https://prometheus.io/docs/alerting/latest/overview/) which is basicaly made of two main components, the [alerting rules](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/) and the [AlertManager](https://prometheus.io/docs/alerting/latest/alertmanager/). Aside from that, take a look at the [Expression Browser](https://prometheus.io/docs/visualization/browser/) which will allow you to quickly visualise collected data over.

Next, and this is a very important thing to learn when working with Prometheus in your project, is to learn about the Prometheus Query Language ([PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/)) which will all you to perform custom queries into the time series database to retrieve specific data metrics about all your monitored systems.

The last thing you will have to learn about, is how to monitor your Docker containers with Prometheus with the help of [cAdvisor](https://prometheus.io/docs/guides/cadvisor/).

## Complementary Resources

- [Prometheus Offical Documentation](https://prometheus.io/docs/introduction/overview/)
- [Getting Started with Prometheus](https://prometheus.io/docs/tutorials/getting_started/)
- [Understanding metrics types](https://prometheus.io/docs/tutorials/understanding_metric_types/)
- [Alerting based on metrics](https://prometheus.io/docs/tutorials/alerting_based_on_metrics/)
- [Pluralsight - Getting Started with Prometheus](https://app.pluralsight.com/library/courses/getting-started-prometheus/table-of-contents)
- [Pluralsight - Alerting on Issues with Prometheus Alertmanager](https://app.pluralsight.com/library/courses/alerting-issues-prometheus-alertmanager/table-of-contents)

<br />
<br />
<br />
<br />

![Briefing's GIF](https://1vi3uw2r71ig39fow2sdrab1-wpengine.netdna-ssl.com/wp-content/uploads/2021/11/Screen-Shot-2021-11-30-at-2.32.19-PM.png)
