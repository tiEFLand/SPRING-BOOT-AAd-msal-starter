package middleware

import (
	"crypto-user/api"
	"crypto-user/utils"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Token struct {
	UID      string
	Username string
	Expire   float64
}

func JwtMiddleware() *jwt.GinJWTMiddleware {
	secretKey, _ := utils.GetConfig().Get("jwt.secret")

	// the jwt middleware
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "user",
		Key:        []byte(secretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		IdentityHandler: func(c *gin.Context) interface{}