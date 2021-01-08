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
		return &user, nil
	}
	return nil, errors.New("user not found")
}

func (user *User) genOKConfig() okex.Config {
	apiKey := user.OkexKey.APIKEY
	secretKey := user.OkexKey.SecretKey
	passphrase := user.OkexKey.PassPhrase
	var config okex.Config
	config.Endpoint = "https://www.okex.com/"
	config.WSEndpoint = "wss://real.okex.com:10442/"
	config.ApiKey = apiKey
	config.SecretKey = secretKey
	config.Passphrase = passphrase
	config.TimeoutSecond = 45
	config.IsPrint = false
	config.I18n = okex.ENGLISH
	return config
}

func (user *User) GetOkexSpotCurrency(currency string) (float64, error) {
	client := okex.NewClient(user.genOKConfig())

	spotMap, err := client.GetSpotAccountsCurrency(currency)
	if err != nil {
		return 0, err
	}
	fmt.Println(spotMap)

	currencyFloat, err := strconv.ParseFloat((*spotMap)["available"].(string), 64)
	if err != nil {
		return 0, err
	}
	return currencyFloat, nil
}

/*  -------------------------------HUOBI-------------------------------------*/

func (user *User) GetHuobiSpotCurrency(ticker string) (float64, error) {
	// apiBuilder := builder.NewAPIBuilder().HttpTimeout(5 * time.Second)
	// access_key := user.HuobiKey.APIKEY
	// secret_key := user.HuobiKey.SecretKey

	// api := apiBuilder.APIKey(access_key).APISecretkey(secret_key).Build(goex.HUOBI_PRO)
	// account, err := api.GetAccount()
	// if err != nil {
	// 	return 0, nil
	// }
	return 0, nil
	//return account.SubAccounts()., nil
}
