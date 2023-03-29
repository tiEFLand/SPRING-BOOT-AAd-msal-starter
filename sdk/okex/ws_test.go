
package okex

/*
 OKEX ws api websocket test & sample
 @author Lingting Fu
 @date 2018-12-27
 @version 1.0.0
*/

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hash/crc32"
	"testing"
	"time"
)

func TestOKWSAgent_AllInOne(t *testing.T) {
	agent := OKWSAgent{}
	config := GetDefaultConfig()

	// Step1: Start agent.
	agent.Start(config)

	// Step2: Subscribe channel
	// Step2.0: Subscribe public channel swap/ticker successfully.
	agent.Subscribe(CHNL_SWAP_TICKER, "BTC-USD-SWAP", DefaultDataCallBack)

	// Step2.1: Subscribe private channel swap/position before login, so it would be a fail.
	agent.Subscribe(CHNL_SWAP_POSITION, "BTC-USD-SWAP", DefaultDataCallBack)

	// Step3: Wait for the ws server's pushed table responses.
	time.Sleep(60 * time.Second)

	// Step4. Unsubscribe public channel swap/ticker
	agent.UnSubscribe(CHNL_SWAP_TICKER, "BTC-USD-SWAP")
	time.Sleep(1 * time.Second)

	// Step5. Login
	agent.Login(config.ApiKey, config.Passphrase)
	time.Sleep(1 * time.Second)

	// Step6. Subscribe private channel swap/position after login, so it would be a success.
	agent.Subscribe(CHNL_SWAP_POSITION, "BTC-USD-SWAP", DefaultDataCallBack)
	time.Sleep(120 * time.Second)

	// Step7. Stop all the go routine run in background.
	agent.Stop()
	time.Sleep(1 * time.Second)
}

func TestOKWSAgent_Depths(t *testing.T) {
	agent := OKWSAgent{}
	config := GetDefaultConfig()

	// Step1: Start agent.
	agent.Start(config)

	// Step2: Subscribe channel
	// Step2.0: Subscribe public channel swap/depths successfully.
	agent.Subscribe(CHNL_SWAP_DEPTH, "BTC-USD-SWAP", DefaultDataCallBack)

	// Step3: Client receive depths from websocket server.
	// Step3.0: Receive partial depths
	// Step3.1: Receive update depths (It may take a very long time to see Update Event.)

	time.Sleep(60 * time.Second)

	// Step4. Stop all the go routine run in background.
	agent.Stop()
	time.Sleep(1 * time.Second)
}

func TestOKWSAgent_mergeDepths(t *testing.T) {
	oldDepths := [][4]interface{}{
		{"5088.59", "34000", 0, 1},
		{"7200", "1", 0, 1},
		{"7300", "1", 0, 1},
	}

	// Case1.
	newDepths1 := [][4]interface{}{
		{"5088.59", "32000", 0, 1},
	}
	expectedMerged1 := [][4]interface{}{
		{"5088.59", "32000", 0, 1},
		{"7200", "1", 0, 1},
		{"7300", "1", 0, 1},
	}

	m1, e1 := mergeDepths(oldDepths, newDepths1)
	assert.True(t, e1 == nil)
	assert.True(t, len(*m1) == len(expectedMerged1) && (*m1)[0][1] == expectedMerged1[0][1] && (*m1)[0][1] == "32000")

	// Case2.
	newDepths2 := [][4]interface{}{
		{"7200", "0", 0, 1},
	}
	expectedMerged2 := [][4]interface{}{
		{"5088.59", "34000", 0, 1},
		{"7300", "1", 0, 1},
	}
	m2, e2 := mergeDepths(oldDepths, newDepths2)
	assert.True(t, e2 == nil)
	assert.True(t, len(*m2) == len(expectedMerged2) && (*m2)[0][1] == expectedMerged2[0][1] && (*m2)[0][1] == "34000")

	// Case3.
	newDepths3 := [][4]interface{}{
		{"5000", "1", 0, 1},
		{"7400", "1", 0, 1},
	}
	expectedMerged3 := [][4]interface{}{
		{"5000", "1", 0, 1},
		{"5088.59", "34000", 0, 1},
		{"7200", "1", 0, 1},
		{"7300", "1", 0, 1},
		{"7400", "1", 0, 1},
	}
	m3, e3 := mergeDepths(oldDepths, newDepths3)
	assert.True(t, e3 == nil)
	assert.True(t, len(*m3) == len(expectedMerged3) && (*m3)[0][1] == expectedMerged3[0][1] && (*m3)[0][1] == "1")

}

func TestOKWSAgent_calCrc32(t *testing.T) {

	askDepths := [][4]interface{}{
		{"5088.59", "34000", 0, 1},
		{"7200", "1", 0, 1},