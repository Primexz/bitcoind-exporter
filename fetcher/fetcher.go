package fetcher

import (
	"context"
	"time"

	"github.com/Primexz/bitcoind-exporter/config"
	prometheus "github.com/Primexz/bitcoind-exporter/prometheus/metrics"
	goprom "github.com/prometheus/client_golang/prometheus"

	"github.com/Primexz/bitcoind-exporter/util"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(logrus.Fields{
	"prefix": "fetcher",
})

func Start() {
	for {
		NewRunner().run()

		time.Sleep(time.Duration(config.C.FetchInterval) * time.Second)
	}
}

type Runner struct {
	client *Client
}

func NewRunner() *Runner {
	return &Runner{
		client: NewClient(),
	}
}

func (r *Runner) run() {
	start := time.Now()

	blockChainInfo := r.getBlockchainInfo()
	memPoolInfo := r.getMempoolInfo()
	memoryInfo := r.getMemoryInfo()
	indexInfo := r.getIndexInfo()
	networkInfo := r.getNetworkInfo()

	feeRate2 := r.getSmartFee(2)
	feeRate5 := r.getSmartFee(5)
	feeRate20 := r.getSmartFee(20)

	hasRateLatest := r.getNetworkHashrate(-1)
	hashRate1 := r.getNetworkHashrate(1)
	hasthRate120 := r.getNetworkHashrate(120)

	netTotals := r.getNetTotals()

	if util.AnyNil(blockChainInfo, memPoolInfo, memoryInfo, indexInfo, networkInfo, feeRate2, feeRate5, feeRate20, hasRateLatest, hashRate1, hasthRate120, netTotals) {
		log.Error("Failed to fetch data")
		return
	}

	//Blockchain
	prometheus.BlockchainBlocks.Set(float64(blockChainInfo.Blocks))
	prometheus.BlockchainHeaders.Set(float64(blockChainInfo.Headers))
	prometheus.BlockchainVerificationProgress.Set(blockChainInfo.VerificationProgress)
	prometheus.BlockchainSizeOnDisk.Set(float64(blockChainInfo.SizeOnDisk))

	//Mempool
	prometheus.MempoolUsage.Set(float64(memPoolInfo.Usage))
	prometheus.MempoolMax.Set(float64(memPoolInfo.MaxMempool))
	prometheus.MempoolTransactionCount.Set(float64(memPoolInfo.Size))

	//Memory
	prometheus.MemoryUsed.Set(float64(memoryInfo.Locked.Used))
	prometheus.MemoryFree.Set(float64(memoryInfo.Locked.Free))
	prometheus.MemoryTotal.Set(float64(memoryInfo.Locked.Total))
	prometheus.MemoryLocked.Set(float64(memoryInfo.Locked.Locked))
	prometheus.ChunksUsed.Set(float64(memoryInfo.Locked.ChunksUsed))
	prometheus.ChunksFree.Set(float64(memoryInfo.Locked.ChunksFree))

	//TxIndex
	prometheus.TxIndexSynced.Set(float64(util.BoolToFloat64(indexInfo.TxIndex.Synced)))
	prometheus.TxIndexBestHeight.Set(float64(indexInfo.TxIndex.BestBlockHeight))

	//Network
	prometheus.TotalConnections.Set(float64(networkInfo.TotalConnections))
	prometheus.ConnectionsIn.Set(float64(networkInfo.ConnectionsIn))
	prometheus.ConnectionsOut.Set(float64(networkInfo.TotalConnections - networkInfo.ConnectionsIn))
	prometheus.TotalBytesRecv.Set(float64(netTotals.TotalBytesRecv))
	prometheus.TotalBytesSent.Set(float64(netTotals.TotalBytesSent))

	//SmartFee
	prometheus.SmartFee.With(goprom.Labels{"blocks": "2"}).Set(util.ConvertBTCkBToSatVb(feeRate2.Feerate))
	prometheus.SmartFee.With(goprom.Labels{"blocks": "5"}).Set(util.ConvertBTCkBToSatVb(feeRate5.Feerate))
	prometheus.SmartFee.With(goprom.Labels{"blocks": "20"}).Set(util.ConvertBTCkBToSatVb(feeRate20.Feerate))

	//Mining
	prometheus.MiningHashrate.With(goprom.Labels{"blocks": "-1"}).Set(hasRateLatest)
	prometheus.MiningHashrate.With(goprom.Labels{"blocks": "1"}).Set(hashRate1)
	prometheus.MiningHashrate.With(goprom.Labels{"blocks": "120"}).Set(hasthRate120)

	//Internal
	prometheus.ScrapeTime.Set(float64(time.Since(start).Milliseconds()))
}

func (r *Runner) getBlockchainInfo() *BlockchainInfo {
	var info *BlockchainInfo
	err := r.client.RpcClient.CallFor(context.TODO(), &info, "getblockchaininfo")
	if err != nil {
		log.WithError(err).Error("Failed to call RPC")
		return nil
	}

	return info
}

func (r *Runner) getMempoolInfo() *MempoolInfo {
	var info *MempoolInfo
	err := r.client.RpcClient.CallFor(context.TODO(), &info, "getmempoolinfo")
	if err != nil {
		log.WithError(err).Error("Failed to call RPC")
		return nil
	}

	return info
}

func (r *Runner) getMemoryInfo() *MemoryInfo {
	var info *MemoryInfo
	err := r.client.RpcClient.CallFor(context.TODO(), &info, "getmemoryinfo")
	if err != nil {
		log.WithError(err).Error("Failed to call RPC")
		return nil
	}

	return info
}

func (r *Runner) getIndexInfo() *IndexInfo {
	var info *IndexInfo
	err := r.client.RpcClient.CallFor(context.TODO(), &info, "getindexinfo")
	if err != nil {
		log.WithError(err).Error("Failed to call RPC")
		return nil
	}

	return info
}

func (r *Runner) getNetworkInfo() *NetworkInfo {
	var info *NetworkInfo
	err := r.client.RpcClient.CallFor(context.TODO(), &info, "getnetworkinfo")
	if err != nil {
		log.WithError(err).Error("Failed to call RPC")
		return nil
	}

	return info
}

func (r *Runner) getSmartFee(blocks int) *SmartFee {
	var info *SmartFee
	err := r.client.RpcClient.CallFor(context.TODO(), &info, "estimatesmartfee", blocks)
	if err != nil {
		log.WithError(err).Error("Failed to call RPC")
		return nil
	}

	return info
}

func (r *Runner) getNetworkHashrate(blocks int) float64 {
	var info float64
	err := r.client.RpcClient.CallFor(context.TODO(), &info, "getnetworkhashps", blocks)
	if err != nil {
		log.WithError(err).Error("Failed to call RPC")
		return 0
	}

	return info
}

func (r *Runner) getNetTotals() *NetTotals {
	var info *NetTotals
	err := r.client.RpcClient.CallFor(context.TODO(), &info, "getnettotals")
	if err != nil {
		log.WithError(err).Error("Failed to call RPC")
		return nil
	}

	return info
}
