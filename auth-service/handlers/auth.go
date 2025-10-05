package handlers

import (
    "log"
    "net/http"
    "time"
    "os"
    
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    
    "auth-service/models"
    "auth-service/storage"
)

type AuthHandler struct {
    userStorage *storage.UserStorage
}

func NewAuthHandler(userStorage *storage.UserStorage) *AuthHandler {
    return &AuthHandler{userStorage: userStorage}
}

func (h *AuthHandler) Root(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "service": "auth-service",
        "version": "1.0",
        "endpoints": []string{
            "POST /register",
            "POST /login",
            "POST /validate",
            "GET /health",
        },
    })
}

func (h *AuthHandler) Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.userStorage.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "User created successfully", 
        "user_id": user.ID,
    })
}

func (h *AuthHandler) Login(c *gin.Context) {
    var creds models.LoginRequest
    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    log.Printf("Login attempt for user: %s", creds.Username)
    
    user, err := h.userStorage.FindUserByUsername(creds.Username)
    if err != nil {
        log.Printf("User lookup error: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
        log.Printf("Password mismatch for user: %s", creds.Username)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        jwtSecret = "simple-secret-12345"
    }
    
    log.Printf("🔐 JWT Secret used for token generation: %s", jwtSecret)
    
    tokenString, err := token.SignedString([]byte(jwtSecret))
    if err != nil {
        log.Printf("🔐 Token signing error: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
        return
    }

    log.Printf("Successful login for user: %s", creds.Username)
    c.JSON(http.StatusOK, gin.H{
        "token": tokenString,
        "user_id": user.ID,
    })
}

func (h *AuthHandler) ValidateToken(c *gin.Context) {
    // Реализация валидации токена
    c.JSON(http.StatusOK, gin.H{"valid": true})
}

func (h *AuthHandler) HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "OK", "service": "auth-service"})
}