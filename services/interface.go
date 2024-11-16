package services

import "github.com/golang-jwt/jwt"

type IJwtService interface {
	GenerateToken(username string) string
	ValidateToken(token string) (*jwt.Token, error)
}
