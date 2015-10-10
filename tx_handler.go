package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/getlevvel/blockchain-digital-notary/Godeps/_workspace/src/github.com/eris-ltd/common/go/ipfs"
	"github.com/getlevvel/blockchain-digital-notary/Godeps/_workspace/src/github.com/eris-ltd/mint-client/mintx/core"
	"github.com/getlevvel/blockchain-digital-notary/Godeps/_workspace/src/github.com/tendermint/tendermint/types"
	"github.com/getlevvel/blockchain-digital-notary/Godeps/_workspace/src/github.com/tendermint/tendermint/wire"
)

func receiveNameTx(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		str := r.URL.Path[1:]
		hash := strings.Split(str, "/")[1]

		fmt.Printf("hash in rnt %v\n", hash)

		txData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("error reading body: %v\n", err)
		}

		tx := new(types.NameTx)
		n := new(int64)
		txD := bytes.NewReader(txData)

		wire.ReadBinary(tx, txD, n, &err)
		if err != nil {
			fmt.Printf("error reading binary: %v\n", err)
		}

		rpcAddr := "http://0.0.0.0:46657"
		_, err1 := core.Broadcast(tx, rpcAddr)
		if err1 != nil {
			fmt.Printf("error broadcasting: %v\n", err1)
		}

		//TODO check Name reg

		endpoint := "http://0.0.0.0:11113/" + "cacheHash/" + hash
		_, err2 := http.Post(endpoint, "", nil)
		if err2 != nil {
			fmt.Printf("cache post error: %v\n", err2)
		}
		w.Write([]byte("success!"))
	}
}

func cacheHash(w http.ResponseWriter, r *http.Request) {
	//XXX this endpoint should also require authentication
	//if name reg update is legit, make post to ipfs hosts w/ hash in URL;
	//endpoint takes hash and pins it locally
	str := r.URL.Path[1:]
	hash := strings.Split(str, "/")[1]

	pinned, err := ipfs.PinToIPFS(hash, bytes.NewBuffer([]byte{}))
	if err != nil {
		fmt.Printf("error pinning to IPFS: %v\n", err)
	}
	fmt.Printf("%v has been added succesfully\n", pinned)
}
