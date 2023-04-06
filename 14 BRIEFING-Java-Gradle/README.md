# Java - Building with Gradle

Type of challenge: **learning** </br>
Duration: **1 day** </br>
Team challenge : **solo**

Up until now we have been we have been working with several languages and framework within our projects. From PHP to NodeJS and more recently Java. But when it comes to working with java projects there are many way a project can be built depending on the build tool that is used with the project. We already covered the Maven build tool in a previous project, in this project we are going to explore how to work with the Gradle build tool. 


## The mission
As a newly hired DevOps Engineer at a compagny called EvilCorp. The company is spcialized in cipher algorithm development implemented in Java. 

Given this [Github Repo](https://github.com/YoussF/caesar-cipher), you have been asked to make sure that the application build process (build/test/release/deploy). Please note that this project as already been setup with [Gradle](https://gradle.org/) as the build/test tool. At least for the testing part. Infact, currently, if a developer tries to build the project, the resultant 'jar' artifact will produce an error when executed. 

```
no main manifest attribute, in build/libs/caesars-cipher.jar
```

This error is due to the lack of configuration within the build.gradle file. In order to fix the error above you have to search the documentation to see how you can tell gradle which java class he has to associate with the produced jar file. This is the output of a correctly running app.

```
Message: in code we trust
Offset: 1
Ciphered message: jo dpef xf usvtu
```

To help you accomplish this task, here are some of the steps you should go through:

- Install the necessary tools (JVM, gradle,...)
- Make a new Github repository with the content of the original [Repo](https://github.com/YoussF/caesar-cipher)
- Try to run the build and the tests manually to make sure everything is ok
- Write the corresponding Dockerfile and run the application localy from docker.
- Push to docker image to DockerHub


Important: this project is known to be running within the following environment:

```
------------------------------------------------------------
Gradle 7.3.3
------------------------------------------------------------

Build time:   2021-12-22 12:37:54 UTC
Revision:     6f556c80f945dc54b50e0be633da6c62dbe8dc71

Kotlin:       1.5.31
Groovy:       3.0.9
Ant:          Apache Ant(TM) version 1.10.11 compiled on July 10 2021
JVM:          11.0.13 (Ubuntu 11.0.13+8-Ubuntu-0ubuntu1.20.04)
OS:           Linux 5.4.72-microsoft-standard-WSL2 amd64
```

![](https://media.makeameme.org/created/one-does-not-5bd659f3aa.jpg)

# Usefull links
- [Gradle build tool](https://gradle.org/)