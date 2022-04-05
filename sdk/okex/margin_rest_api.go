
package okex

import (
	"strings"
)

/*
币币杠杆账户信息
获取币币杠杆账户资产列表，查询各币种的余额、冻结和可用等信息。

限速规则：20次/2s
HTTP请求
GET /api/margin/v3/accounts
*/
func (client *Client) GetMarginAccounts() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, err := client.Request(GET, MARGIN_ACCOUNTS, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
单一币对账户信息
获取币币杠杆某币对账户的余额、冻结和可用等信息。

限速规则：20次/2s
HTTP请求
GET /api/margin/v3/accounts/<instrument_id>
*/
func (client *Client) GetMarginAccountsByInstrument(instrumentId string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := GetInstrumentIdUri(MARGIN_ACCOUNTS_INSTRUMENT, instrumentId)
	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
账单流水查询
列出杠杆帐户资产流水。帐户资产流水是指导致帐户余额增加或减少的行为。流水会分页，并且按时间倒序排序和存储，最新的排在最前面。请参阅分页部分以获取第一页之后的其他纪录。

限速规则：20次/2s
HTTP请求
GET /api/margin/v3/accounts/<instrument_id>/ledger
*/
func (client *Client) GetMarginAccountsLegerByInstrument(instrumentId string, optionalParams *map[string]string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}
	uri := GetInstrumentIdUri(MARGIN_ACCOUNTS_INSTRUMENT_LEDGER, instrumentId)
	if optionalParams != nil && len(*optionalParams) > 0 {
		uri = BuildParams(uri, *optionalParams)
	}

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
杠杆配置信息
获取币币杠杆账户的借币配置信息，包括当前最大可借、借币利率、最大杠杆倍数。

限速规则：20次/2s
HTTP请求
GET /api/margin/v3/accounts/availability
*/
func (client *Client) GetMarginAccountsAvailability() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, err := client.Request(GET, MARGIN_ACCOUNTS_AVAILABILITY, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
某个杠杆配置信息
获取某个币币杠杆账户的借币配置信息，包括当前最大可借、借币利率、最大杠杆倍数。

限速规则：20次/2s
HTTP请求
GET /api/margin/v3/accounts/<instrument_id>/availability
*/
func (client *Client) GetMarginAccountsAvailabilityByInstrumentId(instrumentId string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := GetInstrumentIdUri(MARGIN_ACCOUNTS_INSTRUMENT_AVAILABILITY, instrumentId)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取借币记录
获取币币杠杆帐户的借币记录。这个请求支持分页，并且按时间倒序排序和存储，最新的排在最前面。请参阅分页部分以获取第一页之后的其他纪录。

限速规则：20次/2s
HTTP请求
GET /api/margin/v3/accounts/borrowed
*/
func (client *Client) GetMarginAccountsBorrowed(optionalParams *map[string]string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := MARGIN_ACCOUNTS_BORROWED
	if optionalParams != nil {
		uri = BuildParams(uri, *optionalParams)
	}
	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
某账户借币记录
获取币币杠杆帐户某币对的借币记录。这个请求支持分页，并且按时间倒序排序和存储，最新的排在最前面。请参阅分页部分以获取第一页之后的其他纪录。

限速规则：20次/2s
HTTP请求
GET /api/margin/v3/accounts/<instrument_id>/borrowed
*/
func (client *Client) GetMarginAccountsBorrowedByInstrumentId(instrumentId string, optionalParams *map[string]string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := GetInstrumentIdUri(MARGIN_ACCOUNTS_INSTRUMENT_BORROWED, instrumentId)
	if optionalParams != nil && len(*optionalParams) > 0 {
		uri = BuildParams(uri, *optionalParams)
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
GET /api/margin/v3/orders
*/
func (client *Client) GetMarginOrders(instrumentId string, optionalParams *map[string]string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}
	fullParams := NewParams()
	fullParams["instrument_id"] = instrumentId

	if optionalParams != nil && len(*optionalParams) > 0 {
		for k, v := range *optionalParams {
			fullParams[k] = v
		}
	}

	uri := BuildParams(MARGIN_ORDERS, fullParams)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取订单信息