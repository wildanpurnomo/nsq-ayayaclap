package libs

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type EmailConfirmationTokenClaims struct {
	Email string
	jwt.StandardClaims
}

type ClientAuthTokenClaims struct {
	UserID string
	jwt.StandardClaims
}

type ContextValueWrapper struct {
	GinContext *gin.Context
	UserID     string
}

const (
	AUTH_CONTEXT_KEY  = "auth"
	CONTEXT_VALUE_KEY = "context-value-wrapper"
)

var (
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

func (w *ContextValueWrapper) InvalidateJwtCookie() {
	w.GinContext.SetCookie("jwt", "", 1, "/", "", false, true)
}

func (w *ContextValueWrapper) SetJwtCookie(token string) {
	w.GinContext.SetCookie("jwt", token, 60*60*24, "/", "", false, true)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := ""
		
		cookie, _ := c.Request.Cookie("jwt")
		if cookie != nil {
			userId = ExtractUserIdFromAuthToken(cookie.Value)
		}
		
		c.Request = c.Request.WithContext(
			context.WithValue(
				c.Request.Context(),
				CONTEXT_VALUE_KEY,
				&ContextValueWrapper{
					GinContext: c,
					UserID:     userId,
				},
			),
		)

		c.Next()
	}
}

func ExtractContextValueWrapper(ctx context.Context) *ContextValueWrapper {
	return ctx.Value(CONTEXT_VALUE_KEY).(*ContextValueWrapper)
}

func GenerateClientAuthToken(userId string) string {
	claims := &ClientAuthTokenClaims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := sign.SignedString(JwtSecret)
	if err != nil {
		return ""
	}

	return token
}

func ExtractUserIdFromAuthToken(authToken string) string {
	claims := &ClientAuthTokenClaims{}

	token, err := jwt.ParseWithClaims(authToken, claims, func(t *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return ""
	}

	return claims.UserID
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
