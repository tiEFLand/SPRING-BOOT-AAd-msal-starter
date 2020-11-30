package user_api

import (
	"crypto-user/db"
	"crypto-user/sdk/okex"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type OkexKeyDetail struct {
	APIKEY     string `bson:"api_key" json:"api_key"`
	SecretKey  string `bson:"secret_key" json:"secret_key"`
	PassPhrase string `bson:"passphrase" json:"passphrase"`
}
type HuobiKeyDetail struct {
	APIKEY    string `bson:"api_key" json:"api_key"`
	SecretKey string `bson:"secret_key" json:"secret_key"`
}

type User struct {
	UID       string         `bson:"_id" json:"uid"`
	Username  string         `bson:"username" json:"username"`
	Password  string         `bson:"password" json:"password"` // hex
	Salt      string         `bson:"salt" json:"salt"`
	OkexKey   OkexKeyDetail  `bson:"okex_key" json:"okex_key"`
	HuobiKey  HuobiKeyDetail `bson:"huobi_key" json:"huobi_key"`
	PushUID   int64          `bson:"push_uid" json:"push_uid"`
	Status    string         `bson:"status" json:"status"`
	CreatedTS time.Time      `bson:"created_ts" json:"created_ts"`
}

const (
	USER_STATUS_DELETED = "deleted"
	USER_STATUS_ACTIVE  = "active"
)

func FetchUserKeyByUID(uid string) (*User, error) {
	var user User
	getUserErr := db.FindOneById(db.DB, db.CollectionUser, uid, &user)
	if getUserErr == nil {
		r