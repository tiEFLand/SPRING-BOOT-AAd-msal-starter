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
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &Token{
				UID:      claims["uid"].(string),
				Username: claims["username"].(string),
				Expire:   claims["exp"].(float64),
			}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusBadRequest, api.JSONReply{ErrorCode: -1, ErrorDescription: message, Payload: nil})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization",
		//