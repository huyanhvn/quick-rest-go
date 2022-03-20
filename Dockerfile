# syntax=docker/dockerfile:1

FROM golang:1.17-alpine as builder
RUN apk add --no-cache build-base  # cgo
WORKDIR /go/src
COPY . ./
RUN GO111MODULE=on go build -v -o /quick-rest .

FROM alpine:latest
ENV GIN_MODE=release
EXPOSE 8080
COPY --from=builder /quick-rest /
CMD ["/quick-rest"]
