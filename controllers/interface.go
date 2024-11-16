package controllers

import "github.com/gin-gonic/gin"

type ILoginController interface {
	Login(ctx *gin.Context) string
}
