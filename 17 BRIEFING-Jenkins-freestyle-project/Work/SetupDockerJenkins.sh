# Create a bridge network in Docker
docker network create jenkins

# download and run the docker:dind In order to execute Docker commands inside Jenkins nodes
docker run \
  --name jenkins-docker \
  --rm \
  --detach \
  --privileged \
  --network jenkins \
  --network-alias docker \
  --env DOCKER_TLS_CERTDIR=/certs \
  --volume jenkins-docker-certs:/certs/client \
  --volume jenkins-data:/var/jenkins_home \
  --publish 2376:2376 \
  docker:dind \
  --storage-driver overlay2


# Customise official Jenkins Docker image.
# -> Create Dockerfile with the following content:
# FROM jenkins/jenkins:2.375.3
# USER root
# RUN apt-get update && apt-get install -y lsb-release
# RUN curl -fsSLo /usr/share/keyrings/docker-archive-keyring.asc \
#   https://download.docker.com/linux/debian/gpg
# RUN echo "deb [arch=$(dpkg --print-architecture) \
#   signed-by=/usr/share/keyrings/docker-archive-keyring.asc] \
#   https://download.docker.com/linux/debian \
#   $(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker.list
# RUN apt-get update && apt-get install -y docker-ce-cli
# USER jenkins
# RUN jenkins-plugin-cli --plugins "blueocean docker-workflow"


# Build a new docker image from this Dockerfile and assign the image a meaningful name.
docker build -t myjenkins:2.393-jdk17 .

# Run your own image as a container in Docker
  docker run \
  --name jenkins-petclinic \
  --restart=on-failure \
  --detach \
  --network jenkins \
  --env DOCKER_HOST=tcp://docker:2376 \
  --env DOCKER_CERT_PATH=/certs/client \
  --env DOCKER_TLS_VERIFY=1 \
  --publish 8080:8080 \
  --publish 50000:50000 \
  --volume jenkins-data:/var/jenkins_home \
  --volume jenkins-docker-certs:/certs/client:ro \
  --volume ./home:/home \
  --env JAVA_OPTS="-Dhudson.plugins.git.GitSCM.ALLOW_LOCAL_CHECKOUT=true" \
  myjenkins:2.393-jdk17

sleep 10

#print the password at console.
sudo docker exec jenkins-docker cat /var/jenkins_home/secrets/initialAdminPassword

# Browse to http://localhost:8080 (or whichever port you configured for Jenkins when installing it) and wait until the Unlock Jenkins page appears.


# Accessing the Docker container
# docker exec -it jenkins-petclinic bash

# Accessing the Docker logs
# docker logs jenkins-petclinic
