package clients

import (
    "github.com/ErzhanBersagurov/MVP_multibank/accounts-service/models"
)

// SberClient представляет клиент для работы с Сбербанком (заглушка для демо)
type SberClient struct {
    token string
}

// NewSberClient создает новый экземпляр клиента Сбербанка
func NewSberClient(token string) *SberClient {
    return &SberClient{token: token}
}

// GetAccounts возвращает список счетов в Сбербанке (заглушка)
func (c *SberClient) GetAccounts(userID int) ([]models.Account, error) {
    // ЗАГЛУШКА для демо - возвращаем тестовые данные
    return []models.Account{
        {
            ID:            "sber_789",
            UserID:        userID,
            BankName:      "Сбербанк",
            AccountNumber: "****9012",
            Balance:       275000.75,
            Currency:      "RUB",
            AccountType:   "debit",
        },
        {
            ID:            "sber_345", 
            UserID:        userID,
            BankName:      "Сбербанк",
            AccountNumber: "****3456",
            Balance:       125000.25,
            Currency:      "RUB",
            AccountType:   "savings",
        },
        {
            ID:            "sber_678",
            UserID:        userID,
            BankName:      "Сбербанк",
            AccountNumber: "****7890",
            Balance:       50000.00,
            Currency:      "USD",
            AccountType:   "multi-currency",
        },
    }, nil
}

// GetAccountBalance возвращает баланс счета (заглушка)
func (c *SberClient) GetAccountBalance(accountID string) (float64, error) {
    // ЗАГЛУШКА для демо
    switch accountID {
    case "sber_789":
        return 275000.75, nil
    case "sber_345":
        return 125000.25, nil
    case "sber_678":
        return 50000.00, nil
    default:
        return 0, nil
    }
}

// GetAccountDetails возвращает детальную информацию о счете (заглушка)
func (c *SberClient) GetAccountDetails(accountID string) (*models.Account, error) {
    // ЗАГЛУШКА для демо
    switch accountID {
    case "sber_789":
        return &models.Account{
            ID:            "sber_789",
            BankName:      "Сбербанк",
            AccountNumber: "****9012",
            Balance:       275000.75,
            Currency:      "RUB",
            AccountType:   "debit",
        }, nil
    case "sber_345":
        return &models.Account{
            ID:            "sber_345",
            BankName:      "Сбербанк",
            AccountNumber: "****3456",
            Balance:       125000.25,
            Currency:      "RUB",
            AccountType:   "savings",
        }, nil
    case "sber_678":
        return &models.Account{
            ID:            "sber_678",
            BankName:      "Сбербанк",
            AccountNumber: "****7890",
            Balance:       50000.00,
            Currency:      "USD",
            AccountType:   "multi-currency",
        }, nil
    default:
        return &models.Account{
            ID:            accountID,
            BankName:      "Сбербанк",
            AccountNumber: "****" + accountID[len(accountID)-4:],
            Balance:       100000.00,
            Currency:      "RUB",
            AccountType:   "debit",
        }, nil
    }
}

// GetTotalBalance возвращает общий баланс по всем счетам в Сбербанке
func (c *SberClient) GetTotalBalance(userID int) (float64, error) {
    accounts, err := c.GetAccounts(userID)
    if err != nil {
        return 0, err
    }
    
    total := 0.0
    for _, account := range accounts {
        total += account.Balance
    }
    
    return total, nil
}