docker run -d --name 'gitlab-runner-docker' --restart always \
  -v gitlab-runner-config:/etc/gitlab-runner \
  -v /var/run/docker.sock:/var/run/docker.sock \
  gitlab/gitlab-runner:latest
  
docker run -d --name 'gitlab-runner-docker' --restart always \
  -v gitlab-runner-config:/etc/gitlab-runner \
  -v /var/run/docker.sock:/var/run/docker.sock \
  gitlab/gitlab-runner:alpine

docker run --rm -v gitlab-runner-config:/etc/gitlab-runner gitlab/gitlab-runner register \
  --non-interactive \
  --executor "docker" \
  --docker-image "docker:20.10.16" \
  --url "https://gitlab.com/" \
  --registration-token "GR1348941_5JuMYB8QhMuKjix7yo1" \
  --description "docker-runner" \
  --maintenance-note "Free-form maintainer notes about this runner" \
  --tag-list "docker,linux,shell" \
  --run-untagged="true" \
  --locked="false" \
  --access-level="not_protected"
