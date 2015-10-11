package main

import (
	. "github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/eris-ltd/common/go/log"
)

var logger *Logger

func init() {
	logger = AddLogger("main")
}