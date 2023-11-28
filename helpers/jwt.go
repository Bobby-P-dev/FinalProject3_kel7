package helpers

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
	Secret := os.Getenv("JWT")
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(Secret))
	return signedToken
}

func VerifToken(c *gin.Context) (interface{}, error) {
	errRes := errors.New("sign in to proces")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	key := os.Getenv("KEY")

	if !bearer {
		return nil, errRes
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errRes
		}
		return []byte(key), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); !ok && token.Valid {
		return nil, errRes
	}
	return token.Claims.(jwt.MapClaims), nil
}
