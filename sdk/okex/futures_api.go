
package okex

import (
	"net/http"
	"strings"
)

/*
 OKEX futures contract api
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

/*
 =============================== Futures market api ===============================
*/
/*
 The exchange rate of legal tender pairs
*/
func (client *Client) GetFuturesExchangeRate() (ExchangeRate, error) {
	var exchangeRate ExchangeRate
	_, err := client.Request(GET, FUTURES_RATE, nil, &exchangeRate)
	return exchangeRate, err
}

/*
  Get all of futures contract list
*/
func (client *Client) GetFuturesInstruments() ([]FuturesInstrumentsResult, error) {
	var Instruments []FuturesInstrumentsResult
	_, err := client.Request(GET, FUTURES_INSTRUMENTS, nil, &Instruments)
	return Instruments, err
}

/*
 Get the futures contract currencies
*/
func (client *Client) GetFuturesInstrumentCurrencies() ([]FuturesInstrumentCurrenciesResult, error) {
	var currencies []FuturesInstrumentCurrenciesResult
	_, err := client.Request(GET, FUTURES_CURRENCIES, nil, &currencies)
	return currencies, err
}

/*
	获取深度数据
	获取币对的深度列表。这个请求不支持分页，一个请求返回整个深度列表。

	限速规则：20次/2s
	HTTP请求
	GET /api/spot/v3/instruments/<instrument_id>/book

	签名请求示例
	2018-09-12T07:57:09.130ZGET/api/spot/v3/instruments/LTC-USDT/book?size=10&depth=0.001

*/
func (client *Client) GetFuturesInstrumentBook(InstrumentId string, optionalParams map[string]string) (FuturesInstrumentBookResult, error) {
	var book FuturesInstrumentBookResult
	params := NewParams()
	if optionalParams != nil && len(optionalParams) > 0 {
		params["size"] = optionalParams["size"]
		params["depth"] = optionalParams["depth"]
	}
	requestPath := BuildParams(GetInstrumentIdUri(FUTURES_INSTRUMENT_BOOK, InstrumentId), params)
	_, err := client.Request(GET, requestPath, nil, &book)
	return book, err
}

/*
 Get the futures contract Instrument all ticker
*/
func (client *Client) GetFuturesInstrumentAllTicker() ([]FuturesInstrumentTickerResult, error) {
	var tickers []FuturesInstrumentTickerResult
	_, err := client.Request(GET, FUTURES_TICKERS, nil, &tickers)
	return tickers, err
}

/*
 Get the futures contract Instrument ticker
*/
func (client *Client) GetFuturesInstrumentTicker(InstrumentId string) (FuturesInstrumentTickerResult, error) {
	var ticker FuturesInstrumentTickerResult
	_, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_TICKER, InstrumentId), nil, &ticker)
	return ticker, err
}

/*
 Get the futures contract Instrument trades
*/
func (client *Client) GetFuturesInstrumentTrades(InstrumentId string) ([]FuturesInstrumentTradesResult, error) {
	var trades []FuturesInstrumentTradesResult
	_, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_TRADES, InstrumentId), nil, &trades)
	return trades, err
}

/*
 Get the futures contract Instrument candles
 granularity: @see  file: futures_constants.go
*/
func (client *Client) GetFuturesInstrumentCandles(InstrumentId string, optionalParams map[string]string) ([][]string, error) {
	var candles [][]string
	params := NewParams()

	if optionalParams != nil && len(optionalParams) > 0 {
		params["start"] = optionalParams["start"]
		params["end"] = optionalParams["end"]
		params["granularity"] = optionalParams["granularity"]
	}
	requestPath := BuildParams(GetInstrumentIdUri(FUTURES_INSTRUMENT_CANDLES, InstrumentId), params)
	_, err := client.Request(GET, requestPath, nil, &candles)
	return candles, err
}

/*
 Get the futures contract Instrument index
*/
func (client *Client) GetFuturesInstrumentIndex(InstrumentId string) (FuturesInstrumentIndexResult, error) {
	var index FuturesInstrumentIndexResult
	_, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_INDEX, InstrumentId), nil, &index)
	return index, err
}

