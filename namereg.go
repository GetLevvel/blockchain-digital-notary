package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/getlevvel/blockchain-digital-notary/Godeps/_workspace/src/github.com/eris-ltd/mint-client/mintx/core"
	"github.com/getlevvel/blockchain-digital-notary/Godeps/_workspace/src/github.com/tendermint/tendermint/wire"
)

//TODO pass amt in from url. (make estimate of length)
func UpdateNameReg(fileName, hash, amt string) (string, error) {
	nodeAddr := os.Getenv("MINTX_NODE_ADDR")
	signAddr := os.Getenv("MINTX_SIGN_ADDR")
	chainID := os.Getenv("MINTX_CHAINID")
	pubkey := os.Getenv("MINTX_PUBKEY")
	addr := ""
	amtS := "1000000"
	nonceS := ""
	feeS := "0"
	name := fileName
	data := hash

	//build and sign a nameTx, send it away for broadcasting
	nTx, err := core.Name(nodeAddr, pubkey, addr, amtS, nonceS, feeS, name, data)
	if err != nil {
		fmt.Printf("corename error: %v\n", err)
	}

	//sign but don't broadcast
	_, err = core.SignAndBroadcast(chainID, "", signAddr, nTx, true, false, false)
	if err != nil {
		fmt.Printf("sign error: %v\n", err)
	}

	n := new(int64)
	w := new(bytes.Buffer)
	wire.WriteBinary(nTx, w, n, &err)

	err = postTxData(nodeAddr, hash, w.Bytes())
	if err != nil {
		fmt.Printf("post error: %v\n", err)
	}
	return "", nil
}

func postTxData(nodeAddr, hash string, txData []byte) error {

	// post needs to be to toadserver endpoint, which'll route the TX to the node
	// or should the toadserver have an endpoint that does that?
	txD := bytes.NewReader(txData)
	//it can also query for th name reg to ensure things are good
	endpoint := "http://0.0.0.0:11113/" + "receiveNameTx/" + hash
	_, err := http.Post(endpoint, "", txD)
	if err != nil {
		fmt.Printf("post error: %v\n", err)
	}
	return nil
}
