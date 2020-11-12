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
	SecretKey  string `bson:"secret_ke