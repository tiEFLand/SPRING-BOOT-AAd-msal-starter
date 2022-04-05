
package okex

/*
 OKEX futures contract api response results
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

type ServerTime struct {
	Iso   string `json:"iso"`
	Epoch string `json:"epoch"`
}

type ExchangeRate struct {
	InstrumentId string  `json:"instrument_id"`
	Rate         float64 `json:"rate,string"`
	Timestamp    string  `json:"timestamp"`
}

type BizWarmTips struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Msg     string `json:"msg"`
}

type Result struct {
	Result bool `json:"result"`
}

type PageResult struct {
	From  int
	To    int
	Limit int
}

type FuturesPosition struct {
	BizWarmTips
	Result
	MarginMode    string
	CrossPosition []FuturesCrossPositionHolding
	FixedPosition []FuturesFixedPositionHolding
}

type FuturesCrossPosition struct {
	Result
	MarginMode    string                        `json:"margin_mode"`
	CrossPosition []FuturesCrossPositionHolding `json:"holding"`
}

type FuturesFixedPosition struct {
	Result
	MarginMode    string                        `json:"margin_mode"`
	FixedPosition []FuturesFixedPositionHolding `json:"holding"`
}

type FuturesCrossPositionHolding struct {
	FuturesPositionBase
	LiquidationPrice float64 `json:"liquidation_price,string"`
	Leverage         float64 `json:"leverage,string"`
}

type FuturesFixedPositionHolding struct {
	FuturesPositionBase
	LongMargin      float64 `json:"long_margin,string"`
	LongLiquiPrice  float64 `json:"long_liqui_price,string"`
	LongPnlRatio    float64 `json:"long_pnl_ratio,string"`
	LongLeverage    float64 `json:"long_leverage,string"`
	ShortMargin     float64 `json:"short_margin,string"`
	ShortLiquiPrice float64 `json:"short_liqui_price,string"`
	ShortPnlRatio   float64 `json:"short_pnl_ratio,string"`
	ShortLeverage   float64 `json:"short_leverage,string"`
}

type FuturesPositionBase struct {
	LongQty              float64 `json:"long_qty,string"`
	LongAvailQty         float64 `json:"long_avail_qty,string"`
	LongAvgCost          float64 `json:"long_avg_cost,string"`
	LongSettlementPrice  float64 `json:"long_settlement_price,string"`
	RealizedPnl          float64 `json:"realized_pnl,string"`
	ShortQty             float64 `json:"short_qty,string"`
	ShortAvailQty        float64 `json:"short_avail_qty,string"`
	ShortAvgCost         float64 `json:"short_avg_cost,string"`
	ShortSettlementPrice float64 `json:"short_settlement_price,string"`
	InstrumentId         string  `json:"instrument_id"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
}

type FuturesAccount struct {
	BizWarmTips
	Result
	MarginMode   string
	CrossAccount map[string]FuturesCrossAccount
	FixedAccount map[string]FuturesFixedAccount
}

type FuturesMarkdown struct {
	BizWarmTips
	InstrumentId string  `json:"instrument_id"`
	Timestamp    string  `json:"timestamp"`
	MarkPrice    float32 `json:"mark_price"`
}

type FuturesFixedAccountInfo struct {
	Result
	Info map[string]FuturesFixedAccount `json:"info"`
}

type FuturesCrossAccountInfo struct {
	Result
	Info map[string]FuturesCrossAccount `json:"info"`
}

type FuturesFixedAccount struct {
	MarginMode        string                         `json:"margin_mode"`
	Equity            float64                        `json:"equity,string"`
	TotalAvailBalance float64                        `json:"total_avail_balance,string"`
	Contracts         []FuturesFixedAccountContracts `json:"contracts"`
}

type FuturesFixedAccountContracts struct {
	AvailableQty      float64 `json:"available_qty,string"`
	FixedBalance      float64 `json:"fixed_balance,string"`
	InstrumentId      string  `json:"instrument_id"`
	MarginFixed       float64 `json:"margin_fixed,string"`
	MarginForUnfilled float64 `json:"margin_for_unfilled,string"`
	MarginFrozen      float64 `json:"margin_frozen,string"`
	RealizedPnl       float64 `json:"realized_pnl,string"`
	UnrealizedPnl     float64 `json:"unrealizedPnl,string"`
}

type FuturesCrossAccount struct {
	Equity            float64 `json:"equity,string"`
	Margin            float64 `json:"margin,string"`
	MarginMode        string  `json:"margin_mode"`
	MarginRatio       float64 `json:"margin_ratio,string"`
	RealizedPnl       float64 `json:"realized_pnl,string"`
	UnrealizedPnl     float64 `json:"unrealized_pnl,string"`
	TotalAvailBalance float64 `json:"total_avail_balance,string"`
}

type FuturesCurrencyAccount struct {
	BizWarmTips
	Result
	MarginMode   string
	CrossAccount FuturesCrossAccount
	FixedAccount FuturesFixedAccount
}

type FuturesCurrencyLedger struct {
	LedgerId  int64                        `json:"ledger_id,string"`
	Amount    float64                      `json:"amount,string"`
	Balance   float64                      `json:"balance,string"`
	Currency  string                       `json:"currency"`
	Type      string                       `json:"type"`
	Timestamp string                       `json:"timestamp"`
	Details   FuturesCurrencyLedgerDetails `json:"details"`
}

type FuturesCurrencyLedgerDetails struct {
	OrderId      int64  `json:"order_id"`
	InstrumentId string `json:"instrument_id"`
}

type FuturesAccountsHolds struct {
	InstrumentId string  `json:"instrument_id"`
	Amount       float64 `json:"amount,string"`
	Timestamp    string  `json:"timestamp"`
}

type FuturesNewOrderResult struct {
	BizWarmTips
	Result
	ClientOid string `json:"client_oid"`
	OrderId   string `json:"order_id"`
}

type FuturesBatchNewOrderResult struct {
	Result
	OrderInfo []OrderInfo `json:"order_info"`
}
