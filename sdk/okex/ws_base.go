
package okex

/*
 OKEX websocket api wrapper
 @author Lingting Fu
 @date 2018-12-27
 @version 1.0.0
*/

import (
	"bytes"
	"errors"
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"
)

type BaseOp struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

func subscribeOp(sts []*SubscriptionTopic) (op *BaseOp, err error) {

	strArgs := []string{}

	for i := 0; i < len(sts); i++ {
		channel, err := sts[i].ToString()
		if err != nil {
			return nil, err
		}
		strArgs = append(strArgs, channel)
	}

	b := BaseOp{
		Op:   "subscribe",
		Args: strArgs,
	}
	return &b, nil
}

func unsubscribeOp(sts []*SubscriptionTopic) (op *BaseOp, err error) {

	strArgs := []string{}

	for i := 0; i < len(sts); i++ {
		channel, err := sts[i].ToString()
		if err != nil {
			return nil, err
		}
		strArgs = append(strArgs, channel)
	}

	b := BaseOp{
		Op:   CHNL_EVENT_UNSUBSCRIBE,
		Args: strArgs,
	}
	return &b, nil
}

func loginOp(apiKey string, passphrase string, timestamp string, sign string) (op *BaseOp, err error) {
	b := BaseOp{
		Op:   "login",
		Args: []string{apiKey, passphrase, timestamp, sign},
	}
	return &b, nil
}

type SubscriptionTopic struct {
	channel string
	filter  string `default:""`
}

func (st *SubscriptionTopic) ToString() (topic string, err error) {
	if len(st.channel) == 0 {
		return "", ERR_WS_SUBSCRIOTION_PARAMS
	}

	if len(st.filter) > 0 {
		return st.channel + ":" + st.filter, nil
	} else {
		return st.channel, nil
	}
}

type WSEventResponse struct {
	Event   string `json:"event"`
	Success string `json:success`
	Channel string `json:"channel"`
}

func (r *WSEventResponse) Valid() bool {
	return len(r.Event) > 0 && len(r.Channel) > 0
}

type WSTableResponse struct {
	Table  string        `json:"table"`
	Action string        `json:"action",default:""`
	Data   []interface{} `json:"data"`
}

func (r *WSTableResponse) Valid() bool {
	return (len(r.Table) > 0 || len(r.Action) > 0) && len(r.Data) > 0
}

type WSDepthItem struct {
	InstrumentId string           `json:"instrument_id"`
	Asks         [][4]interface{} `json:"asks"`
	Bids         [][4]interface{} `json:"bids"`
	Timestamp    string           `json:"timestamp"`
	Checksum     int32            `json:"checksum"`
}

func mergeDepths(oldDepths [][4]interface{}, newDepths [][4]interface{}) (*[][4]interface{}, error) {

	mergedDepths := [][4]interface{}{}
	oldIdx, newIdx := 0, 0

	for oldIdx < len(oldDepths) && newIdx < len(newDepths) {

		oldItem := oldDepths[oldIdx]
		newItem := newDepths[newIdx]

		oldPrice, e1 := strconv.ParseFloat(oldItem[0].(string), 10)
		newPrice, e2 := strconv.ParseFloat(newItem[0].(string), 10)
		if e1 != nil || e2 != nil {
			return nil, fmt.Errorf("Bad price, check why. e1: %+v, e2: %+v", e1, e2)
		}

		if oldPrice == newPrice {
			newNum := StringToInt64(newItem[1].(string))

			if newNum > 0 {
				mergedDepths = append(mergedDepths, newItem)
			}

			oldIdx++
			newIdx++
		} else if oldPrice > newPrice {
			mergedDepths = append(mergedDepths, newItem)
			newIdx++
		} else if oldPrice < newPrice {
			mergedDepths = append(mergedDepths, oldItem)
			oldIdx++
		}
	}

	for ; oldIdx < len(oldDepths); oldIdx++ {
		mergedDepths = append(mergedDepths, oldDepths[oldIdx])
	}

	for ; newIdx < len(newDepths); newIdx++ {
		mergedDepths = append(mergedDepths, newDepths[newIdx])
	}

	return &mergedDepths, nil
}

func (di *WSDepthItem) update(newDI *WSDepthItem) error {
	newAskDepths, err1 := mergeDepths(di.Asks, newDI.Asks)
	if err1 != nil {
		return err1
	}

	newBidDepths, err2 := mergeDepths(di.Bids, newDI.Bids)
	if err2 != nil {
		return err2
	}

	crc32BaseBuffer, expectCrc32 := calCrc32(newAskDepths, newBidDepths)

	if expectCrc32 != newDI.Checksum {
		return fmt.Errorf("Checksum's not correct. LocalString: %s, LocalCrc32: %d, RemoteCrc32: %d",
			crc32BaseBuffer.String(), expectCrc32, newDI.Checksum)
	} else {
		di.Checksum = newDI.Checksum
		di.Bids = *newBidDepths
		di.Asks = *newAskDepths
		di.Timestamp = newDI.Timestamp
	}

	return nil
}