
# blue-discount

The project uses the following packages and tools:

[Docker](https://www.docker.com/)
[Kubernetes](https://kubernetes.io/)
[Tilt](https://tilt.dev/) (development)
[Gingko](https://github.com/onsi/ginkgo) (BDD for Go)
[Gomega](https://github.com/onsi/gomega)
[Gorm](https://gorm.io/index.html)
[Viper](https://github.com/spf13/viper)

## Running locally

The project was created to run in Kubernetes. You dont need to install dependencies, such as Postgres, in your machine, because all k8s objects are used by Tilt during development. Tilt is a tool to setup the project in k8s with support to reload on changes. You get the docs to install it [here](https://docs.tilt.dev/install.html), and to create a local k8s cluster is recommended [Kind](https://github.com/tilt-dev/kind-local).

To interact with the cluster:

```
$ kubectl cluster-info --context kind-kind
```

Create a dev namespace:

```
$ kubectl create namespace dev
```

To run the project:

```
$ tilt up
```

After that, you can access the service on http://localhost:8000.

## Testing

To test the service:

```
$ make test
```

To test with coverage:

```
$ make test-cov
```
