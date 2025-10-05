package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "accounts-service/middleware"
    "accounts-service/services"
    "accounts-service/storage"
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

    // Health check (без аутентификации)
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "OK", "service": "accounts-service"})
    })

    // Добавляем JWT аутентификацию ко всем остальным эндпоинтам
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
        userID := c.GetInt("userID")
        authToken := c.GetHeader("Authorization")

        account, err := bankService.GetAccountDetail(accountID, userID, authToken)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, account)
    })

    // Эндпоинт для получения истории транзакций
    r.GET("/transfer/history", func(c *gin.Context) {
        userID := c.GetInt("userID")
        authToken := c.GetHeader("Authorization")
        
        transactions, err := bankService.GetTransactionHistory(userID, authToken)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        
        c.JSON(http.StatusOK, gin.H{"transactions": transactions})
    })

    // Эндпоинт для обновления балансов (используется transfer-service)
    r.POST("/balance/update", func(c *gin.Context) {
        var req struct {
            FromAccount string  `json:"from_account"`
            ToAccount   string  `json:"to_account"`
            Amount      float64 `json:"amount"`
        }
        
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        
        balanceStorage := storage.GetInstance()
        
        // Уменьшаем баланс счета-источника
        balanceStorage.UpdateBalance(req.FromAccount, -req.Amount)
        
        // Увеличиваем баланс счета-получателя
        balanceStorage.UpdateBalance(req.ToAccount, req.Amount)
        
        log.Printf("Updated balances: %s -%f, %s +%f", req.FromAccount, req.Amount, req.ToAccount, req.Amount)
        
        c.JSON(http.StatusOK, gin.H{"message": "Balances updated successfully"})
    })

    log.Println("Accounts service starting on :8081")
    r.Run(":8081")
}