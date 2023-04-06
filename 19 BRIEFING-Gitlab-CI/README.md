# Gitlab - CI

Type of challenge: **learning** </br>
Duration: **2 day** </br>
Team challenge : **solo**

In today's challenge, we are going to learn how Continuous Integration can be achieved when working with Gitlab in your projects. 

Learning Objectives:

- Gitlab-CI structure and synthax
- Gitlab Container Registry
- Gitlab Runner
  
## The mission

As a DevOps engineer working at ACME a consulting compagny that helps their clients to optimize their software development process. You a currently working on two different projects ([Project 1](https://github.com/g0t4/jgsu-spring-petclinic) / [Project 2](https://github.com/lewisemm/converter)) for which you will have to implement a CI/CD pipeline in Gitlab. Make sure that the application build process (build/test/release/deploy) can be done each time a developer makes a commit on the main branch. Also, each project has to be containerized (which is not the case for now).

For the release step, you will have to push your docker image to Gitlab's Container Registry.

### Quick note

- Use you VM to install your Gitlab Runner.
- Deployement has to be done one your VM.
- Use gitlab.com and not a self hosted version of gitlab.

# Usefull links

- [Gitlab-CI](https://docs.gitlab.com/ee/ci/)
- [Gitab Runner](https://docs.gitlab.com/runner/)
- [Gitlab Container registry](https://docs.gitlab.com/ee/user/packages/container_registry/)