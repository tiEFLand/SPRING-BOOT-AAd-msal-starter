package okex

/*
 OKEX Swap Api
 @author Lingting Fu
 @date 2018-12-27
 @version 1.0.0
*/

import (
	"errors"
	"strings"
)

/*
获取某个合约的持仓信息
GET /api/swap/v3/<instrument_id>/position
*/
func (client *Client) GetSwapPositionByInstrument(instrumentId string) (*SwapPosition, error) {

	sp := SwapPosition{}
	if _, err := client.Request(GET, GetInstrumentIdUri(SWAP_INSTRUMENT_POSITION, instrumentId), nil, &sp); err != nil {
		return nil, err
	}
	return &sp, nil
}

/*
所有合约持仓信息
获取所有合约的持仓信息
限速规则：1次/10s
GET /api/swap/v3/position
*/
func (client *Client) GetSwapPositions() (*SwapPositionList, error) {

	sp := SwapPositionList{}
	if _, err := client.Request(GET, SWAP_POSITION, nil, &sp); err != nil {
		return nil, err
	}
	return &sp, nil
}

func (client *Client) getSwapAccounts(uri string) (*SwapAccounts, error) {
	sa := SwapAccounts{}
	if _, err := client.Request(GET, uri, nil, &sa); err != nil {
		return nil, err
	}
	return &sa, nil
}

/*
获取所有币种合约的账户信息
HTTP请求
GET /api/swap/v3/accounts
*/
func (client *Client) GetSwapAccounts() (*SwapAccounts, error) {
	return client.getSwapAccounts(SWAP_ACCOUNTS)
}

/*
单个币种合约账户信息
HTTP请求
GET /api/swap/v3/<instrument_id>/accounts
*/
func (client *Client) GetSwapAccount(instrumentId string) (*SwapAccount, error) {

	sa := SwapAccount{}
	uri := GetInstrumentIdUri(SWAP_INSTRUMENT_ACCOUNT, instrumentId)
	if _, err := client.Request(GET, uri, nil, &sa); err != nil {
		return nil, err
	}
	return &sa, nil
}

/*
获取某个合约的杠杆倍数，持仓模式

HTTP请求
GET /api/swap/v3/accounts/<instrument_id>/settings
*/
func (client *Client) GetSwapAccountsSettingsByInstrument(instrumentId string) (*SwapAccountsSetting, error) {
	as := SwapAccountsSetting{}
	if _, err := client.Request(GET, GetInstrumentIdUri(SWAP_ACCOUNTS_SETTINGS, instrumentId), nil, &as); err != nil {
		return nil, err
	}
	return &as, nil
}

/*
设定某个合约的杠杆倍数

HTTP请求
POST /api/swap/v3/accounts/<instrument_id>/leverage
*/
func (client *Client) PostSwapAccountsLeverage(instrumentId string, leverage string, side string) (*SwapAccountsSetting, error) {
	params := make(map[string]string)
	params["leverage"] = leverage
	params["side"] = side
	as := SwapAccountsSetting{}
	if _, err := client.Request(POST, GetInstrumentIdUri(SWAP_ACCOUNTS_LEVERAGE, instrumentId), params, &as); err != nil {
		return nil, err
	}
	return &as, nil
}

/*
账单流水查询
列出账户资产流水，账户资产流水是指导致账户余额增加或减少的行为。流水会分页，每页100条数据，并且按照时间倒序排序和存储，最新的排在最前面。

HTTP请求
GET /api/swap/v3/accounts/<instrument_id>/ledger
*/
func (client *Client) GetSwapAccountLedger(instrumentId string, optionalParams map[string]string) (*SwapAccountsLedgerList, error) {
	baseUri := GetInstrumentIdUri(SWAP_ACCOUNTS_LEDGER, instrumentId)
	uri := baseUri
	if optionalParams != nil {
		uri = BuildParams(baseUri, optionalParams)
	}
	ll := SwapAccountsLedgerList{}
	if _, err := client.Request(GET, uri, nil, &ll); err != nil {
		return nil, err
	}
	return &ll, nil
}

/*
API交易提供限价单下单模式，只有当您的账户有足够的资金才能下单。一旦下单，您的账户资金将在订单生命周期内被冻结，被冻结的资金以及数量取决于订单指定的类型和参数。

HTTP请求
POST /api/swap/v3/order
*/
func (client *Client) PostSwapOrder(instrumentId string, order *BasePlaceOrderInfo) (*SwapOrderResult, error) {
	or := SwapOrderResult{}
	info := PlaceOrderInfo{*order, instrumentId}
	if _, err := client.Request(POST, SWAP_ORDER, info, &or); err != nil {
		return nil, err
	}
	return &or, nil
}

/*
批量进行下单请求。

HTTP请求
POST /api/swap/v3/orders
*/
func (client *Client) PostSwapOrders(instrumentId string, orders []*BasePlaceOrderInfo) (*SwapOrdersResult, error) {
	sor := SwapOrdersResult{}
	orderData := PlaceOrdersInfo{InstrumentId: instrumentId, OrderData: orders}
	if _, err := client.Request(POST, SWAP_ORDERS, orderData, &sor); err != nil {
		return nil, err
	}
	return &sor, nil
}

/*
撤销之前下的未完成订单。

HTTP请求
POST /api/swap/v3/cancel_order/<instrument_id>/<order_id>
*/
func (client *Client) PostSwapCancelOrder(instrumentId string, orderId string) (*SwapCancelOrderResult, error) {
	uri := "/api/swap/v3/cancel_order/" + instrumentId + "/" + orderId
	or := SwapCancelOrderResult{}
	if _, err := client.Request(POST, uri, nil, &or); err != nil {
		return nil, err
	}
	return &or, nil

}

/*
批量撤销之前下的未完成订单。

HTTP请求
POST /api/swap/v3/cancel_batch_orders/<instrument_id>
*/
func (client *Client) PostSwapBatchCancelOrders(instrumentId string, orderIds []string) (*SwapCancelOrderResult, error) {
	uri := GetInstrumentIdUri(SWAP_CANCEL_BATCH_ORDERS, instrumentId)
	or := SwapCancelOrderResult{}

	params := map[string]interface{}{}
	params["ids"] = orderIds

	if _, err := client.Request(POST, uri, params, &or); err != nil {
		return nil, err
	}
	return &or, nil
}

/*
列出您当前所有的订单信息。

HTTP请求
GET /api/swap/v3/orders/<instrument_id>

请求示例
GET /api/swap/v3/orders/BTC-USD-SWAP?status=2&from=4&limit=30
*/
func (client *Client) GetSwapOrderByInstrumentId(instrumentId string, paramMap map[string]string) (*SwapOrdersInfo, error) {
	if paramMap["status"] == "" || len(instrumentId) == 0 {
		return nil, errors.New("Request Parameter's not correct, instrument_id and status is required.")
	}

	baseUri := GetInstrumentIdUri(SWAP_INSTRUMENT_ORDER_LIST, instrumentId)
	kvParams := BuildOrderParams(paramMap)
	uri := baseUri + "?" + kvParams
	soi := SwapOrdersInfo{}

	if _, err := client.Request(GET, uri, nil, &soi); err != nil {
		return nil, err
	}
	return &soi, nil
}

/*
通过订单id获取单个订单信息。

HTTP请求
GET 