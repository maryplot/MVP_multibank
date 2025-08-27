package clients

import (
    "github.com/ErzhanBersagurov/MVP_multibank/accounts-service/models"
)

type TinkoffClient struct {
    token string
}

func NewTinkoffClient(token string) *TinkoffClient {
    return &TinkoffClient{token: token}
}

func (c *TinkoffClient) GetAccounts(userID int) ([]models.Account, error) {
    // ЗАГЛУШКА для демо - возвращаем тестовые данные
    return []models.Account{
        {
            ID:            "tinkoff_123",
            UserID:        userID,
            BankName:      "Тинькофф",
            AccountNumber: "****1234",
            Balance:       150000.50,
            Currency:      "RUB",
            AccountType:   "debit",
        },
        {
            ID:            "tinkoff_456",
            UserID:        userID,
            BankName:      "Тинькофф",
            AccountNumber: "****5678",
            Balance:       89000.30,
            Currency:      "RUB",
            AccountType:   "investment",
        },
    }, nil
}

func (c *TinkoffClient) GetAccountBalance(accountID string) (float64, error) {
    // ЗАГЛУШКА для демо
    switch accountID {
    case "tinkoff_123":
        return 150000.50, nil
    case "tinkoff_456":
        return 89000.30, nil
    default:
        return 0, nil
    }
}

func (c *TinkoffClient) GetAccountDetails(accountID string) (*models.Account, error) {
    // ЗАГЛУШКА для демо
    return &models.Account{
        ID:            accountID,
        BankName:      "Тинькофф",
        AccountNumber: "****" + accountID[len(accountID)-4:],
        Balance:       150000.50,
        Currency:      "RUB",
        AccountType:   "debit",
    }, nil
}