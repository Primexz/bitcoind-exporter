package fetcher

import (
	"fmt"
	"os"
	"strings"

	"github.com/Primexz/bitcoind-exporter/config"
	"github.com/Primexz/bitcoind-exporter/util"
	"github.com/ybbus/jsonrpc/v3"
)

// "http://my-rpc-service:8080/rpc"
var rpcClient = jsonrpc.NewClientWithOpts(computeAddress(), &jsonrpc.RPCClientOpts{
	CustomHeaders: map[string]string{
		"Authorization": "Basic " + computeBasicAuth(),
	},
})

func computeBasicAuth() string {
	user := config.C.RPCUser
	pass := config.C.RPCPass
	cookieFile := config.C.RPCCookieFile

	if cookieFile != "" {
		cookie, err := os.ReadFile(cookieFile)
		if err != nil {
			log.WithError(err).Fatal("Failed to read cookie file")
			return ""
		}
		cookieStr := strings.TrimSpace(string(cookie))

		if !strings.Contains(cookieStr, ":") {
			log.WithError(err).Fatal("Invalid cookie file format")
			return ""
		}

		return util.StringToBase64(cookieStr)
	}

	return util.StringToBase64(fmt.Sprintf("%s:%s", user, pass))
}

func computeAddress() string {
	address := config.C.RPCAddress

	if strings.HasPrefix(address, "http://") || strings.HasPrefix(address, "https://") {
		return address
	} else {
		return fmt.Sprintf("http://%s", address)
	}
}
