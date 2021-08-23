package libs

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

type EmailConfirmationTokenClaims struct {
	Email string
	jwt.StandardClaims
}

func GenerateEmailConfirmationToken(email string) (string, error) {
	claims := &EmailConfirmationTokenClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := sign.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
