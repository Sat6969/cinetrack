package middlewares

import (
    "github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(401, gin.H{"error": "unauthorized"})
        c.Abort()
        return
    }
    c.Next()
}