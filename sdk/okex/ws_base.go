
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
