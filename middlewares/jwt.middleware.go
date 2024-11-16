package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"simple-kong-gateway/services"
	"strings"
)

const TokenType = "Bearer"

func AuthorizeToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		tokenReq := strings.Trim(authHeader[len(TokenType):], " ")
		tokenResp, err := services.NewJwtService().ValidateToken(tokenReq)
		if tokenResp.Valid {
			ctx.Next()
		} else {
			log.Printf("Error validating token: %v", err)
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		}
	}
}
