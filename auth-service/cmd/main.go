package main

import (
	"log"
    "os"
    
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

    // Запуск сервера
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    log.Printf("Auth service starting on :%s", port)
    if err := authHandler.StartServer(port); err != nil {
        log.Fatal("Server failed:", err)
    }
}