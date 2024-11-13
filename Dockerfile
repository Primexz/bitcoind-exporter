FROM golang:1.22.4-alpine3.20 AS build

COPY . /go/src
WORKDIR /go/src

RUN go mod download
RUN GOOS=linux go build -o /go/bin/bitcoind-exporter -v .

FROM alpine
COPY --from=build /go/bin/bitcoind-exporter /usr/bin/bitcoind-exporter

ENTRYPOINT ["/usr/bin/bitcoind-exporter"]