/*
 Get the futures contract Instrument estimated price
*/
func (client *Client) GetFuturesInstrumentEstimatedPrice(InstrumentId string) (FuturesInstrumentEstimatedPriceResult, error) {
	var estimatedPrice FuturesInstrumentEstimatedPriceResult
	_, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_ESTIMATED_PRICE, InstrumentId), nil, &estimatedPrice)
	return estimatedPrice, err
}

/*
 Get the futures contract Instrument holds
*/
func (client *Client) GetFuturesInstrumentOpenInterest(InstrumentId string) (FuturesInstrumentOpenInterestResult, error) {
	var openInterest FuturesInstrumentOpenInterestResult
	_, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_OPEN_INTEREST, InstrumentId), nil, &openInterest)
	return openInterest, err
}

/*
 Get the futures contract Instrument limit price
*/
func (client *Client) GetFuturesInstrumentPriceLimit(InstrumentId string) (FuturesInstrumentPriceLimitResult, error) {
	var priceLimit FuturesInstrumentPriceLimitResult
	_, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_PRICE_LIMIT, InstrumentId), nil, &priceLimit)
	return priceLimit, err
}

/*
 Get the futures contract liquidation
*/
func (client *Client) GetFuturesInstrumentLiquidation(InstrumentId string, status, from, to, limit int) (FuturesInstrumentLiquidationListResult, error) {
	var liquidation []FuturesInstrumentLiquidationResult
	params := NewParams()
	params["status"] = Int2String(status)
	params["from"] = Int2String(from)
	params["to"] = Int2String(to)
	params["limit"] = Int2String(limit)
	requestPath := BuildParams(GetInstrumentIdUri(FUTURES_INSTRUMENT_LIQUIDATION, InstrumentId), params)
	response, err := client.Request(GET, requestPath, nil, &liquidation)
	var list FuturesInstrumentLiquidationListResult
	page := parsePage(response)
	list.Page = page
	list.LiquidationList = liquidation
	return list, err
}

/*
 =============================== Futures trade api ===============================
*/

/*
 Get all of futures contract position list.
 return struct: FuturesPositions
*/
func (client *Client) GetFuturesPositions() (FuturesPosition, error) {
	response, err := client.Request(GET, FUTURES_POSITION, nil, nil)
	return parsePositions(response, err)
}

/*
 Get all of futures contract position list.
 return struct: FuturesPositions
*/
func (client *Client) GetFuturesInstrumentPosition(InstrumentId string) (FuturesPosition, error) {
	response, err := client.Request(GET, GetInstrumentIdUri(FUTURES_INSTRUMENT_POSITION, InstrumentId), nil, nil)
	return parsePositions(response, err)
}

/*
 Get all of futures contract account list
 return struct: FuturesAccounts
*/
func (client *Client) GetFuturesAccounts() (FuturesAccount, error) {
	response, err := client.Request(GET, FUTURES_ACCOUNTS, nil, nil)
	return parseAccounts(response, err)
}

/*
 Get the futures contract currency account @see file : futures_constants.go
 return struct: FuturesCurrencyAccounts
*/
func (client *Client) GetFuturesAccountsByCurrency(currency string) (FuturesCurrencyAccount, error) {
	response, err := client.Request(GET, GetCurrencyUri(FUTURES_ACCOUNT_CURRENCY_INFO, currency), nil, nil)
	return parseCurrencyAccounts(response, err)
}

/*
 Get the futures contract currency ledger
*/
func (client *Client) GetFuturesAccountsLedgerByCurrency(currency string, from, to, limit int) ([]FuturesCurrencyLedger, error) {
	var ledger []FuturesCurrencyLedger
	params := NewParams()
	params["from"] = Int2String(from)
	params["to"] = Int2String(to)
	params["limit"] = Int2String(limit)
	requestPath := BuildParams(GetCurrencyUri(FUTURES_ACCOUNT_CURRENCY_LEDGER, currency), params)
	_, err := client.Request(GET, requestPath, nil, &ledger)
	return ledger, err
}

/*
 Get the futures contract Instrument holds
*/
func (client *Client) GetFuturesAccountsHoldsByInstrumentId(InstrumentId string) (FuturesAccountsHolds, error) {
	var holds FuturesAccountsHolds
	_, err := client.Request(GET, GetInstrumentIdUri(FUTURES_ACCOUNT_INSTRUMENT_HOLDS, InstrumentId), nil, &holds)
	return holds, err