package middlewares

import (
	"github.com/gin-gonic/gin"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(401, gin.H{"error": "invalid authorization header"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
	)

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "invalid or expired token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		c.JSON(401, gin.H{"error": "invalid token claims"})
		c.Abort()
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)

	if !ok {
		c.JSON(401, gin.H{"error": "user id not found in token"})
		c.Abort()
		return
	}

	userID := uint(userIDFloat)

	c.Set("user_id", userID)

	c.Next()
}
