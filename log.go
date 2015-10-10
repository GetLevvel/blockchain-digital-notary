import (
	. "github.com/eris-ltd/eris-keys/Godeps/_workspace/src/github.com/eris-ltd/common/log"
)

var logger *Logger

func init() {
	logger = AddLogger("main")
}