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
	AvailPosition    stri