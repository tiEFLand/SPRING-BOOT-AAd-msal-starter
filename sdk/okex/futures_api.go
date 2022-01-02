
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