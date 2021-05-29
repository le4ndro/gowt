FROM golang:1.16-alpine3.13 AS builder

RUN apk add --no-cache git gcc musl-dev make

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0  

WORKDIR /build
COPY  . /build

RUN make vendor

RUN make build

FROM alpine:3.13

RUN apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=builder /build/out/bin/gowt .
COPY --from=builder /build/templates ./templates
COPY --from=builder /build/.env .

ENTRYPOINT ["./gowt"]