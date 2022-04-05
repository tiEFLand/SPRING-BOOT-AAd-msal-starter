
package okex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	InstrumentId = "BTC-USD-181228"
	currency     = "BTC"
)

/*
 OKEX general api's testing
*/
func TestGetServerTime(t *testing.T) {
	serverTime, err := NewTestClient().GetServerTime()
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("OKEX's server time: ", serverTime)
}

func TestGetFuturesExchangeRate(t *testing.T) {
	exchangeRate, err := NewTestClient().GetFuturesExchangeRate()
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Current exchange rate: ", exchangeRate)
}

/*
 Futures market api's testing
*/
func TestGetFuturesInstruments(t *testing.T) {
	Instruments, err := NewTestClient().GetFuturesInstruments()
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instruments: ", Instruments)
}

func TestGetFuturesInstrumentsCurrencies(t *testing.T) {
	currencies, err := NewTestClient().GetFuturesInstrumentCurrencies()
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instrument currencies: ", currencies)
}

func TestGetFuturesInstrumentBook(t *testing.T) {
	insId := getValidInstrumentId()
	book, err := NewTestClient().GetFuturesInstrumentBook(insId, nil)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instrument book: ", book)
}

func TestGetFuturesInstrumentBook2(t *testing.T) {
	params := NewParams()
	params["size"] = "10"
	params["depth"] = "0.1"
	insId := getValidInstrumentId()
	r, err := NewTestClient().GetFuturesInstrumentBook(insId, nil)

	simpleAssertTrue(r, err, t, false)
}

func TestGetFuturesInstrumentAllTicker(t *testing.T) {
	tickers, err := NewTestClient().GetFuturesInstrumentAllTicker()
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instrument all ticker: ", tickers)
}

func TestGetFuturesInstrumentTicker(t *testing.T) {
	ticker, err := NewTestClient().GetFuturesInstrumentTicker(InstrumentId)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instrument ticker: ", ticker)
}

func TestGetFuturesInstrumentTrades(t *testing.T) {
	trades, err := NewTestClient().GetFuturesInstrumentTrades(InstrumentId)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instrument trades: ", trades)
}

func TestGetFuturesInstrumentCandles(t *testing.T) {
	//start := "2018-06-20T02:31:00Z"
	//end := "2018-06-20T02:55:00Z"
	granularity := CANDLES_1MIN

	optional := map[string]string{}
	//optional["start"] = start
	//optional["end"] = end
	optional["granularity"] = Int2String(granularity)

	insId := getValidInstrumentId()

	candles, err := NewTestClient().GetFuturesInstrumentCandles(insId, optional)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Futures Instrument candles:")
	for i, outLen := 0, len(candles); i < outLen; i++ {
		candle := candles[i]
		for j, inLen := 0, 7; j < inLen; j++ {
			if j == 0 {
				fmt.Print("timestamp:")
				fmt.Print(candle[j])
			} else if j == 1 {
				fmt.Print(" open:")
				fmt.Print(candle[j])
			} else if j == 2 {
				fmt.Print(" high:")
				fmt.Print(candle[j])
			} else if j == 3 {
				fmt.Print(" low:")
				fmt.Print(candle[j])