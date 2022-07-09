
package okex

type SpotGetOrderResult struct {
	OrderId        string  `json:"order_id"`
	Price          string  `json:"price"`
	Size           float64 `json:"size,string"`
	Notional       string  `json:"notional"`
	InstrumentId   string  `json:"instrument_id"`
	Side           string  `json:"side"`
	Type           string  `json:"type"`
	Timestamp      string  `json:"timestamp"`
	FilledSize     float64 `json:"filled_size,string"`
	FilledNotional float64 `json:"filled_notional,string"`