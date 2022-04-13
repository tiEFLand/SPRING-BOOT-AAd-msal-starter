
package okex

import "strings"

/*
币币账户信息
获取币币账户资产列表(仅展示拥有资金的币对)，查询各币种的余额、冻结和可用等信息。

限速规则：20次/2s
HTTP请求
GET /api/spot/v3/accounts

*/
func (client *Client) GetSpotAccounts() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, err := client.Request(GET, SPOT_ACCOUNTS, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
单一币种账户信息
获取币币账户单个币种的余额、冻结和可用等信息。

限速规则：20次/2s
HTTP请求
GET /api/spot/v3/accounts/<currency>
*/
func (client *Client) GetSpotAccountsCurrency(currency string) (*map[string]interface{}, error) {
	r := map[string]interface{}{}
	uri := GetCurrencyUri(SPOT_ACCOUNTS_CURRENCY, currency)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
账单流水查询
列出账户资产流水。账户资产流水是指导致账户余额增加或减少的行为。流水会分页，并且按时间倒序排序和存储，最新的排在最前面。请参阅分页部分以获取第一页之后的其他记录。

限速规则：20次/2s
HTTP请求
GET /api/spot/v3/accounts/<currency>/ledger
*/
func (client *Client) GetSpotAccountsCurrencyLeger(currency string, optionalParams *map[string]string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	baseUri := GetCurrencyUri(SPOT_ACCOUNTS_CURRENCY_LEDGER, currency)
	uri := baseUri
	if optionalParams != nil && len(*optionalParams) > 0 {
		uri = BuildParams(baseUri, *optionalParams)
	}

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取订单列表
列出您当前所有的订单信息。这个请求支持分页，并且按时间倒序排序和存储，最新的排在最前面。请参阅分页部分以获取第一页之后的其他纪录。

限速规则：20次/2s
HTTP请求
GET /api/spot/v3/orders
*/
func (client *Client) GetSpotOrders(status, instrument_id string, options *map[string]string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	fullOptions := NewParams()
	fullOptions["instrument_id"] = instrument_id
	fullOptions["status"] = status
	if options != nil && len(*options) > 0 {
		fullOptions["from"] = (*options)["from"]
		fullOptions["to"] = (*options)["to"]
		fullOptions["limit"] = (*options)["limit"]
	}

	uri := BuildParams(SPOT_ORDERS, fullOptions)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取所有未成交订单
列出您当前所有的订单信息。这个请求支持分页，并且按时间倒序排序和存储，最新的排在最前面。请参阅分页部分以获取第一页之后的其他纪录。

限速规则：20次/2s
HTTP请求
GET /api/spot/v3/orders_pending
*/
func (client *Client) GetSpotOrdersPending(options *map[string]string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	fullOptions := NewParams()
	uri := SPOT_ORDERS_PENDING
	if options != nil && len(*options) > 0 {
		fullOptions["instrument_id"] = (*options)["instrument_id"]
		fullOptions["from"] = (*options)["from"]
		fullOptions["to"] = (*options)["to"]
		fullOptions["limit"] = (*options)["limit"]
		uri = BuildParams(SPOT_ORDERS_PENDING, fullOptions)
	}

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取订单信息
通过订单ID获取单个订单信息。

限速规则：20次/2s

HTTP请求
GET /api/spot/v3/orders/<order_id>
或者
GET /api/spot/v3/orders/<client_oid>
*/
func (client *Client) GetSpotOrdersById(instrumentId, orderOrClientId string) (SpotGetOrderResult, error) {
	r := SpotGetOrderResult{}
	uri := strings.Replace(SPOT_ORDERS_BY_ID, "{order_client_id}", orderOrClientId, -1)
	options := NewParams()
	options["instrument_id"] = instrumentId
	uri = BuildParams(uri, options)

	_, err := client.Request(GET, uri, nil, &r)
	return r, err
}

/*
获取成交明细
获取最近的成交明细表。这个请求支持分页，并且按时间倒序排序和存储，最新的排在最前面。请参阅分页部分以获取第一页之后的其他记录。

限速规则：20次/2s
HTTP请求
GET /api/spot/v3/fills
*/
func (client *Client) GetSpotFills(order_id, instrument_id string, options *map[string]string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	fullOptions := NewParams()
	fullOptions["instrument_id"] = instrument_id
	fullOptions["order_id"] = order_id
	if options != nil && len(*options) > 0 {
		fullOptions["from"] = (*options)["from"]
		fullOptions["to"] = (*options)["to"]
		fullOptions["limit"] = (*options)["limit"]
	}

	uri := BuildParams(SPOT_FILLS, fullOptions)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取币对信息
用于获取行情数据，这组公开接口提供了行情数据的快照，无需认证即可调用。

获取交易币对的列表，查询各币对的交易限制和价格步长等信息。

限速规则：20次/2s
HTTP请求
GET /api/spot/v3/instruments
*/
func (client *Client) GetSpotInstruments() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, err := client.Request(GET, SPOT_INSTRUMENTS, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取深度数据
获取币对的深度列表。这个请求不支持分页，一个请求返回整个深度列表。

限速规则：20次/2s
HTTP请求
GET /api/spot/v3/instruments/<instrument_id>/book
*/
func (client *Client) GetSpotInstrumentBook(instrumentId string, optionalParams *map[string]string) (*map[string]interface{}, error) {
	r := map[string]interface{}{}
	uri := GetInstrumentIdUri(SPOT_INSTRUMENT_BOOK, instrumentId)
	if optionalParams != nil && len(*optionalParams) > 0 {
		optionals := NewParams()
		optionals["size"] = (*optionalParams)["size"]
		optionals["depth"] = (*optionalParams)["depth"]
		uri = BuildParams(uri, optionals)
	}

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取全部ticker信息
获取平台全部币对的最新成交价、买一价、卖一价和24小时交易量的快照信息。

限速规则：50次/2s
HTTP请求
GET /api/spot/v3/instruments/ticker
*/
func (client *Client) GetSpotInstrumentsTicker() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, err := client.Request(GET, SPOT_INSTRUMENTS_TICKER, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
