package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/getlevvel/blockchain-digital-notary/Godeps/_workspace/src/github.com/eris-ltd/common/go/ipfs"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//take filename & send ask chain for hash
		str := r.URL.Path[1:]

		fn := strings.Split(str, "/")
		hash := getInfos(fn[1])

		err := ipfs.GetFromIPFS(hash, fn[1], "", bytes.NewBuffer([]byte{}))
		if err != nil {
			fmt.Printf("error getting file from IPFS: %v\n", err)
		}
		contents, err := ioutil.ReadFile(fn[1])
		if err != nil {
			fmt.Printf("error reading file: %v\n", err)
		}
		w.Write(contents)

		//or don't remove? & just output like above
		err = os.Remove(fn[1])
		if err != nil {
			fmt.Printf("error removing file: %v\n", err)
		}
	}
}
