
package okex

/*
 utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
 signing a message
 using: hmac sha256 + base64
  eg:
    message = Pre_hash function comment
    secretKey = E65791902180E9EF4510DB6A77F6EBAE

  return signed string = TO6uwdqz+31SIPkd4I+9NiZGmVH74dXi+Fd5X0EzzSQ=
*/
func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

/*
 the pre hash string
  eg:
    timestamp = 2018-03-08T10:59:25.789Z
    method  = POST
    request_path = /orders?before=2&limit=30
    body = {"product_id":"BTC-USD-0309","order_id":"377454671037440"}

  return pre hash string = 2018-03-08T10:59:25.789ZPOST/orders?before=2&limit=30{"product_id":"BTC-USD-0309","order_id":"377454671037440"}
*/
func PreHashString(timestamp string, method string, requestPath string, body string) string {
	return timestamp + strings.ToUpper(method) + requestPath + body
}

/*
  md5 sign
*/
func Md5Signer(message string) string {
	data := []byte(message)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

/*
 int convert string
*/
func Int2String(arg int) string {
	return strconv.Itoa(arg)
}

/*
 int64 convert string
*/
func Int642String(arg int64) string {
	return strconv.FormatInt(int64(arg), 10)
}

/*
  json string convert struct
*/
func JsonString2Struct(jsonString string, result interface{}) error {
	jsonBytes := []byte(jsonString)
	err := json.Unmarshal(jsonBytes, result)
	return err
}

/*
  json byte array convert struct
*/
func JsonBytes2Struct(jsonBytes []byte, result interface{}) error {
	err := json.Unmarshal(jsonBytes, result)
	return err
}

/*
 struct convert json string
*/
func Struct2JsonString(structt interface{}) (jsonString string, err error) {
	data, err := json.Marshal(structt)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
  ternary operator replace language: a == b ? c : d
*/
func T3O(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

/*
 Get a epoch time
  eg: 1521221737.376
*/
func EpochTime() string {
	millisecond := time.Now().UnixNano() / 1000000
	epoch := strconv.Itoa(int(millisecond))
	epochBytes := []byte(epoch)
	epoch = string(epochBytes[:10]) + "." + string(epochBytes[10:])
	return epoch
}

/*
 Get a iso time
  eg: 2018-03-16T18:02:48.284Z
*/
func IsoTime() string {
	utcTime := time.Now().UTC()
	iso := utcTime.String()
	isoBytes := []byte(iso)
	iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	return iso
}

/*
 Get utc +8 -- 1540365300000 -> 2018-10-24 15:15:00 +0800 CST
*/
func LongTimeToUTC8(longTime int64) time.Time {
	timeString := Int64ToString(longTime)
	sec := timeString[0:10]
	nsec := timeString[10:len(timeString)]
	return time.Unix(StringToInt64(sec), StringToInt64(nsec))
}

/*
 1540365300000 -> 2018-10-24 15:15:00
*/
func LongTimeToUTC8Format(longTime int64) string {
	return LongTimeToUTC8(longTime).Format("2006-01-02 15:04:05")
}

/*
  iso time change to time.Time
  eg: "2018-11-18T16:51:55.933Z" -> 2018-11-18 16:51:55.000000933 +0000 UTC
*/
func IsoToTime(iso string) (time.Time, error) {
	nilTime := time.Now()
	if iso == "" {
		return nilTime, errors.New("illegal parameter")