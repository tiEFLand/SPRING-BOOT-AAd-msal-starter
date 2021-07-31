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