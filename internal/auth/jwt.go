package auth

import (
	// "github.com/golang-jwt/jwt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	// todo use value from config
	expiration := time.Hour * 24

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration),
	})

	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
