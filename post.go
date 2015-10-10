package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/getlevvel/blockchain-digital-notary/Godeps/_workspace/src/github.com/eris-ltd/common/go/ipfs"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//TODO get amt from url
		str := r.URL.Path[1:]
		fn := strings.Split(str, "/")[0]

		body := r.Body
		b, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Printf("error reading body: %v\n", err)
		}

		// which file perms?
		err = ioutil.WriteFile(fn, b, 0666)
		if err != nil {
			fmt.Printf("error writing temp file: %v\n", err)
		}
		//should just put on whoever is doing the sending's gateway; since cacheHash won't send it there anyways
		hash, err := ipfs.SendToIPFS(fn, "eris", bytes.NewBuffer([]byte{}))
		if err != nil {
			fmt.Printf("error sending to IPFS: %v\n", err)
		}
		//TODO rm file

		_, _ = UpdateNameReg(fn, hash, "")

	}
}
