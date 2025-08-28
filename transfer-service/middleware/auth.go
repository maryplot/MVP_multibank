package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
    
)

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        // TODO: Реальная валидация через auth-service
        // Пока пропускаем все валидные JWT
        c.Set("userID", 1) // Заглушка
        c.Next()
    }
}