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
	// Create a 