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
	InstrumentId     string `json:"inst