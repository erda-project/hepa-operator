# docker build -t ${REGISTRY}/hapi-buildkit:1.0.0-$(date "+%Y%m%d-%H%M%S")-$(git rev-parse --short HEAD 2>/dev/null) -f dockerfiles/toolkit/Dockerfile .

FROM registry.cn-hangzhou.aliyuncs.com/dspo/centos:7

MAINTAINER zhongrun.czr@alibaba-inc.com

# args
ARG GO_VERSION=1.19.1
ARG KUBEBUILDER_VERSION=3.6.0
ARG KUSTOMIZE_VERSION=v4.5.7
ARG GO_PKG=go$GO_VERSION.linux-amd64.tar.gz
ARG KUSTOMIZE_PKG="kustomize_"$KUSTOMIZE_VERSION"_linux_amd64.tar.gz"

# install wget
RUN yum install -y curl \
    make \
    gcc \
    docker \
    dnf

# install git
RUN dnf install -y git-all

# install golang
# curl -L -o $GO_PKG https://dl.google.com/go/$GO_PKG
# e.g. curl -L -o https://dl.google.com/go/go1.19.1.linux-amd64.tar.gz
ARG GO_URL=https://dl.google.com/go/$GO_PKG
RUN curl -sLO https://dl.google.com/go/$GO_PKG && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf $GO_PKG
ENV PATH $PATH:/usr/local/go/bin
ENV GOBIN /usr/local/go/bin
ENV GOPROXY https://goproxy.cn,direct

# curl -L -o kubebuilder.linux.amd64 https://go.kubebuilder.io/dl/$KUSTOMIZE_VERSION/linux/amd64
# e.g. curl -L -o kubebuilder.linux.amd64 https://go.kubebuilder.io/dl/3.6.0/linux/amd64
ARG KUBEBUILDER_URL=https://go.kubebuilder.io/dl/$KUBEBUILDER_VERSION/linux/amd64
RUN curl -L -o /usr/local/bin/kubebuilder https://go.kubebuilder.io/dl/$KUSTOMIZE_VERSION/linux/amd64 && \
    chmod +x /usr/local/bin/kubebuilder

# curl -L -o $KUSTOMIZE_PKG https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/$KUSTOMIZE_VERSION/$KUSTOMIZE_PKG
# e.g. curl -L -o kustomize_v4.5.7_linux_amd64.tar.gz https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v4.5.7/kustomize_v4.5.7_linux_amd64.tar.gz
ARG KUSTOMIZE_URL=https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/$KUSTOMIZE_VERSION/$KUSTOMIZE_PKG
RUN curl -sLO $KUSTOMIZE_URL && \
    tar -xzf $KUSTOMIZE_PKG && \
    mv kustomize /usr/local/bin/kustomize && \
    chmod +x /usr/local/bin/kustomize
