
package okex

/*
 OKEX Swap Rest Api. Unit test
 @author Lingting Fu
 @date 2018-12-27
 @version 1.0.0
*/

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestGetSwapInstrumentPosition(t *testing.T) {
	c := NewTestClient()
	sp, err := c.GetSwapPositionByInstrument("BTC-USD-SWAP")
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(sp)
	println(jstr)
}

func TestGetSwapPositions(t *testing.T) {
	c := NewTestClient()
	sp, err := c.GetSwapPositions()
	simpleAssertTrue(sp, err, t, false)
}

func TestGetSwapAccount(t *testing.T) {
	c := NewTestClient()
	sa, err := c.GetSwapAccount("BTC-USD-SWAP")
	simpleAssertTrue(sa, err, t, false)
}

type JsonTestType struct {
	Asks [][]interface{} `json:"asks"`
}

func TestJson(t *testing.T) {
	a := string(`{"asks" : [["411.5","9",4,3]]}`)
	i := JsonTestType{}
	JsonString2Struct(a, &i)
	simpleAssertTrue(i, nil, t, false)

	str, _ := Struct2JsonString(i)
	println(str)
}

func TestMap(t *testing.T) {
	m := map[string]string{}
	m["a"] = "1999"

	r := m["b"]
	assert.True(t, r == "" && len(r) == 0)

	r2 := m["a"]
	assert.True(t, r2 == "1999")
}

func TestClient_PublicAPI(t *testing.T) {
	c := NewTestClient()
	instrumentId := "BTC-USD-SWAP"
	histList, err := c.GetSwapHistoricalFundingRateByInstrument(instrumentId, nil)
	fmt.Printf("%+v", err)
	assert.True(t, histList != nil && err == nil)
	fmt.Printf("%+v", histList)

	r, err := c.GetSwapMarkPriceByInstrument(instrumentId)
	fmt.Printf("Result: %+v, Error: %+v", r, err)
	assert.True(t, r != nil && err == nil)

	r1, err := c.GetSwapFundingTimeByInstrument(instrumentId)
	fmt.Printf("Result: %+v, Error: %+v", r1, err)
	assert.True(t, r1 != nil && err == nil)

	r2, err := c.GetSwapAccountsHoldsByInstrument(instrumentId)
	fmt.Printf("Result: %+v, Error: %+v", r2, err)
	assert.True(t, r2 != nil && err == nil)

	r3, err := c.GetSwapLiquidationByInstrument(instrumentId, "1", nil)
	fmt.Printf("Result: %+v, Error: %+v", r3, err)
	assert.True(t, r3 != nil && err == nil)

	optionalParams := map[string]string{}
	optionalParams["from"] = "1"
	optionalParams["to"] = "5"
	optionalParams["limit"] = "50"

	r4, err := c.GetSwapLiquidationByInstrument(instrumentId, "0", optionalParams)
	fmt.Printf("Result: %+v, Error: %+v", r4, err)
	assert.True(t, r4 != nil && err == nil)

	r5, err := c.GetSwapPriceLimitByInstrument(instrumentId)
	fmt.Printf("Result: %+v, Error: %+v", r5, err)
	assert.True(t, r5 != nil && err == nil)

	r6, err := c.GetSwapOpenInterestByInstrument(instrumentId)
	fmt.Printf("Result: %+v, Error: %+v", r6, err)
	assert.True(t, r6 != nil && err == nil)

	r7, err := c.GetSwapIndexByInstrument(instrumentId)
	fmt.Printf("Result: %+v, Error: %+v", r7, err)
	assert.True(t, r7 != nil && err == nil)

	//lingting.fu. 20190225. No Kline in test enviroment, contact LiWei to solve env problem.
	//r8, err := c.GetSwapCandlesByInstrument(instrumentId, nil)
	//fmt.Printf("Result: %+v, Error: %+v", r8, err)
	//assert.True(t, r8 != nil && err == nil)

	r9, err := c.GetSwapTradesByInstrument(instrumentId, nil)
	fmt.Printf("Result: %+v, Error: %+v", r9, err)
	assert.True(t, r9 != nil && err == nil)

	r10, err := c.GetSwapTickerByInstrument(instrumentId)
	fmt.Printf("Result: %+v, Error: %+v", r10, err)
	assert.True(t, r10 != nil && err == nil)

	r11, err := c.GetSwapInstruments()
	fmt.Printf("Result: %+v, Error: %+v", r11, err)
	assert.True(t, r11 != nil && err == nil)

	r12, err := c.GetSwapRate()
	fmt.Printf("Result: %+v, Error: %+v", r12, err)
	assert.True(t, r12 != nil && err == nil)

	r13, err := c.GetSwapInstrumentsTicker()
	simpleAssertTrue(r13, err, t, false)

	r14, err := c.GetSwapDepthByInstrumentId(instrumentId, "1")
	simpleAssertTrue(r14, err, t, false)

}

func simpleAssertTrue(result interface{}, err error, t *testing.T, doprint bool) bool {
	if doprint {
		fmt.Fprintf(os.Stderr, "Result: %+v, Error: %+v", result, err)
	}
	assert.True(t, result != nil && err == nil)
	return result != nil && err == nil
}

func TestClient_PrivateAPI(t *testing.T) {

	c := NewTestClient()
	instrumentId := "BTC-USD-SWAP"

	// Fore. 20190225. CleanUp history test order before new test start.
	cleanUpOrders(c, instrumentId)

	r1, err := c.GetSwapPositionByInstrument(instrumentId)
	simpleAssertTrue(r1, err, t, false)

	r2, err := c.GetSwapAccounts()
	simpleAssertTrue(r2, err, t, false)

	r3, err := c.GetSwapAccountsSettingsByInstrument(instrumentId)
	simpleAssertTrue(r3, err, t, false)

	r4, err := c.PostSwapAccountsLeverage(instrumentId, "20", "3")
	simpleAssertTrue(r4, err, t, false)

	r5, err := c.PostSwapAccountsLeverage(instrumentId, "50", "3")
	simpleAssertTrue(r5, err, t, false)