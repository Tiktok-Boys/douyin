FROM golang:1.20.7-alpine AS builder

WORKDIR /build

COPY . .

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod download && \
    go build .

FROM scratch

COPY --from=builder /build/messageservice /app/messageservice

ENV MYSQL_HOST="172.17.0.4" \
    MYSQL_PORT=3306 \
    MYSQL_USERNAME="root" \
    MYSQL_PASSWORD="tiktok" \
    MYSQL_DATABASE="tiktok"

EXPOSE 8666

ENTRYPOINT ["/app/messageservice"]
