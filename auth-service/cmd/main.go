package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"auth-service/database"
	"auth-service/handlers"
	"auth-service/storage"
)

func main() {
    // Инициализация БД
    db, err := database.InitDB()
    if err != nil {
        log.Fatal("Database initialization failed:", err)
    }
    defer db.Close()

    // Инициализация хранилища
    userStorage := storage.NewUserStorage(db)
    
    // Инициализация обработчиков
    authHandler := handlers.NewAuthHandler(userStorage)

    // Настройка CORS
    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
        AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // Регистрация маршрутов
    // Группировка маршрутов под префиксом /api
    api := r.Group("/api")
    {
        api.GET("/", authHandler.Root)
        api.GET("/health", authHandler.HealthCheck)
        api.POST("/login", authHandler.Login)
        api.POST("/register", authHandler.Register)
        api.GET("/validate", authHandler.ValidateToken)
    }

    // Запуск сервера
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    log.Printf("Auth service starting on :%s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Server failed:", err)
    }
}