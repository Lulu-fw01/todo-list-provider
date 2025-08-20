package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todo-list-provider/configs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	authHeader   = "Authorization"
	userIDKey    = "sub"
	expiredAtKey = "exp"
)

type AuthServcie struct {
	cfg *configs.AuthConfig
}

func NewAuthService(cfg *configs.AuthConfig) AuthServcie {
	return AuthServcie{
		cfg: cfg,
	}
}

func (s *AuthServcie) CreateJWT(userId int) (string, error) {
	expiration := time.Second * time.Duration(s.cfg.JwtExperationSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIDKey:    strconv.Itoa(userId),
		expiredAtKey: time.Now().Add(expiration),
	})

	tokenStr, err := token.SignedString([]byte(s.cfg.JwtSecretPhrase))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (s *AuthServcie) JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader(authHeader)
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty auth token"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("")
			}
			return []byte(s.cfg.JwtSecretPhrase), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		payload := token.Claims.(jwt.MapClaims)
		userIDStr := payload[userIDKey].(string)
		userId, _ := strconv.Atoi(userIDStr)
		// todo check user exist

		c.Set(userIDKey, userId)
		c.Next()
	}
}

func GetUserIDFromCtx(c *gin.Context) int {
	val, exist := c.Get(userIDKey)
	if !exist {
		return -1
	}
	return val.(int)
}
