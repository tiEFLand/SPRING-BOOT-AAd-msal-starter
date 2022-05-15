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
	ac, err := c.GetSpotOrders("filled", "BTC-USDT", nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

	// Fore. 20190305. TODO: {"message":"System error"} returned by following request.
	// Url: http://coinmainweb.new.docker.okex.com/api/spot/v3/fills?instrument_id=BTC-USDT&order_id=2365709152770048
	filledOrderId := (*ac)[0]["order_id"].(string)
	sf, err := c.GetSpotFills(filledOrderId, "BTC-USDT", nil)
	assert.True(t, sf != nil && err == nil)
}

func TestGetSpotOrdersPending(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetSpotOrdersPending(nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

	options := NewParams()
	options["instrument_id"] = "BTC-USDT"
	ac, err = c.GetSpotOrdersPending(&options)
	assert.True(t, err == nil)
	jstr, _ = Struct2JsonString(ac)
	println(jstr)

	testOrderId := (*ac)[0]["order_id"]
	so, err := c.GetSpotOrdersById("BTC-USDT", testOrderId.(string))
	assert.True(t, so != nil