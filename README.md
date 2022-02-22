# ToDo
- [front] write test code
- [front] apply gRPC
- [back] apply gRPC
# ToDo Optional
- [front] improve ui/ux
- [back] peformance tuning

# Flow
1. Install `docker`, `kubernetes` and `dapr`
    - https://docs.docker.com/get-docker/
    - https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
    - https://docs.dapr.io/getting-started/install-dapr-cli/
2. Initialize dapr on kubernetes cluster
    - `dapr init --kubernetes --wait`
    - [more details](https://github.com/dapr/quickstarts/tree/v1.6.0/hello-kubernetes)
3. Build docker images (images were not pushed to DockerHub)
    - `docker build -t othello/front:prod -f front/prod.Dockerfile .`
    - `docker build -t othello/board:prod -f board/prod.Dockerfile .`
    - `docker build -t othello/cp:prod -f cp/prod.Dockerfile .`
4. Apply manifests (skip the redis settings for dapr as we don't use it)
    - `kubectl apply -f manifests`
    - `kubectl port-forward service/othello-front 3000:80`

# Development for application without dapr
1. Install `docker` and `docker-compose`
    - https://docs.docker.com/get-docker/
    - https://docs.docker.com/compose/install/
2. Uncomment entrypoint and run below command
    - `docker-compose up -d --build`


# Related Repository
- https://github.com/suragnair/alpha-zero-general
