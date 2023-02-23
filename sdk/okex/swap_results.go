package okex

/*
 OKEX api result definition
 @author Lingting Fu
 @date 2018-12-27
 @version 1.0.0
*/

type SwapPositionHolding struct {
	LiquidationPrice string `json:"liquidation_price"`
	Position         string `json:"position"`
	AvailPosition    string `json:"avail_position"`
	AvgCost          string `json:"avg_cost"`
	SettlementPrice  string `json:"settlement_price"`
	InstrumentId     string `json:"instrument_id"`
	Leverage         string `json:"leverage"`
	RealizedPnl      string `json:"realized_pnl"`
	Side             string `json:"side"`
	Timestamp        string `json:"timestamp"`
	Margin           string `json:"margin";default:""`
}

type SwapPosition struct {
	BizWarmTips
	MarginMode string                `json:"margin_mode"`
	Holding    []SwapPositionHolding `json:"holding"`
}

type SwapPositionList []SwapPosition

type SwapAccountInfo struct {
	InstrumentId      string `json:"instrument_id"`
	Timestamp         string `json:"timestamp"`
	MarginFrozen      string `json:"margin_frozen"`
	TotalAvailBalance string `json:"total_avail_balance"`
	MarginRatio       string `json:"margin_ratio"`
	RealizedPnl       string `json:"realized_pnl"`
	UnrealizedPnl     string `json:"unrealized_pnl"`
	FixedBalance      string `json:"fixed_balance"`
	Equity            string `json:"equity"`
	Margin            string `json:"margin"`
	MarginMode        string `json:"margin_mode"`
}

type SwapAccounts struct {
	BizWarmTips
	Info []SwapAccountInfo `json:"info"`
}

type SwapAccount struct {
	Info SwapAccountInfo `json:"info"`
}

type BaseSwapOrderResult struct {
	OrderId      string `json:"order_id"`
	ClientOid    string `json:"client_oid"`
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
	Result       string `json:"result"`
}

type SwapOrderResult struct {
	BaseSwapOrderResult
	BizWarmTips
}

type SwapOrdersResult struct {
	BizWarmTips
	OrderInfo []BaseSwapOrderResult `json:"order_info"`
}

type SwapCancelOrderResult struct {
	Err