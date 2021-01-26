FROM golang:1.13-alpine3.11 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/heru-wijaya/go-grpc-skeleton
COPY go.mod go.sum ./
COPY . .
RUN go mod vendor
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./cmd

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]