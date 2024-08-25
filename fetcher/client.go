package fetcher

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Primexz/bitcoind-exporter/config"
	"github.com/ybbus/jsonrpc/v3"
)

// Create RPC client with cookie-based authentication
var rpcClient = jsonrpc.NewClientWithOpts(computeAddress(), &jsonrpc.RPCClientOpts{
	CustomHeaders: map[string]string{
		"Authorization": "Basic " + computeCookieAuth(),
	},
})

// Reads the RPC cookie for authentication and encodes it for Basic Auth
func computeCookieAuth() string {
	cookieFilePath := config.C.RPCCookieFile
	if cookieFilePath == "" {
		panic("RPCCookieFile path is not set in the configuration")
	}

	// Read the contents of the cookie file
	cookie, err := ioutil.ReadFile(cookieFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read RPC cookie file from '%s': %v", cookieFilePath, err))
	}

	// The cookie file contains the format: <username>:<password>
	// Trim whitespace and ensure valid format
	cookieAuth := strings.TrimSpace(string(cookie))
	if !strings.Contains(cookieAuth, ":") {
		panic(fmt.Sprintf("Invalid cookie format in file '%s'", cookieFilePath))
	}

	// Base64 encode the cookie value for Basic Auth
	return base64.StdEncoding.EncodeToString([]byte(cookieAuth))
}

// Computes the RPC server address
func computeAddress() string {
	address := config.C.RPCAddress

	if strings.HasPrefix(address, "http://") || strings.HasPrefix(address, "https://") {
		return address
	} else {
		return fmt.Sprintf("http://%s", address)
	}
}
