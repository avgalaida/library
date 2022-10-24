FROM golang:1.19-alpine3.16 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/avgalaida/library

COPY go.mod go.sum ./

COPY domain domain
COPY application application
COPY infrastructure infrastructure

RUN GO111MODULE=on go install ./...

FROM alpine:3.16
WORKDIR /usr/bin
COPY --from=build /go/bin .
