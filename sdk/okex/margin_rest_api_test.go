
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
