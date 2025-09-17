package services

import (
    "fmt"
    "log"

    "github.com/ErzhanBersagurov/MVP_multibank/accounts-service/clients"
    "github.com/ErzhanBersagurov/MVP_multibank/accounts-service/models"
    "github.com/ErzhanBersagurov/MVP_multibank/accounts-service/storage"
)

type BankService struct {
    alfaClient    *clients.AlfaClient
    tinkoffClient *clients.TinkoffClient
    sberClient    *clients.SberClient
}

func NewBankService(tinkoffToken, sberToken string) *BankService {
    return &BankService{
        alfaClient:    clients.NewAlfaClient("alfa_demo_token"),
        tinkoffClient: clients.NewTinkoffClient(tinkoffToken),
        sberClient:    clients.NewSberClient(sberToken),
    }
}

// GetAllAccounts возвращает все счета из всех банков
func (s *BankService) GetAllAccounts(userID int) ([]models.Account, error) {
    log.Printf("Getting accounts for user %d from all banks", userID)
    
    var allAccounts []models.Account

    // Получаем счета из Альфа-банка (первым)
    alfaAccounts, err := s.alfaClient.GetAccounts(userID)
    if err != nil {
        log.Printf("Error getting Alfa accounts: %v", err)
    } else {
        allAccounts = append(allAccounts, alfaAccounts...)
    }

    // Получаем счета из Тинькофф
    tinkoffAccounts, err := s.tinkoffClient.GetAccounts(userID)
    if err != nil {
        log.Printf("Error getting Tinkoff accounts: %v", err)
    } else {
        allAccounts = append(allAccounts, tinkoffAccounts...)
    }

    // Получаем счета из Сбербанка
    sberAccounts, err := s.sberClient.GetAccounts(userID)
    if err != nil {
        log.Printf("Error getting Sber accounts: %v", err)
    } else {
        allAccounts = append(allAccounts, sberAccounts...)
    }

    // Применяем изменения балансов из transfer-service
    balanceStorage := storage.GetInstance()
    for i := range allAccounts {
        accountID := allAccounts[i].ID
        balanceChange := balanceStorage.GetBalanceChange(accountID)
        allAccounts[i].Balance += balanceChange
    }

    log.Printf("Retrieved %d accounts total", len(allAccounts))
    return allAccounts, nil
}

// GetTotalBalance возвращает общий баланс по всем банкам
func (s *BankService) GetTotalBalance(userID int) (float64, error) {
    accounts, err := s.GetAllAccounts(userID)
    if err != nil {
        return 0, err
    }

    total := 0.0
    for _, account := range accounts {
        total += account.Balance
    }

    return total, nil
}

// GetBankAccounts возвращает счета конкретного банка
func (s *BankService) GetBankAccounts(userID int, bankName string) ([]models.Account, error) {
    switch bankName {
    case "alfa", "Альфа-Банк":
        return s.alfaClient.GetAccounts(userID)
    case "tinkoff", "Тинькофф":
        return s.tinkoffClient.GetAccounts(userID)
    case "sber", "Сбербанк":
        return s.sberClient.GetAccounts(userID)
    default:
        return nil, fmt.Errorf("unknown bank: %s", bankName)
    }
}

// GetAccountDetail возвращает детальную информацию о счете
func (s *BankService) GetAccountDetail(accountID string) (*models.Account, error) {
    // Определяем банк по префиксу accountID
    switch {
    case len(accountID) >= 4 && accountID[:4] == "alfa":
        return s.alfaClient.GetAccountDetails(accountID)
    case len(accountID) >= 6 && accountID[:6] == "tinkoff":
        return s.tinkoffClient.GetAccountDetails(accountID)
    case len(accountID) >= 4 && accountID[:4] == "sber":
        return s.sberClient.GetAccountDetails(accountID)
    default:
        return nil, fmt.Errorf("unknown account type: %s", accountID)
    }
}