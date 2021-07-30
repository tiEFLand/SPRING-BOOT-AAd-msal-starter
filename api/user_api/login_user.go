
package user_api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"crypto-user/api"
	"crypto-user/utils"

	"crypto-user/db"

	"gopkg.in/mgo.v2/bson"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/**
启动作业
*/
func LoginUserHandler(c *gin.Context) {
	var user_request LoginUserRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	var user User
	if err := db.FindOne(db.DB, db.CollectionUser, bson.M{"username": user_request.Username}, nil, &user); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user not found", Payload: nil})
		return
	}

	// if password error
	if !checkPassword(user_request.Password, user.Password, user.Salt) {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user password incorrect", Payload: nil})
		return
	}

	secretKey, _ := utils.GetConfig().Get("jwt.secret")
	expireTime, _ := utils.GetConfig().GetInt("jwt.expire_time")

	// sign jwt and reply
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	uid := fmt.Sprintf("%d", user.PushUID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":           user.UID,
		"username":      user.Username,
		"okex_api_set":  maskRight(user.OkexKey.APIKEY, len(user.OkexKey.APIKEY)/4),
		"huobi_api_set": maskRight(user.HuobiKey.APIKEY, len(user.HuobiKey.APIKEY)/4),
		"push_uid_set":  maskRight(uid, len(uid)/4),
		"exp":           time.Now().Local().Add(time.Hour * time.Duration(expireTime)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(tokenString, err)
		c.JSON(http.StatusInternalServerError, api.JSONReply{ErrorCode: -1, ErrorDescription: "jwt sign err", Payload: nil})
		return
	}

	c.JSON(http.StatusOK, api.JSONReply{ErrorCode: 0, ErrorDescription: "success", Payload: struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}})

}

func genStar(len int) string {
	return strings.Repeat("•", len)
}

func maskRight(s string, length int) string {
	if length > 0 {
		rs := []rune(s)
		for i := len(rs) - 1; i >= length; i-- {
			rs[i] = '•'
		}
		return string(rs)
	}
	return ""
}