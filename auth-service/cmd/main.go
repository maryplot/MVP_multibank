package main

import (
    "database/sql"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    _ "github.com/lib/pq"
    "golang.org/x/crypto/bcrypt"

    "github.com/ErzhanBersagurov/MVP_multibank/auth-service/models"
    "github.com/ErzhanBersagurov/MVP_multibank/auth-service/storage"
)

func main() {
    // Подключение к PostgreSQL
    db, err := sql.Open("postgres", "host=postgres user=postgres password=postgres dbname=auth_db sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    storage := storage.NewStorage(db)
    r := gin.Default()

    // Регистрация
    r.POST("/register", func(c *gin.Context) {
        var user models.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := storage.CreateUser(&user); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
    })

    // Логин и выдача JWT
    r.POST("/login", func(c *gin.Context) {
        var loginReq models.LoginRequest
        if err := c.ShouldBindJSON(&loginReq); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        user, err := storage.FindUserByUsername(loginReq.Username)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        // Генерация JWT токена
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "user_id": user.ID,
            "exp":     time.Now().Add(time.Hour * 24).Unix(),
        })

        tokenString, _ := token.SignedString([]byte("your-secret-key"))
        c.JSON(http.StatusOK, gin.H{"token": tokenString})
    })

    r.Run(":8080")
}