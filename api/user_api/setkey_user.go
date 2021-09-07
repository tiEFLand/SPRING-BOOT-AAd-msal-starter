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

	clai