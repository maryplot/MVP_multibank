package main

import (
    "log"
    "net/http"
    

    "github.com/gin-gonic/gin"
    "github.com/ErzhanBersagurov/MVP_multibank/transfer-service/middleware"
    "github.com/ErzhanBersagurov/MVP_multibank/transfer-service/models"
    "github.com/ErzhanBersagurov/MVP_multibank/transfer-service/services"
)

func main() {
    transferService := services.NewTransferService()
    
    r := gin.Default()

    // JWT аутентификация
    r.Use(middleware.JWTAuth())

    // Логирование всех запросов
    r.Use(func(c *gin.Context) {
        log.Printf("📍 Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)
        c.Next()
    })

    // Перевод между своими счетами
    r.POST("/transfer/internal", func(c *gin.Context) {
        userID := c.GetInt("userID")
        log.Printf("🔄 Internal transfer request from user %d", userID)
        
        var req models.TransferRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        transaction, err := transferService.InternalTransfer(userID, req)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Transfer completed successfully",
            "transaction": transaction,
        })
    })

    // История транзакций
    r.GET("/transfer/history", func(c *gin.Context) {
        userID := c.GetInt("userID")
        transactions := transferService.GetTransactionHistory(userID)
        c.JSON(http.StatusOK, gin.H{"transactions": transactions})
    })

    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "OK", "service": "transfer-service"})
    })

    // Отладочный эндпоинт - список всех routes
    r.GET("/debug/routes", func(c *gin.Context) {
        routes := r.Routes()
        c.JSON(http.StatusOK, gin.H{"routes": routes})
    })

    log.Println("🚀 Transfer service starting on :8082")
    r.Run(":8082")
}