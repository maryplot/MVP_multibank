package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/ErzhanBersagurov/MVP_multibank/accounts-service/middleware"
    "github.com/ErzhanBersagurov/MVP_multibank/accounts-service/services"
)

func main() {
    // Инициализируем сервис с заглушками-токенами
    bankService := services.NewBankService("tinkoff_demo_token", "sber_demo_token")

    // Создаем Gin роутер
    r := gin.Default()

    // Middleware для логирования
    r.Use(func(c *gin.Context) {
        log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
        c.Next()
    })

    // Добавляем JWT аутентификацию ко всем эндпоинтам
    r.Use(middleware.JWTAuth())

    // Эндпоинт для получения всех счетов
    r.GET("/accounts", func(c *gin.Context) {
        userID := c.GetInt("userID") // Теперь из JWT токена!
        
        accounts, err := bankService.GetAllAccounts(userID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"accounts": accounts})
    })

    // Эндпоинт для получения счетов конкретного банка
    r.GET("/accounts/:bank", func(c *gin.Context) {
        userID := c.GetInt("userID") // Из JWT токена
        bankName := c.Param("bank")

        accounts, err := bankService.GetBankAccounts(userID, bankName)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"bank": bankName, "accounts": accounts})
    })

    // Эндпоинт для общего баланса
    r.GET("/balance", func(c *gin.Context) {
        userID := c.GetInt("userID") // Из JWT токена

        totalBalance, err := bankService.GetTotalBalance(userID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"total_balance": totalBalance})
    })

    // Эндпоинт для деталей счета
    r.GET("/accounts/detail/:id", func(c *gin.Context) {
        accountID := c.Param("id")

        account, err := bankService.GetAccountDetail(accountID)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, account)
    })

    // Health check (без аутентификации)
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "OK", "service": "accounts-service"})
    })

    log.Println("Accounts service starting on :8081")
    r.Run(":8081")
}