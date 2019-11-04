package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"time"
	"runtime"
)


func main() {
	runtime.GOMAXPROCS(4)
	start := time.Now()
	stockSymbols := []string{
		"goog",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	numComplete := 0

	for _, symbol  := range stockSymbols {
		go func(symbol string) {
			resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s : %.2f\n", quote.Name, quote.LastPrice)
			numComplete ++
		}(symbol)
	}


	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}


type QuoteResponse  struct {
	Status string
	Name string
	LastPrice float32
	Change string
	ChangePercent float32
	TimeStamp string
	MSDate string
	Volumne int
	ChangeYTD float32
	ChangePercentYTD float32
	High float32
	Low float32
	Open float32
}

