package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = []byte(os.Getenv("JWTSECRET"))

type Claims struct {
	UserID uint `json:"user_id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateToken (userID uint, name string) (string, error) {
	claims := Claims{
		UserID: userID,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
	
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		head := c.GetHeader("Authorization")
		if head == "" || (strings.Contains(head, "Bearer") == false) {
			c.JSON(401, gin.H{"error": "未携带token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(head, "Bearer ")
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "token无效或已过期"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("name", claims.Name)
		c.Next()
	}
}
