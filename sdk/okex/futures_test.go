
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
			} else if j == 4 {
				fmt.Print(" close:")
				fmt.Print(candle[j])
			} else if j == 5 {
				fmt.Print(" volume:")
				fmt.Print(candle[j])
			} else if j == 6 {
				fmt.Print(" currency_volume:")
				fmt.Println(candle[j])
			}
		}
	}
}

func TestGetFuturesInstrumentIndex(t *testing.T) {
	index, err := NewTestClient().GetFuturesInstrumentIndex(InstrumentId)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instrument index: ", index)
}

func TestGetFuturesInstrumentEstimatedPrice(t *testing.T) {
	estimatedPrice, err := NewTestClient().GetFuturesInstrumentEstimatedPrice(InstrumentId)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instrument estimated price: ", estimatedPrice)
}

func TestGetFuturesInstrumentOpenInterest(t *testing.T) {
	priceLimit, err := NewTestClient().GetFuturesInstrumentOpenInterest(InstrumentId)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instrument open interest: ", priceLimit)
}

func TestGetFuturesInstrumentPriceLimit(t *testing.T) {
	priceLimit, err := NewTestClient().GetFuturesInstrumentPriceLimit(InstrumentId)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instrument price limit: ", priceLimit)
}

func TestGetFuturesInstrumentLiquidation(t *testing.T) {
	InstrumentIdx := "EOS-USD-181228"
	status, from, to, limit := 1, 1, 0, 5
	liquidation, err := NewTestClient().GetFuturesInstrumentLiquidation(InstrumentIdx, status, from, to, limit)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures Instrument liquidation: ", liquidation)
}

/*
 Futures trade api's testing
*/
func TestGetFuturesPositions(t *testing.T) {
	position, err := NewTestClient().GetFuturesPositions()
	if err != nil {
		t.Error(err)
	}
	if position.MarginMode == "crossed" {
		FmtPrintln("Futures crossed position: ", position)
	} else if position.MarginMode == "fixed" {
		FmtPrintln("Futures fixed position: ", position)
	} else {
		FmtPrintln("Futures position failed: ", position)
	}
}

func TestGetFuturesInstrumentPosition(t *testing.T) {
	position, err := NewTestClient().GetFuturesInstrumentPosition(InstrumentId)
	if err != nil {
		t.Error(err)
	}
	if position.MarginMode == "crossed" {
		FmtPrintln("Futures crossed position: ", position)
	}
	if position.MarginMode == "fixed" {
		FmtPrintln("Futures fixed position: ", position)
	} else {
		FmtPrintln("Futures position failed: ", position)
	}
}

func TestGetFuturesAccounts(t *testing.T) {
	account, err := NewTestClient().GetFuturesAccounts()
	if err != nil {
		t.Error(err)
	}
	if account.MarginMode == "crossed" {
		FmtPrintln("Futures crossed account: ", account)
	} else if account.MarginMode == "fixed" {
		FmtPrintln("Futures fixed account: ", account)
	} else {
		FmtPrintln("Futures account failed: ", account)
	}
}

func TestGetFuturesAccountsByCurrency(t *testing.T) {
	currencyAccounts, err := NewTestClient().GetFuturesAccountsByCurrency(currency)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures currency accounts: ", currencyAccounts)
}

func TestGetFuturesAccountsLedgerByCurrency(t *testing.T) {
	from, to, limit := 1, 0, 2
	ledger, err := NewTestClient().GetFuturesAccountsLedgerByCurrency(currency, from, to, limit)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures currency ledger: ", ledger)
}

func TestGetFuturesAccountsHoldsByInstrumentId(t *testing.T) {
	holds, err := NewTestClient().GetFuturesAccountsHoldsByInstrumentId(InstrumentId)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures currency holds: ", holds)
}

func TestFuturesOrder(t *testing.T) {
	var newOrderParams FuturesNewOrderParams
	newOrderParams.ClientOid = "od12345678"
	newOrderParams.InstrumentId = InstrumentId
	newOrderParams.Type = IntToString(OPEN_SHORT)
	newOrderParams.Price = "100000.00"
	newOrderParams.Size = "1"
	newOrderParams.MatchPrice = "0"
	newOrderParams.Leverage = "20"

	result, err := NewTestClient().FuturesOrder(newOrderParams)
	if err != nil {
		t.Error(err)
	}
	FmtPrintln("Futures new order: ", result)
}

func TestFuturesOrders(t *testing.T) {
	var batchNewOrder FuturesBatchNewOrderParams
	batchNewOrder.InstrumentId = InstrumentId
	batchNewOrder.Leverage = "20"
	var ordersData [5]FuturesBatchNewOrderItem
	for i, loop := 1, 6; i < loop; i++ {
		var item FuturesBatchNewOrderItem
		item.ClientOid = "od" + IntToString(12345670+i)
		item.Type = IntToString(OPEN_SHORT)
		item.Price = IntToString(100000 + i)
		item.Size = "1"
		item.MatchPrice = "0"
		ordersData[i-1] = item
	}
	json, err := Struct2JsonString(ordersData)
	if err != nil {
		t.Error(err)
	}
	batchNewOrder.OrdersData = json
	result, err := NewTestClient().FuturesOrders(batchNewOrder)
	if err != nil {
		t.Error(err)
	}