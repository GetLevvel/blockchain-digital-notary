package main

import (
	"fmt"
	"net/http"
	"os"
)

//TODO use struct
type Infos struct {
	Filename string `json:"filename"`
	Hash     string `json:"hash"`
	//	BTCaddr  string `json:"btc_addr"`
}

func main() {
	fmt.Printf("blockchain notary service starting...\n")

	mux := http.NewServeMux()
	mux.HandleFunc("/postfile/", postHandler) //post a file with its contents to gateway, returns hash
	//ts makes & signs a nameTx, then posts to a node, which does the broadcasting
	mux.HandleFunc("/receiveNameTx/", receiveNameTx) //unpack tx, if valid, add to chain
	mux.HandleFunc("/cacheHash/", cacheHash)         //also if valid, pin hash on all hosts (except one that sent it :~( )

	mux.HandleFunc("/getfile/", getHandler) // request by name, receive contents

	http.ListenAndServe(":11113", mux)
	fmt.Printf("blockchain notary service listening...\n")
}

//-------------------------------------------------------
//--helpers

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func ifExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
