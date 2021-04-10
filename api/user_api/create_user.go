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
	count, err := db.FindCount(db.DB, db.CollectionUser, bson.M{"username": user_request.Username})
	if err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "db err", Payload: nil})
		return
	}

	if count != 0 {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user already exist", Payload: nil})
		return
	}

	// add salt && sha256 password
	salt := utils.GenRandomStr(time.Now().UnixNano(), 64)
	h := sha256.New()
	h.Write([]byte(user_request.Password + salt))
	password := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("%s", password)
	user := User{
		UID:       utils.GenId(),
		Username:  user_request.Username,
		Password:  password,
		Salt:      salt,
		Status:    USE