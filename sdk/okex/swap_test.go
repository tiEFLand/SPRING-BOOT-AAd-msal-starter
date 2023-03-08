
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