package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const defaultJWTSecret = "abracadabra"

func getJWTSecretValue() string {
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		return secret
	}
	return defaultJWTSecret
}

func GetJWTSecret() []byte {
	return []byte(getJWTSecretValue())
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, email string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "taskmanager",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetJWTSecret())
}
