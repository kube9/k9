FROM golang:1.9 as builder
RUN mkdir -p /go/src/github.com/kube9/k9/
WORKDIR /go/src/github.com/kube9/k9/
COPY . .
RUN make k9server-build

FROM alpine:latest as final
WORKDIR /
COPY --from=builder /go/src/github.com/kube9/k9/bin/k9server .
CMD ["./k9server"]