package fetcher

type BlockchainInfo struct {
	Chain                string  `json:"chain"`
	Blocks               int     `json:"blocks"`
	Headers              int     `json:"headers"`
	BestBlockhash        string  `json:"bestblockhash"`
	Difficulty           float64 `json:"difficulty"`
	Time                 int     `json:"time"`
	MedianTime           int     `json:"mediantime"`
	VerificationProgress float64 `json:"verificationprogress"`
	InitialBlockDownload bool    `json:"initialblockdownload"`
	ChainWork            string  `json:"chainwork"`
	SizeOnDisk           int     `json:"size_on_disk"`
	Pruned               bool    `json:"pruned"`
	Warnings             string  `json:"warnings"`
}

type MempoolInfo struct {
	Loaded              bool    `json:"loaded"`
	Size                int     `json:"size"`
	Bytes               int     `json:"bytes"`
	Usage               int     `json:"usage"`
	TotalFee            float64 `json:"total_fee"`
	MaxMempool          int     `json:"maxmempool"`
	MempoolMinFee       float64 `json:"mempoolminfee"`
	MinRelayTxFee       float64 `json:"minrelaytxfee"`
	IncrementalRelayFee float64 `json:"incrementalrelayfee"`
	UnbroadcastCount    int     `json:"unbroadcastcount"`
	FullRBF             bool    `json:"fullrbf"`
}

type MemoryInfo struct {
	Locked struct {
		Used       int `json:"used"`
		Free       int `json:"free"`
		Total      int `json:"total"`
		Locked     int `json:"locked"`
		ChunksUsed int `json:"chunks_used"`
		ChunksFree int `json:"chunks_free"`
	} `json:"locked"`
}

type IndexInfo struct {
	TxIndex struct {
		Synced          bool `json:"synced"`
		BestBlockHeight int  `json:"best_block_height"`
	}
}

type NetworkInfo struct {
	Version            int      `json:"version"`
	Subversion         string   `json:"subversion"`
	ProtocolVersion    int      `json:"protocolversion"`
	LocalServices      string   `json:"localservices"`
	LocalServicesNames []string `json:"localservicesnames"`
	LocalRelay         bool     `json:"localrelay"`
	Timeoffset         int      `json:"timeoffset"`
	TotalConnections   int      `json:"connections"`
	ConnectionsIn      int      `json:"connections_in"`
	ConnectionsOut     int      `json:"connections_out"`
	RelayFee           float64  `json:"relayfee"`
	IncrementalFee     float64  `json:"incrementalfee"`
}
