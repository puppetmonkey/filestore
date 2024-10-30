package ipfs

import (
	"fmt"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

var ipfs *ipfsapi.Shell

// Init 初始化IPFS连接
func Init(host string) error {
	ipfs = ipfsapi.NewShell(host)
	if ipfs == nil {
		return fmt.Errorf("failed to connect to IPFS")
	}
	return nil
}
