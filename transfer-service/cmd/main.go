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

    // Перевод между своими счетами
    r.POST("/transfer/internal", func(c *gin.Context) {
        userID := c.GetInt("userID")
        
        var req models.TransferRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // TODO: Проверка что оба счета принадлежат пользователю

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

    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "OK", "service": "transfer-service"})
    })

    log.Println("Transfer service starting on :8082")
    r.Run(":8082")
}