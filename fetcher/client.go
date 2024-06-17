package fetcher

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/Primexz/bitcoind-exporter/config"
	"github.com/ybbus/jsonrpc/v3"
)

// "http://my-rpc-service:8080/rpc"
var rpcClient = jsonrpc.NewClientWithOpts(computeAddress(), &jsonrpc.RPCClientOpts{
	CustomHeaders: map[string]string{
		"Authorization": "Basic " + computeBasicAuth(),
	},
})

func computeBasicAuth() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.C.RPCUser, config.C.RPCPass)))
}

func computeAddress() string {
	address := config.C.RPCAddress

	if strings.HasPrefix(address, "http://") {
		return address
	} else {
		return fmt.Sprintf("http://%s", address)
	}
}
