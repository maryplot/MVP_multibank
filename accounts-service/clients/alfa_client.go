package clients

import (
    "github.com/ErzhanBersagurov/MVP_multibank/accounts-service/models"
)

// AlfaClient представляет клиент для работы с Альфа-банком (заглушка для демо)
type AlfaClient struct {
    token string
}

// NewAlfaClient создает новый экземпляр клиента Альфа-банка
func NewAlfaClient(token string) *AlfaClient {
    return &AlfaClient{token: token}
}

// GetAccounts возвращает список счетов в Альфа-банке (заглушка)
func (c *AlfaClient) GetAccounts(userID int) ([]models.Account, error) {
    // ЗАГЛУШКА для демо - возвращаем тестовые данные
    return []models.Account{
        {
            ID:            "alfa_111",
            UserID:        userID,
            BankName:      "Альфа-Банк",
            AccountNumber: "****1111",
            Balance:       320000.00,
            Currency:      "RUB",
            AccountType:   "зарплатная карта",
        },
    }, nil
}

// GetAccountBalance возвращает баланс счета (заглушка)
func (c *AlfaClient) GetAccountBalance(accountID string) (float64, error) {
    // ЗАГЛУШКА для демо
    switch accountID {
    case "alfa_111":
        return 320000.00, nil
    case "alfa_222":
        return 180000.50, nil
    default:
        return 0, nil
    }
}

// GetAccountDetails возвращает детальную информацию о счете (заглушка)
func (c *AlfaClient) GetAccountDetails(accountID string) (*models.Account, error) {
    // ЗАГЛУШКА для демо
    switch accountID {
    case "alfa_111":
        return &models.Account{
            ID:            "alfa_111",
            BankName:      "Альфа-Банк",
            AccountNumber: "****1111",
            Balance:       320000.00,
            Currency:      "RUB",
            AccountType:   "debit",
        }, nil
    case "alfa_222":
        return &models.Account{
            ID:            "alfa_222",
            BankName:      "Альфа-Банк",
            AccountNumber: "****2222",
            Balance:       180000.50,
            Currency:      "RUB",
            AccountType:   "credit",
        }, nil
    default:
        return &models.Account{
            ID:            accountID,
            BankName:      "Альфа-Банк",
            AccountNumber: "****" + accountID[len(accountID)-4:],
            Balance:       200000.00,
            Currency:      "RUB",
            AccountType:   "debit",
        }, nil
    }
}

// GetTotalBalance возвращает общий баланс по всем счетам в Альфа-банке
func (c *AlfaClient) GetTotalBalance(userID int) (float64, error) {
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