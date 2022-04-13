
package okex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	CurrencyPairInstrument = "BTC-USDT"
)

func TestGetMarginAccounts(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginAccounts()
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetMarginAccountsByInstrument(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginAccountsByInstrument(CurrencyPairInstrument)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetMarginAccountsLegerByInstrument(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginAccountsLegerByInstrument(CurrencyPairInstrument, nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetMarginAccountsAvailability(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginAccountsAvailability()
	assert.True(t, err == nil && ac != nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetMarginAccountsAvailabilityByInstrumentId(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginAccountsAvailabilityByInstrumentId(CurrencyPairInstrument)
	assert.True(t, err == nil && ac != nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetMarginAccountsBorrowed(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginAccountsBorrowed(nil)
	assert.True(t, err == nil && ac != nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetMarginAccountsBorrowedByInstrumentId(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginAccountsBorrowedByInstrumentId(CurrencyPairInstrument, nil)
	assert.True(t, err == nil && ac != nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetMarginOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginOrders(CurrencyPairInstrument, nil)
	assert.True(t, err == nil && ac != nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetMarginOrdersById(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginOrdersById(CurrencyPairInstrument, "123456")
	assert.True(t, err == nil && ac != nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetMarginOrdersPending(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginOrdersPending(nil)
	assert.True(t, err == nil && ac != nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetMarginFills(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetMarginFills(CurrencyPairInstrument, "23239", nil)
	assert.True(t, err == nil && ac != nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestPostMarginAccountsBorrow(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostMarginAccountsBorrow(CurrencyPairInstrument, "usdt", "0.1")
	assert.True(t, err == nil && ac != nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestPostMarginAccountsRepayment(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostMarginAccountsRepayment(CurrencyPairInstrument, "usdt", "0.1", nil)
	assert.True(t, err == nil && ac != nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestPostMarginOrders_AllInOne(t *testing.T) {

	c := NewTestClient()
	optionals := NewParams()
	optionals["type"] = "limit"
	optionals["price"] = "100"
	optionals["size"] = "0.01"
	optionals["margin_trading"] = "2"

	r, err := c.PostMarginOrders("sell", CurrencyPairInstrument, &optionals)
	assert.True(t, r != nil && err == nil)
	jstr, _ := Struct2JsonString(r)
	println(jstr)

	orderId := (*r)["order_id"].(string)
	r, err = c.PostMarginCancelOrdersById(CurrencyPairInstrument, orderId)
	assert.True(t, r != nil && err == nil)
	jstr, _ = Struct2JsonString(r)
	println(jstr)

}

func TestPostMarginBatchOrders(t *testing.T) {
	c := NewTestClient()

	orderInfos := []map[string]string{
		map[string]string{"client_oid": "w20180728w", "instrument_id": "btc-usdt", "side": "sell", "type": "limit", "size": "0.001", "price": "10001", "margin_trading ": "1"},
		map[string]string{"client_oid": "r20180728r", "instrument_id": "btc-usdt", "side": "sell", "type": "limit", " size ": "0.001", "notional": "10002", "margin_trading ": "1"},
	}

	r, err := c.PostMarginBatchOrders(&orderInfos)
	assert.True(t, r != nil && err == nil)
	jstr, _ := Struct2JsonString(r)
	println(jstr)
}