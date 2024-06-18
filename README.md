# Bitcoind Prometheus Exporter ₿

**Prometheus metrics for a bitcoin node made simple**

![Buid](https://img.shields.io/github/actions/workflow/status/primexz/bitcoind-exporter/release.yml)
![License](https://img.shields.io/github/license/primexz/krakendca)

## 🔍 About the project

A Prometheus Exporter, which provides a deep insight into a Bitcoin full node.


## ⚙️ Configuration

This tool is configured via environment variables. Some environment variables are required and some activate additional functionalities.


| Variable | Description | Required | Default |
| --- | --- | --- | --- |
| `RPC_ADDRESS` | The RPC address for the Bitcoin full node, e.g. ``http://127.0.0.1:8332`` | ✅ | |
| `RPC_USER` |The user name that was defined in the Bitcoin Node configuration | ✅ | |
| `RPC_PASS` | The password that was set in the Bitcoin Node configuration | ✅  |  |
| `ZMQ_ADDRESS` | The address to the ZeroMQ interface of the Bitcoin Fullnode. This variable is required to determine the transcation rates. e.g. ``127.0.0.1:28333`` | ❌ |  |
| `FETCH_INTERVAL` | The interval at which the metrics are to be recalculated. | ❌ | `10` |
| `METRIC_PORT` | The port via which the metrics are provided. | ❌ | `3000` |
| `LOG_LEVEL` | The log level for the service | ❌ | `info` |

## 💻 Grafana Dashboard

The official Grafana dashboard can be found here: https://grafana.com/grafana/dashboards/21351

## 🐳 Run with Docker

###  Docker-CLI

```bash
docker run -d --name kraken_dca \
  -e RPC_ADDRESS=http://127.0.0.1:8332 \
  -e RPC_USER=mempool \
  -e RPC_PASS=mempool \
  -e ZMQ_ADDRESS=127.0.0.1:28333 \
   ghcr.io/primexz/bitcoind-exporter:latest
```

### 🚀 Docker-Compose

```bash
vim docker-compose.yml
```

```yaml
version: "3.8"
services:
  kraken_dca:
    image: ghcr.io/primexz/bitcoind-exporter:latest
    environment:
      - RPC_ADDRESS=http://127.0.0.1:8332
      - RPC_USER=mempool
      - RPC_PASS=mempool
      - ZMQ_ADDRESS=27.0.0.1:28333
    restart: always
```

```bash
docker-compose up -d
```