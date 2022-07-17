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
	if _, err := client.Request(GET, GetInstrumentIdUri(SWAP_INSTRUMENT_POSITION, instrumen