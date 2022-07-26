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
GET /api/swap/v3/<in