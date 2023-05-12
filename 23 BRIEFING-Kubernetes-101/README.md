# Kubernetes 101

Module: **Orchestration** </br>
Competency: is able to explain kubernetes architechture & componants <br />
Duration: **2 days** </br>
Participant(s): **solo** </br>

## Introduction

Kubernetes (K8s) is simply a container orchestration framework. It means is that K8s is a system designed to automate the lifecycle of containerized applications — from predictability, scalability to availability.

The Reason behind the rise and need for K8s stems from the increasing use of microservices, away from traditional monolithic-type applications. As a result, containers provide the perfect host for these individual microservices as containers manage dependencies, are independent, OS-agnostic and ephemeral, amongst other benefits.

Complex applications that have many components are often made up of hundreds or even thousands of microservices. Scaling these microservices up while ensuring availability is an extremely painful process if we were to manage all these different components using custom-written programs or scripts, resulting in the demand for a proper way of managing these components and this is where container orchestration tools like K8s come into play.

## The mission

Today your mission will be understand the core concepts of Kubernetes, the architecture of the system, the problems it solves, and the model that it uses to handle containerised deployments and scaling.

After diving into the theory, you should be able to explain the architecture of a k8s cluster, list different constituents of a master node and their respective roles (API Server, Scheduler, Controller Manager, etcd). As well as what makes a worker node a worker node (Kublet, container runtime, kube-proxy,...). After focusing on the architecture, you will now focus on the different components of a k8s system, such as pods, services, deployments, ingress, ingress controllers, ConfigMaps, Secrets, Persistent Volumes, Persistent Volumes Claims, DeamonSets, ReplicaSets, StatefullSets, etc...

Remember that you should be able to explain what each of these components does as if you are in an interview for a job.

Also, make sure you do not forget to learn about the tooling (kubectl, kubeadm, minikube, etc...)

Finally, I suggest you get your hands dirty with a few [labs on the basic k8s modules](https://kubernetes.io/docs/tutorials/kubernetes-basics/).

## Complementary Resources

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [Pluralsight - Getting Started with Kubernetes](https://app.pluralsight.com/library/courses/kubernetes-getting-started/table-of-contents)
- [Medium - 8 Best Free Kubernetes Courses for Beginners in 2022](https://medium.com/javarevisited/7-free-online-courses-to-learn-kubernetes-in-2020-3b8a68ec7abc)
- [Edx - Introduction to Kubernetes](https://www.edx.org/course/introduction-to-kubernetes?source=aw&awc=6798_1660107853_6f3cd051bc758e5aca33cb66504e75cb&utm_source=aw&utm_medium=affiliate_partner&utm_content=text-link&utm_term=631878_javarevisited)
  <br />
  <br />
  <br />

![](https://starecat.com/content/wp-content/uploads/this-is-a-ship-shipping-ship-shipping-shipping-ships.jpg)
