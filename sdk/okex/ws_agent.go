
package okex

/*
 OKEX websocket API agent
 @author Lingting Fu
 @date 2018-12-27
 @version 1.0.0
*/

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io/ioutil"

	"github.com/gorilla/websocket"

	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

var ReconnectCh = make(chan *WSSubscriptionCache, 999)

// save sub details into cache
// so when connection die, we are able to reconnect
type WSSubscriptionCache struct {
	Channel, Filter string
	Callback        ReceivedDataCallback
}

type OKWSAgent struct {
	baseUrl string
	config  *Config
	conn    *websocket.Conn

	wsEvtCh  chan interface{}
	wsErrCh  chan interface{}
	wsTbCh   chan interface{}
	stopCh   chan interface{}
	errCh    chan error
	signalCh chan os.Signal

	subMap         map[string][]ReceivedDataCallback
	activeChannels map[string]bool
	hotDepthsMap   map[string]*WSHotDepths

	processMut sync.Mutex
	cache      WSSubscriptionCache
}

func (a *OKWSAgent) Start(config *Config) error {
	a.baseUrl = config.WSEndpoint + "ws/v3?compress=true"
	for {
		// log.Printf("Connecting to %s", a.baseUrl)
		c, _, err := websocket.DefaultDialer.Dial(a.baseUrl, nil)
		if err != nil {
			log.Printf("dial:%+v", err)
			time.Sleep(time.Second * 1)
			continue
		} else {
			// log.Printf("Connected to %s", a.baseUrl)
			a.conn = c
			a.config = config

			a.wsEvtCh = make(chan interface{})
			a.wsErrCh = make(chan interface{})
			a.wsTbCh = make(chan interface{})
			a.errCh = make(chan error)
			a.stopCh = make(chan interface{}, 16)
			a.signalCh = make(chan os.Signal)
			a.activeChannels = make(map[string]bool)
			a.subMap = make(map[string][]ReceivedDataCallback)
			a.hotDepthsMap = make(map[string]*WSHotDepths)

			signal.Notify(a.signalCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

			go a.work()
			go a.receive()
			go a.finalize()
		}
		return nil