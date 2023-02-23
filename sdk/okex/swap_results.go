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

type SwapPositionList []SwapPosit