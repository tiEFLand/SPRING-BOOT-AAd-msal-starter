package user_api

import (
	"fmt"
	"net/http"
	"time"

	"crypto-user/api"
	"crypto-user/utils"

	"crypto-user/db"

	jwt_gin "github.com/appleboy/gin-jwt/v2"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

/**
启动作业
*/
func RefreshTokenHandler(c *gin.Context) {

	claims := jwt_gin.ExtractClaims(c)

	var user User
	if err := db.FindOne(db.DB, db.CollectionUser, bson.M{"_id": claims["uid"]}, nil, &user); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user not found", Payload: nil})
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
		fm