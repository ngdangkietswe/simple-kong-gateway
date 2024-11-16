package models

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
}
