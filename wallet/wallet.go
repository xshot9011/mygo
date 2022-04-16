package wallet

import (
	"encoding/csv"
	"os"

	log "github.com/xshot9011/mygo/logger"
)

var (
	Delimiter = ','
	Comment   = '#'
)

type Wallet interface {
	getData() ([]string, [][]string)
}

type WalletCSV struct {
	header  []string
	records [][]string
}

func getReader(file *os.File) *csv.Reader {
	log.Trace(log.GetCurrentLog())

	reader := csv.NewReader(file)
	reader.Comma = Delimiter
	reader.Comment = Comment

	log.Trace(log.GetEndCurrentLog())
	return reader
}

func getFile(fileName string) *os.File {
	log.Trace(log.GetCurrentLog())

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Trace(log.GetEndCurrentLog())
	return file
}

func getData(reader *csv.Reader) ([]string, [][]string) {
	log.Trace(log.GetCurrentLog())

	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("Header: %v", header)
	log.Debugf("Records: %s", records)
	log.Trace(log.GetEndCurrentLog())
	return header, records
}

func New(fileName string) *WalletCSV {
	log.Trace(log.GetCurrentLog())

	wallet := WalletCSV{}

	file := getFile(fileName)
	reader := getReader(file)
	header, records := getData(reader)

	wallet.header = header
	wallet.records = records

	log.Trace(log.GetEndCurrentLog())
	return &wallet
}
