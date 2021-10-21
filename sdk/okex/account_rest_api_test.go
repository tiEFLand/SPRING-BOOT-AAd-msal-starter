
package okex

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAccountCurrencies(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountCurrencies()

	fmt.Printf("%+v, %+v", ac, err)

	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetAccountWallet(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountWallet()
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetAccountWithdrawalFeeByCurrency(t *testing.T) {
	c := NewTestClient()