package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"simple-kong-gateway/configs"
	"simple-kong-gateway/controllers"
	"simple-kong-gateway/middlewares"
	"simple-kong-gateway/services"
)

var (
	jwtService      = services.NewJwtService()
	loginController = controllers.NewLoginController(jwtService)
)

func main() {
	router := gin.New()

	// Init proxy server
	host := fmt.Sprintf("%s:%d", configs.GlobalConfig.Host, configs.GlobalConfig.Port)
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   host,
	})

	// Routes setup
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Welcome to Simple Kong Gateway"})
	})

	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		}
	})

	router.Any("/api/*path", middlewares.AuthorizeToken(), func(ctx *gin.Context) {
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})

	err := router.Run(fmt.Sprintf(":%d", 8005))
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
		return
	}
}
