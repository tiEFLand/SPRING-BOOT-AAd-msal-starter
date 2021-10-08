
package main

import (
	"runtime"

	"crypto-user/api/account_api"
	"crypto-user/api/middleware"
	"crypto-user/api/user_api"
	"crypto-user/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"x-url-path", "content-type", "Authorization"}
	config.AllowMethods = []string{"POST", "OPTIONS", "GET"}
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	userGroup := r.Group("/api/user")
	userGroup.POST("/create", user_api.CreateUserHandler)
	userGroup.POST("/delete", user_api.DeleteUserHandler)
	userGroup.POST("/login", user_api.LoginUserHandler)
	userGroup.POST("/wallet", account_api.GetWalletHandler)
	userGroup.Use(middleware.JwtMiddleware().MiddlewareFunc())
	{
		userGroup.POST("/setkey", user_api.SetKeyUserHandler)
	}
	port, _ := utils.GetConfig().Get("api.port")
	mode, _ := utils.GetConfig().Get("gin.mode")
	if mode == "" {
		mode = "debug"
	}
	if port == "" {
		port = "5001"
	}
	gin.SetMode(mode)
	r.Run(":" + port)
}