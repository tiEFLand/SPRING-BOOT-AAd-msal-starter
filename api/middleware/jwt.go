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

func JwtMiddleware() *jwt.GinJWTMiddle