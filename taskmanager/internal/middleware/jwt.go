package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const defaultJWTSecret = "abracadabra" // override via JWT_SECRET env

func getJWTSecretValue() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return defaultJWTSecret
	}
	return secret
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
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "taskmanager",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetJWTSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
