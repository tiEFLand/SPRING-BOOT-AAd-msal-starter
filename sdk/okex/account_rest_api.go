
package okex

/*
获取平台所有币种列表。并非所有币种都可被用于交易。在ISO 4217标准中未被定义的币种代码可能使用的是自定义代码。

HTTP请求
GET /api/account/v3/currencies

*/
func (client *Client) GetAccountCurrencies() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, err := client.Request(GET, ACCOUNT_CURRENCIES, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
钱包账户信息
获取钱包账户所有资产列表，查询各币种的余额、冻结和可用等信息。

HTTP请求
GET /api/account/v3/wallet
*/
func (client *Client) GetAccountWallet() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, err := client.Request(GET, ACCOUNT_WALLET, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
单一币种账户信息
获取钱包账户单个币种的余额、冻结和可用等信息。

HTTP请求
GET /api/account/v3/wallet/<currency>

请求示例
GET /api/account/v3/wallet/btc
*/
func (client *Client) GetAccountWalletByCurrency(currency string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := GetCurrencyUri(ACCOUNT_WALLET_CURRENCY, currency)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
单一币种账户信息
获取钱包账户单个币种的余额、冻结和可用等信息。

HTTP请求
GET /api/account/v3/wallet/<currency>
*/
func (client *Client) GetAccountWithdrawalFeeByCurrency(currency *string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := ACCOUNT_WITHRAWAL_FEE
	if currency != nil && len(*currency) > 0 {
		params := NewParams()
		params["currency"] = *currency
		uri = BuildParams(uri, params)
	}

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
查询最近所有币种的提币记录

HTTP请求
GET /api/account/v3/withdrawal/history
*/
func (client *Client) GetAccountWithdrawalHistory() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, err := client.Request(GET, ACCOUNT_WITHRAWAL_HISTORY, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
查询单个币种的提币记录。

HTTP请求
GET /api/account/v3/withdrawal/history/<currency>
*/
func (client *Client) GetAccountWithdrawalHistoryByCurrency(currency string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	uri := GetCurrencyUri(ACCOUNT_WITHRAWAL_HISTORY_CURRENCY, currency)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
获取充值地址
获取各个币种的充值地址，包括曾使用过的老地址。

HTTP请求
GET /api/account/v3/deposit/address

请求示例
GET /api/account/v3/deposit/address?currency=btc
*/
func (client *Client) GetAccountDepositAddress(currency string) (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}
	params := NewParams()
	params["currency"] = currency

	uri := BuildParams(ACCOUNT_DEPOSIT_ADDRESS, params)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}