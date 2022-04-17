package main

import (
	log "github.com/xshot9011/mygo/logger"
	"github.com/xshot9011/mygo/wallet"
)

var (
	fileName     = "./user.csv"
	selectFields = []string{"income", "regular", "learning", "happiness"}
)

func init() {
}

func main() {
	log.Trace(log.GetCurrentLog())

	wallet := wallet.New(fileName, selectFields)
	log.Debugf("Main get wallet: %v", wallet)

	log.Trace(log.GetEndCurrentLog())
}
