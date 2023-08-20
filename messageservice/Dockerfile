FROM golang:1.20.7-alpine AS builder

WORKDIR /build

COPY . .

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod download && \
    go build .

FROM scratch

COPY --from=builder /build/messageservice /app/messageservice

ENV MYSQL_HOST="kasperxms.xyz" \
    MYSQL_PORT=3306 \
    MYSQL_USERNAME="tiktok" \
    MYSQL_PASSWORD="TikTokBoys123" \
    MYSQL_DATABASE="tiktok" \
    ETCD_HOST="172.17.0.8" \
    ETCD_PORT=2379

EXPOSE 8666

ENTRYPOINT ["/app/messageservice"]
