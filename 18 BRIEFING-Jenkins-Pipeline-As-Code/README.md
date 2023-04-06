# Jenkins - Pipeline As Code

Type of challenge: **learning** </br>
Duration: **2 day** </br>
Team challenge : **solo**

Up until now we have been creating and configuring jobs on the jenkins web interface with the help of Freestyle Projects. But, around december 2014, the Jenkins project introduced scripted synthax pipelines ! And, around february 2017, declarative pipelines made their entrance !

Now, you may be asking what is a jenkins pipeline ? Well jenkins pipeline is a suite of plugins that makes it possible to create the definition of a job within a file.

You may be familiar with freestyle jobs, maven jobs or other types of jobs. But, pipeline is the only one that makes it possible to be able to save the definition of your job to an scm such as git.

Today you're gonna learn you what jenkins file is, how to create it and what the syntax looks like. So, what is jenkins file ? What are the differences between scripted pipelines and declarative pipelines ?

## The mission

As a newly hired DevOps Engineer at a compagny called EvilCorp. The company is spcialized in cipher algorithm development implemented in Java.

Given this [Github Repo](https://github.com/YoussF/caesar-cipher), you have been asked to make sure that the application build process (build/test/release/deploy) can be done each time a developer makes a commit on the main branch using Jenkins declarative pipeline living inside a 'jenkinsfile' inside the repo. Please note that this project as already been setup with [Gradle](https://gradle.org/) as the build/test tool. At least for the testing part. Infact, currently, if a developer tries to build the project, the resultant 'jar' artifact will produce an error when executed.

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

- Install the necessary tools (JVM, gradle, jenkins...)
- Make a new Github repository with the content of the original [Repo](https://github.com/YoussF/caesar-cipher)
- Try to run the build and the tests manually to make sure everything is ok
- Create a Jenkinsfile, and write all the necessary steps, then push it to the remote repo.
- Create a pipeline in jenkins...

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

- [Jenkins Documentation](https://www.jenkins.io/doc/)
- [Gradle build tool](https://gradle.org/)
