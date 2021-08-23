package libs

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

var (
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

type EmailConfirmationTokenClaims struct {
	Email string
	jwt.StandardClaims
}

func ExtractEmailFromRedirToken(redirToken string) (string, error) {
	var claims EmailConfirmationTokenClaims
	token, err := jwt.ParseWithClaims(redirToken, &claims, func(t *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("Invalid token")
	}

	return claims.Email, nil
}
