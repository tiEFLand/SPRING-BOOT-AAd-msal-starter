package okex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSpotAccounts(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotAccounts()

	fmt.Printf("%+v, %+v", ac, err)

	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetSpotAccountsCurrency(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotAccountsCurrency("BTC")
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetSpotAccountsCurrencyLeger(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotAccountsCurrencyLeger("btc", nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

	options := map[string]string{}
	options["from"] = "1"
	options["to"] = "2"
	options["limit"] = "100"

	ac2, err2 := c.GetSpotAccountsCurrencyLeger("btc", &options)
	assert.True(t, ac2 != nil && err2 == nil)
}

func TestGetSpotOrders(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpo