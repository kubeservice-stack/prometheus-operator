ARG ARCH="amd64"
ARG OS="linux"

# Build the manager binary
FROM golang:1.20.3-alpine as builder
RUN apk add --no-cache gcc musl-dev libc6-compat

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd/ cmd/
COPY pkg/ pkg/

RUN go mod tidy
RUN go mod vendor

# Build
RUN CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -ldflags "-linkmode external -extldflags -static" -o prometheus-config-reloader ./cmd/prometheus-config-reloader/

FROM quay.io/prometheus/busybox-${OS}-${ARCH}:latest

COPY --from=builder /workspace/prometheus-config-reloader /bin/pingmesh-agent

USER nobody

ENTRYPOINT ["/bin/prometheus-config-reloader"]
