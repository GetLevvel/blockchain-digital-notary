package config

import (
	// We can't use github.com/tendermint/tendermint/logger
	// because that would create a dependency cycle.
	"github.com/getlevvel/blockchain-digital-notary/Godeps/_workspace/src/github.com/tendermint/log15"
)

var log = log15.New("module", "config")
