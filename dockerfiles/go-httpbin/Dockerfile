# docker build -t ${REGISTRY}/go-httpbin:$(date "+%Y%m%d-%H%M%S")-$(git rev-parse --short HEAD 2>/dev/null) -f dockerfiles/toolkit/Dockerfile .

# Build the manager binary
# retag from golang:1.19
FROM registry.cn-hangzhou.aliyuncs.com/dspo/golang:1.19 as builder

WORKDIR /workspace

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go mod tidy && \
    go build -a -o go-httpbin tools/go-httpbin/main.go


# retag from centos:7
FROM registry.cn-hangzhou.aliyuncs.com/dspo/centos:7

MAINTAINER zhongrun.czr@alibaba-inc.com

VOLUME /workspace
WORKDIR /
COPY --from=builder /workspace/go-httpbin /usr/local/bin/go-httpbin

USER 65532:65532

ENTRYPOINT ["go-httpbin"]