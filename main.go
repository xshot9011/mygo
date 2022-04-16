package main

import (
	log "github.com/xshot9011/mygo/logger"
	"github.com/xshot9011/mygo/wallet"
)

var (
	fileName string = "./user.csv"
)

func init() {
}

func main() {
	log.Trace(log.GetCurrentLog())

	wallet := wallet.New(fileName)
	log.Debugf("Get wallet: %v", wallet)

	log.Trace(log.GetEndCurrentLog())
}
