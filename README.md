# hepa-operator

The Hepa Operator is a Kubernetes Operator, implemented using the kubebuilder scaffolding tool.

Hepa Operator is designed to abstract the reverse proxy of a cluster API gateway and the policies attached to the HTTP interface. Currently supported and intended gateway products are:

- Aliyun MSE
- Kong

## Usage

Prerequisite: K8s must be installed first. You can install K8s using tools such as [KIND](https://sigs.k8s.io/kind) or Docker Desktop.

### Build and Deploy

```sh
make release-completelty
```

A `release.yml` file is generated in the project root, which contains all the K8s resources needed to install Hepa Operator (including CRDs, Namespace, Deployment, Role, etc.).

```sh
make apply
```

All K8s resources of Hepa Operator will be created or updated.

### Local Operation

```sh
make manifest; make install; make run
```

to run locally.

### Apply Samples

```sh
make apply-samples
```

### Uninstallation

```sh
kubectl delete --ignore-not-found -f release.yml
```
