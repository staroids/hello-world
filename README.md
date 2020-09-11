# Hello world example on Staroid

Minimal example project running on [Staroid](https://staroid.com).

## Run on Staroid

[![Run](https://staroid.com/api/run/button.svg)](https://staroid.com/api/run)

Check [here](https://docs.staroid.com/project/run_button.html) to know more about 'Run on staroid' button.

## Run locally for development

Open a terminal and run

```bash
$ skaffold dev --port-forward
```

and browse http://localhost:8080.

[skaffold](https://skaffold.dev) CLI and a Kubernetes (e.g. [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)) are required to run locally.

## Fork, edit and have your own version.

You can fork this repository and [connect](https://staroid.com/projects/settings) from Staroid. Edit code and release your own version of the project.

## Files

| file | description |
| ---- | --------- |
| [Dockerfile](https://github.com/staroids/hello-world/blob/master/Dockerfile) | Dockerfile to build container image |
| [skaffold.yaml](https://github.com/staroids/hello-world/blob/master/skaffold.yaml) | [Skaffold](https://skaffold.dev) configuration to execute build and deploy |
| [staroid.yaml](https://github.com/staroids/hello-world/blob/master/staroid.yaml) | [Staroid configuration](https://docs.staroid.com/references/staroid_yaml.html)
| [k8s-deployment.yaml](https://github.com/staroids/hello-world/blob/master/k8s-deployment.yaml) | Kubernetes resource file to create [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) controller |
| [k8s-service.yaml](https://github.com/staroids/hello-world/blob/master/k8s-service.yaml) | Kubernetes Resource file to deploy [Service](https://kubernetes.io/docs/concepts/services-networking/service/) |
| [main.go](https://github.com/staroids/hello-world/blob/master/main.go) | Application source code |


For more sophisticated examples, check https://docs.staroid.com/examples/index.html