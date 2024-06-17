FROM alpine
COPY bitcoind-exporter /usr/bin/bitcoind-exporter
ENTRYPOINT ["/usr/bin/bitcoind-exporter"]