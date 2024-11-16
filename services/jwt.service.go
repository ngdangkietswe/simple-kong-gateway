package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"simple-kong-gateway/configs"
	"time"
)

type jwtService struct {
}

type jwtClaim struct {
	name               string
	jwt.StandardClaims // This struct contains the standard JWT claims
}

// GenerateToken generates a JWT token
func (j jwtService) GenerateToken(username string) string {
	claims := &jwtClaim{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * configs.GlobalConfig.JwtExp).Unix(),
			Issuer:    configs.GlobalConfig.JwtIssuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSign, err := token.SignedString([]byte(configs.GlobalConfig.JwtSecret))
	if err != nil {
		panic(err)
	}

	return tokenSign
}

// ValidateToken validates a JWT token
func (j jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return []byte(configs.GlobalConfig.JwtSecret), nil
	})
}

func NewJwtService() IJwtService {
	return &jwtService{}
}
