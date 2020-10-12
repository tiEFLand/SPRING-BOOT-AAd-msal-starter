
package account_api

import (
	"errors"
	"net/http"

	"crypto-user/api"
	"crypto-user/api/user_api"

	"github.com/gin-gonic/gin"
)

/**
获取Wallet
*/
func GetWalletHandler(c *gin.Context) {
	var user_request GetAccountRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	var user user_api.User
	// if err := db.FindOneById(db.DB, db.CollectionUser, claims["uid"], &user); err == nil {
	// 	//c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user not found", Payload: nil})
	// 	return
	// }

	user.OkexKey = user_api.OkexKeyDetail{
		APIKEY:     "7b805931-4721-4d13-8e6c-941d412c7729",
		SecretKey:  "4E736995A41CEB12F8C364C01F981F6D",
		PassPhrase: "cbt2018",
	}

	switch user_request.TickerType {
	case "spot":
		var available float64
		var err error
		if user_request.Ex == "OKEX" {
			available, err = user.GetOkexSpotCurrency(user_request.Ticker)
		} else if user_request.Ex == "HUOBI" {
			available, err = user.GetHuobiSpotCurrency(user_request.Ticker)
		} else {
			err = errors.New("unknow ex")
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: 0, ErrorDescription: "failed getting spot", Payload: err.Error()})
			return
		}
		c.JSON(http.StatusOK, api.JSONReply{ErrorCode: 0, ErrorDescription: "success", Payload: struct {
			Available float64 `json:"available"`
		}{
			Available: available,
		}})
	}
}