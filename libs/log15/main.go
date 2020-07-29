package main

import (
	"fmt"

	log "github.com/inconshreveable/log15"
)

func main() {
	fmt.Println("log15")

	// lvl, _ := log.LvlFromString("warn")

	// glogger.Verbosity(lvl)

	log.Root()

	svrlog := log.New("module", "app/server")

	svrlog.Info("abcS")

	svrlog.Info("abc", "cc", "abc")

	svrlog.Error("ddddddd")

	svrlog.Debug("debug")
}
