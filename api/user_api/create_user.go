package user_api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"crypto-user/api"
	"crypto-user/db"
	"crypto-user/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

/**
创建用户
*/
func CreateUserHandler(c *gin.Context) {
	var user_request CreateUserRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	// username唯一
	count, err := db.FindCount(db.DB, db.CollectionUs