FROM golang:1.15.8-alpine3.13 AS build

RUN apk --no-cache add gcc g++ make ca-certificates

WORKDIR /go/src/github.com/bramalho/go-cqrs

COPY go.mod .
COPY go.sum .
COPY db db
COPY event event
COPY model model
COPY search search
COPY service-pusher service-pusher
COPY service-query service-query
COPY service-todo service-todo
COPY util util

RUN go mod vendor

RUN go install -mod vendor ./...

FROM alpine:3.13.2

WORKDIR /usr/bin

COPY --from=build /go/bin .
