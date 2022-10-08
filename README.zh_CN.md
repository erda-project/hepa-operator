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

### 卸载

```sh
kubectl delete --ignore-not-found -f release.yml
```