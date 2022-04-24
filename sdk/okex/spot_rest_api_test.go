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
	c :=