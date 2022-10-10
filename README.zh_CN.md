# hepa-operator

Hepa Operator 是一个 Kubernetes Operator, 在实现过程中使用了 kubebuilder 脚手架工具.

Hepa Operator 旨在对集群 API 网关的反向代理及附着于 HTTP 接口上的策略进行抽象. 目前已支持和打算支持的网关产品有:

- 阿里云 MSE
- Kong

## 使用方式

前提条件: 首先要安装 K8s. 你可以使用 [KIND](https://sigs.k8s.io/kind) 或 Docker Desktop 等工具安装 K8s.

### 构建与部署

```sh
make release-completelty
```

会在项目根目录生成一个 `release.yml` 的文件, 它包含了安装 Hepa Operator 所需的所有 K8s 资源 (包含 CRDs, Namespace, Deployment, Role 等).

```sh
make apply
```

会创建或更新 Hepa Operator 所有的 K8s 资源.

### 本地运行

```sh
make manifest; make install; make run
```

即可本地运行.

### 运行示例

```sh
make apply-samples
```

或

```sh
make get-samples
```

输出(`apply` 并 `get` 相关示例资源):

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


### 卸载

```sh
kubectl delete --ignore-not-found -f release.yml
```

### 它是如何工作的 ?

```yaml
apiVersion: hepa.erda.cloud/v1
kind: ConfigZone
metadata:
  name: hapi-operator-sample
  namespace: hapi-operator-sample
spec:                                               # spec 描述一些配置信息, 它可以被 Hapi 引用
  policy:
    auth:
      authType: hmac-auth
      switch: true
    safetyIP:
      blackListSourceRange: ""
      domainBlackListSourceRange: ""
      domainWhiteListSourceRange: ""
      ipType: x-peer-ip
      keyRateLimitingValue: 10 query_per_second
      switch: true
      whiteListSourceRange: 123.45.67.1/16,10.10.10.10
```
```yaml
apiVersion: hepa.erda.cloud/v1
kind: Hapi
metadata:
  name: hapi-sample
  namespace: hapi-operator-sample
  labels:
    "configZone": "hapi-operator-sample"            # ## 引用的 ConfigZone 
    "packageId": "c82396e5fc13ef7bbf6bc078502a21e4" # ## 用户自定义的标签
spec:                                               # spec 描述反向代理的规则
  hosts:                                            # # 对外的域名列表
    - hapi-sample.mse-daily.terminus.io
  path: /search                                     # # 路由路径
  backend:                                          # # backend 描述后端转发规则
    redirectBy: url                                 # ## redirectBy 描述按 url 还是 service 转发
    serviceName: go-httpbin                         # ## 如果按 service 转发, 则反向代理到该 namespace 的 go-httpbin:80/s
    servicePort: 80
    upstreamHost: baidu.com                         # ## 如果按 url 转发, 则反向代理到 baidu.com/s
    rewriteTarget: /s
  policy:                                           # # policy 描述路由策略
    auth:                                           # ## auth 描述认证策略
      authType: sign-auth
      global: false                                 # ## global=true 时引用 configZone 的对应策略的配置
      switch: true                                  # ## switch=true 表示启用该策略(global=true 时以 configZone 中该策略的配置为准)
    safetyIP:
      blackListSourceRange: ""
      domainBlackListSourceRange: ""
      domainWhiteListSourceRange: ""
      ipType: x-peer-ip
      keyRateLimitingValue: 12 query_per_second
      global: true
      switch: false
      whiteListSourceRange: 123.45.67.1/18
```

通过定义一个名为 `Hapi` 的 CRD 来表示反向代理规则, 它描述了一对转发关系以及附着于这对转发关系上的策略.

Hepa-Operator 监听了这个 CRD, 并根据其配置信息来控制一些 K8s API 对象(e.g. Ingress, Service) 以及一些外部资源(e.g. Aliyun MSE Gateway Openapi, Kong Admin API) 来实现反向代理.

```sh
% kubectl -n hapi-operator-sample get czr,hapi,ing,svc
NAME                                              SCENE   HOSTS   HAPI_COUNT   POLICIES              PHASE
configzone.hepa.erda.cloud/hapi-operator-sample                   1            ["AUTH","SAFETYIP"]   OK

NAME                               ENDPOINT                                   REDIRECTTO    POLICIES              PHASE
hapi.hepa.erda.cloud/hapi-sample   hapi-sample.mse-daily.terminus.io/search   baidu.com/s   ["SAFETYIP","auth"]   OK

NAME                                    CLASS   HOSTS                               ADDRESS                   PORTS     AGE
ingress.networking.k8s.io/hapi-sample   mse     hapi-sample.mse-daily.terminus.io   **.**.**.**,**.**.**.**   80, 443   9d

NAME                           TYPE           CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
service/external-hapi-sample   ExternalName   <none>        baidu.com     80/TCP,443/TCP   10d
```

其中 `czr` 是 ConfigZone 的短名称(Config Zone Reference).

ConfigZone 的打印列 `HAPI_COUNT` 表示引用该配置的 HAPI 的计数; `POLICIES` 表示该实例上启用的策略列表.

Hapi 的打印列 `ENDPOINT` 和 `REDIRECTTO` 表示一对路由关系, 如示例中表示将对 "hapi-sample.mse-daily.terminus.io/search" 的请求转发到 "baidu.com/s";
`POLICIES` 表示该路由上启用的策略, 大写表示引用全局策略, 驼峰式表示本地策略.

它是如何实现这个转发关系的呢 ? 当网关采用 Aliyun MSE 时, 对于按 service 转发, Hepa Operator 会创建一个 Ingress 反向代理到这个 Service; 对于按 url 转发(如示例),  Hepa Operator 会先额外创建一个 ExternalName 类型的 Service, 其 ExternalIP 是目标域名, 再创建一个 Ingress 反向代理到这个 Service.
当网关采用 Kong 时, 它会创建一个 Ingress 反向代理到 Kong 的服务, 再创建 Kong Service, Route 等对象反向代理到真正的后端服务.

> **implement by Aliyun MSE**
![implement by Aliyun MSE](http://terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2022/10/10/bc9ef280-e3b1-4921-974a-0bbbcb45e18d.png)
![legend](http://terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2022/10/10/aa1a2a79-2bb3-4887-8557-1ed70b70f604.png)

> **implement by Kong**
![implement by Kong](http://terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2022/10/10/7a48f4f6-479f-4ab1-b56d-7338eb0fb6a4.png)
![legend](http://terminus-paas.oss-cn-hangzhou.aliyuncs.com/paas-doc/2022/10/10/aa1a2a79-2bb3-4887-8557-1ed70b70f604.png)