package user_api

import (
	"net/http"

	"crypto-user/api"
	"crypto-user/db"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

/**
启动作业
*/
func DeleteUserHandler(c *gin.Context) {
	var user_request DeleteUserRequest
	if err := c.ShouldBindJSON(&user_request); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "parms err", Payload: nil})
		return
	}

	var user User
	if err := db.FindOneById(db.DB, db.CollectionUser, user_request.UID, &user); err != nil {
		c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: "user not found", Payload: nil})
		return
	}

	user.Status = USER_STATUS_DELETED
	if err := db.U