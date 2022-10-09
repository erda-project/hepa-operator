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

or 

```sh
make get-samples
```

output:

```shell
[root@node-172016174045 hepa-operator]% make get-samples
kubectl apply -f config/samples/_v1_hapi.yaml
namespace/hapi-operator-sample unchanged
deployment.apps/go-httpbin unchanged
service/go-httpbin unchanged
configzone.hepa.erda.cloud/hapi-operator-sample unchanged
hapi.hepa.erda.cloud/hapi-sample unchanged

kubectl -n hapi-operator-sample get czr,hapi,ing,svc,deploy,pod
NAME                                              SCENE   HOSTS   HAPI_COUNT   POLICIES              PHASE
configzone.hepa.erda.cloud/hapi-operator-sample                   1            ["AUTH","SAFETYIP"]   OK

NAME                               ENDPOINT                              REDIRECTTO    POLICIES              PHASE
hapi.hepa.erda.cloud/hapi-sample   hapi-sample.mse-daily.terminus.io/s   baidu.com/s   ["SAFETYIP","auth"]   OK

NAME                                    CLASS   HOSTS                               ADDRESS                   PORTS     AGE
ingress.networking.k8s.io/hapi-sample   mse     hapi-sample.mse-daily.terminus.io   **.**.**.**,**.**.**.**   80, 443   9d

NAME                           TYPE           CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
service/external-hapi-sample   ExternalName   <none>        baidu.com     80/TCP,443/TCP   10d
service/go-httpbin             ClusterIP      **.**.**.**   <none>        80/TCP           37d

NAME                         READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/go-httpbin   1/1     1            1           37d

NAME                              READY   STATUS    RESTARTS   AGE
pod/go-httpbin-68fdb87875-g7b8f   1/1     Running   0          11d
```


### Uninstallation

```sh
kubectl delete --ignore-not-found -f release.yml
```
