package services

import (
    "fmt"       // ← ДОБАВИТЬ
    "sync"
    "time"      // ← ДОБАВИТЬ

    "github.com/ErzhanBersagurov/MVP_multibank/transfer-service/models"
)

type TransferService struct {
    transactions []models.Transaction
    mutex        sync.Mutex
}

func NewTransferService() *TransferService {
    return &TransferService{
        transactions: []models.Transaction{},
    }
}

func (s *TransferService) InternalTransfer(userID int, req models.TransferRequest) (*models.Transaction, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    transaction := models.Transaction{
        ID:          generateTransactionID(),
        FromAccount: req.FromAccount,
        ToAccount:   req.ToAccount,
        Amount:      req.Amount,
        Currency:    req.Currency,
        Status:      "completed",
        CreatedAt:   time.Now(),
    }

    s.transactions = append(s.transactions, transaction)
    return &transaction, nil
}

func (s *TransferService) GetTransactionHistory(userID int) []models.Transaction {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    return s.transactions
}

func generateTransactionID() string {
    return fmt.Sprintf("trx_%d", time.Now().Unix())
}