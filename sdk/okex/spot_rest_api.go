
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