
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