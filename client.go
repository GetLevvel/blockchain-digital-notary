package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	cclient "github.com/getlevvel/blockchain-digital-notary/Godeps/_workspace/src/github.com/tendermint/tendermint/rpc/core_client"
)

var (
	DefaultNodeRPCHost = "0.0.0.0"
	DefaultNodeRPCPort = "46657"
	DefaultNodeRPCAddr = "http://" + DefaultNodeRPCHost + ":" + DefaultNodeRPCPort

	DefaultChainID string

	REQUEST_TYPE = "JSONRPC"
	client       cclient.Client
)

// override the hardcoded defaults with env variables if they're set
func init() {
	nodeAddr := os.Getenv("MINTX_NODE_ADDR")
	if nodeAddr != "" {
		DefaultNodeRPCAddr = nodeAddr
	}

	chainID := os.Getenv("MINTX_CHAINID")
	if chainID != "" {
		DefaultChainID = chainID
	}
}

func getInfos(fileName string) string {
	if fileName == "" {
		//to eventually support an endpoint that lists available files
		_, err := client.ListNames()
		ifExit(err)
		//formatOutput(r)
		return "" //result of format output
	} else {
		_, err := client.GetName(fileName)
		ifExit(err)
		//formatOutput(r)
		return "" //result of format output
	}
}

//this func is just a check
func checkAddr(addr string, w io.Writer) error {
	if addr == "" {
		_, err := client.ListAccounts()
		ifExit(err)
		//formatOutput(r)
		return nil //result of format output
	} else {
		addrBytes, err := hex.DecodeString(addr)
		if err != nil {
			exit(fmt.Errorf("Addr %s is improper hex: %v", addr, err))
		}
		r, err := client.GetAccount(addrBytes)
		ifExit(err)
		if r == nil {
			exit(fmt.Errorf("Account %X does not exist", addrBytes))
		}
		r2 := r.Account
		if r2 == nil {
			exit(fmt.Errorf("Account %X does not exist", addrBytes))
		}
		//formatOutput(c, 1, r2)
	}

	//TODO deal with this gracefully
	//	w.Write([]byte("Permission denied, invalid address\n"))
	return nil //errors.New("Permission denied, invalid address")

	//get more infos (like check if they have perms!)

}
