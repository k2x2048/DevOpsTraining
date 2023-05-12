# Kubernetes - Deployments

Module: **Orchestration** </br>
Competency: is able to deploy stateless & statefull applications on k8s <br />
Duration: **2 days** </br>
Participant(s): **solo** </br>

## Introduction

When deploying an application on a Kubernetes cluster you will have to take into considerations multiple things, like for example:

How to manage the application configuration settings ? [ConfigMaps](https://kubernetes.io/docs/concepts/configuration/configmap/)<br />
How do you handle your applications storage ? [Volumes](https://kubernetes.io/docs/concepts/storage/)<br />
How do we make and kubernetes application highly availible to the outside world ? [Services, Load Balancing, and Networking](https://kubernetes.io/docs/concepts/services-networking/)<br />
What about Kubernetes security standards ? [Security](https://kubernetes.io/docs/concepts/security/) <br />

It absolutly necessary to have a good understanding of all these fundamental concept for you to be able deploy applications on k8s.

## The mission

Today we are going to learn about two different and very important concepts when deploying applications with k8s which are Stateless versus Stateless deployments.

Basicaly you will have to ask yourself the following question : <br>
Does the application I'm deploying need to keep its datas (the state)? if yes, it means that you will have to perform a Statefull deployment. On the other hand, if your application doesn't need to keep it's datas (state) it means that you will need a Stateless deployment.

In order to help you implement these two types of deployments, your mission is to follow the two following tutorials:

- Stateless: [Deploying PHP Guestbook application with Redis](https://kubernetes.io/docs/tutorials/stateless-application/guestbook/)
- Statefull: [Deploying WordPress and MySQL with Persistent Volumes](https://kubernetes.io/docs/tutorials/stateful-application/mysql-wordpress-persistent-volume/)

**Important note:**<br>
Before you start the Statefull tutorial, I highly recommand you to take a look at the [StatefulSet Basics](https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/)

<br /><br /><br />

![](https://res.cloudinary.com/practicaldev/image/fetch/s--L1mx3fun--/c_imagga_scale,f_auto,fl_progressive,h_720,q_auto,w_1280/https://dev-to-uploads.s3.amazonaws.com/i/bvc0q7azyffnp10liamy.jpeg)
