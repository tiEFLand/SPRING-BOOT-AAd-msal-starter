package user_api

import (
	"net/http"

	"crypto-user/api"
	"crypto-user/db"

	"gopkg.in/mgo.v2/bson"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

/**
设置APIkey
*/
func SetKeyUserHandler(c *gin.Context) {
	var user_request SetKeyUserRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	claims := jwt.ExtractClaims(c)

	var user User
	if err := db.FindOne(db.DB, db.CollectionUser, bson.M{"username": claims["username"].(string)}, nil, &user); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user not found", Payload: nil})
		return
	}

	if user_request.OkexKey.APIKEY != "" {
		user.OkexKey.APIKEY = user_request.OkexKey.APIKEY
	}
	if user_request.OkexKey.SecretKey != "" {
		user.OkexKey.SecretKey = user_request.OkexKey.SecretKey
	}
	if user_request.OkexKey.PassPhrase != "" {
		user.OkexKey.PassPhrase = user_request.OkexKey.PassPhrase
	}
	if user_request.HuobiKey.APIKEY != "" {
		user.HuobiKey.APIKEY = user_request.HuobiKey.APIKEY
	}
	if user_request.HuobiKey.SecretKey != "" {
	