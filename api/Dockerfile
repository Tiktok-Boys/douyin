FROM golang:1.20.7-alpine AS builder

WORKDIR /build

COPY . .

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod download && \
    go build .

FROM scratch

COPY --from=builder /build/api /app/api

ENV ETCD_HOST="172.17.0.8" \
    ETCD_PORT=2379

EXPOSE 8080

ENTRYPOINT ["/app/api"]
