package services

import (
    "log"
    "time"  // ← ДОБАВИТЬ ИМПОРТ time

    "github.com/ErzhanBersagurov/MVP_multibank/transfer-service/models"
)

type TransferService struct {
    // В будущем добавим подключение к БД и банковским API
}

func NewTransferService() *TransferService {
    return &TransferService{}
}

func (s *TransferService) InternalTransfer(userID int, req models.TransferRequest) (*models.Transaction, error) {
    log.Printf("Processing internal transfer for user %d: %s -> %s (%.2f %s)",
        userID, req.FromAccount, req.ToAccount, req.Amount, req.Currency)

    // TODO: Реальная проверка принадлежности счетов пользователю
    // TODO: Проверка достаточности средств
    // TODO: Интеграция с банковскими API

    // Заглушка для демо
    transaction := &models.Transaction{
        ID:          "trx_123456",
        FromAccount: req.FromAccount,
        ToAccount:   req.ToAccount,
        Amount:      req.Amount,
        Currency:    req.Currency,
        Status:      "completed",
        CreatedAt:   time.Now(),  // ← Теперь time определен
    }

    return transaction, nil
}

func (s *TransferService) ValidateAccounts(userID int, accountIDs []string) error {
    // TODO: Реальная проверка что счета принадлежат пользователю
    // Пока заглушка
    return nil
}