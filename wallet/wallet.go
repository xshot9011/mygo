package wallet

import (
	"encoding/csv"
	"os"
	"strconv"

	log "github.com/xshot9011/mygo/logger"
)

var (
	Delimiter = ','
	Comment   = '#'
)

type Wallet struct {
	info []WalletCSV
}

type WalletCSV struct {
	income    float64
	regular   float64
	learning  float64
	happiness float64
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

func getHeaderInfo(reader *csv.Reader) map[string]int {
	log.Trace(log.GetCurrentLog())
	log.Trace("Take a look at the first line of the file.")

	headerInfo := make(map[string]int) // "column_name" -> index
	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	for index, field := range header {
		headerInfo[field] = index
	}

	log.Debugf("headerInfo: %v", headerInfo)
	log.Trace(log.GetEndCurrentLog())
	return headerInfo
}

func getRecord(reader *csv.Reader) [][]string {
	log.Trace(log.GetCurrentLog())
	log.Trace("Take a look at following records")

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("records: %v", records)
	log.Trace(log.GetEndCurrentLog())
	return records
}

func readCSV(reader *csv.Reader) (map[string]int, [][]string) {
	return getHeaderInfo(reader), getRecord(reader)
}

func (wallet *Wallet) initValue(headerInfo map[string]int, records [][]string) {
	for _, record := range records {
		income, err := strconv.ParseFloat(record[headerInfo["income"]], 64)
		if err != nil {
			log.Fatal(err)
		}
		regular, err := strconv.ParseFloat(record[headerInfo["regular"]], 64)
		if err != nil {
			log.Fatal(err)
		}
		learning, err := strconv.ParseFloat(record[headerInfo["learning"]], 64)
		if err != nil {
			log.Fatal(err)
		}
		happiness, err := strconv.ParseFloat(record[headerInfo["happiness"]], 64)
		if err != nil {
			log.Fatal(err)
		}
		wallet.info = append(wallet.info, WalletCSV{
			income:    income,
			regular:   regular,
			learning:  learning,
			happiness: happiness,
		})
	}
	log.Debugf("Initial wallet value: %v", wallet)
}

func New(fileName string, selectFields []string) *Wallet {
	log.Trace(log.GetCurrentLog())

	wallet := Wallet{}
	file := getFile(fileName)
	reader := getReader(file)
	headerInfo, records := readCSV(reader)

	wallet.initValue(headerInfo, records)
	log.Debugf("New wallet(%v): %v", len(wallet.info), wallet)

	log.Trace(log.GetEndCurrentLog())
	defer file.Close()
	return &Wallet{}
}
