package middleware

import (
    "net/http"
    "os"
    "strings"
    "log"  // ← ДОБАВИТЬ ИМПОРТ

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        log.Printf("🔐 Auth header: %s", authHeader)  // ← ДОБАВИТЬ ЛОГ
        
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        // Убираем "Bearer " из заголовка
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        log.Printf("🔐 Token string: %s", tokenString)  // ← ДОБАВИТЬ ЛОГ

        jwtSecret := os.Getenv("JWT_SECRET")
        if jwtSecret == "" {
             jwtSecret = "simple-secret-12345" 
        }
        log.Printf("🔐 JWT Secret: %s", jwtSecret)  // ← ДОБАВИТЬ ЛОГ

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(jwtSecret), nil
        })

        if err != nil {
            log.Printf("🔐 JWT Parse error: %v", err)  // ← ДОБАВИТЬ ЛОГ
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        if !token.Valid {
            log.Printf("🔐 Token is invalid")  // ← ДОБАВИТЬ ЛОГ
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims := token.Claims.(jwt.MapClaims)
        userID := int(claims["user_id"].(float64))
        log.Printf("🔐 Authenticated user ID: %d", userID)  // ← ДОБАВИТЬ ЛОГ

        c.Set("userID", userID)
        c.Next()
    }
}