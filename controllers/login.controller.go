package controllers

import (
	"github.com/gin-gonic/gin"
	"simple-kong-gateway/models"
	"simple-kong-gateway/services"
)

type loginController struct {
	tokenService services.IJwtService
}

func (l loginController) Login(ctx *gin.Context) string {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return ""
	}

	token := l.tokenService.GenerateToken(user.Username)
	return token
}

func NewLoginController(loginService services.IJwtService) ILoginController {
	return &loginController{tokenService: loginService}
}
