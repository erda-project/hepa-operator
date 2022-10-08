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
    go build -a -o manager main.go && \
    go build -a -o kong-cli tools/kong-cli/main.go


# retag from centos:7
FROM registry.cn-hangzhou.aliyuncs.com/dspo/centos:7

MAINTAINER zhongrun.czr@alibaba-inc.com

VOLUME /workspace
WORKDIR /
COPY --from=builder /workspace/manager /usr/local/bin/hapi-operator
COPY --from=builder /workspace/kong-cli /usr/local/bin/kong-cli
USER 65532:65532

ENTRYPOINT ["hapi-operator"]
