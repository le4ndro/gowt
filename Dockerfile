FROM golang:1.16-alpine as builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0  

WORKDIR /build
COPY  . /build

RUN go mod download

RUN go build -a -o gowt .

FROM scratch as prod

WORKDIR /root/
COPY --from=builder /build/gowt .
COPY --from=builder /build/templates ./templates
COPY --from=builder /build/.env .

ENTRYPOINT ["./gowt"]