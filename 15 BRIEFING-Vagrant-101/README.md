# Vagrant

Module: **Virtualization** </br>
Competency: [is able to use vagrant to spin up VMs and containers](./evaluation.md) </br>
Duration: **5 days** </br>
Participant(s): **solo** </br>
Output: **pull request**

## The mission

With `docker` and its `Dockerfile` you had a glimpse of how useful configuration files can be to manage infrastructures. Allowing you to spin up containers with **your desired specifications** and more importantly in a **reproducible** manner. Wouldn't it be great if there was such a tool for virtual machines?! You guessed it, there is and it's called `vagrant`.

First thing first though, what's `vagrant` exactly? Well to answer that question there is nothing better than to **follow some tutorials**! Below is a curated list, but feel free to search and learn from other sources.

* official [getting started](https://learn.hashicorp.com/collections/vagrant/getting-started) by _HashiCorp_
* introduction to the [base commands](https://github.com/Zuehlke/vagrant-chef-exercises) (_only the "Vagrant Introduction" part_)
* article covering [the fundamentals](https://hackernoon.com/devops101-vagrant-6737c8c29904)

Now that you've read it all, let's practice! If not done already, install vagrant (_and all its requirements_) then try to solve the exercises below.

1. Follow this first "_[amuse-bouche](https://github.com/dwyl/learn-vagrant)_" to get acquainted with the tech.
2. You might already have read it somewhere, but `vagrant` has a `docker` provisioner (_backend_) to create **Docker containers**. Try it out with this [exercise](https://github.com/sighupio/interview-basic-docker-vagrant)!
3. Finally, and as a main course, try to deploy the [COGIP website](https://github.com/becodeorg/BXL-Lovelace-7.34/issues/94) using `vagrant`. There are multiple ways to achieve this deployment (_VMs_, _containers_, _containers in VMs_, _etc..._), but your solution should at the very least have the **website running** and a **reverse proxy** to manage an _HTTPS connection_.

> **NOTE**: If you're feeling adventurous, you can also try to setup a [load balancer](https://en.wikipedia.org/wiki/Load_balancing_(computing)).

## Complementary Resources

* [Documentation](https://www.vagrantup.com/docs)
* [Awesome List of Resources](https://github.com/iJackUA/awesome-vagrant)
* [Cloud Architecture Lab](https://hackmd.io/@dme26/ByFJ5w7xr)
* [DevOps implementation](https://www.rohitsalecha.com/post/practical_devops_infrastructure_as_code_vagrant_ansible_docker/)
* [Vagrant and IAC](https://hackernoon.com/devops101-itinfrastructure-54337d6a148b)

## Final words

When it comes to IaC (_infrastructure as code_), `vagrant` serves as a primordial step to achieve a **reproducible setup**, but further combined with tools such as `ansible`, `terraform` it allows for amazing levels of automation.

![Briefing's GIF](https://c.tenor.com/fQlt4OtCfXsAAAAd/i-have-the-power-power.gif)
