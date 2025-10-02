package models

import "time"

type Account struct {
    ID                 string        `json:"id"`
    UserID             int           `json:"user_id"`
    BankName           string        `json:"bank_name"`
    AccountNumber      string        `json:"account_number"`
    Balance            float64       `json:"balance"`
    Currency           string        `json:"currency"`
    AccountType        string        `json:"account_type"`  // debit, credit, savings
    CreatedAt          time.Time     `json:"created_at"`
    UpdatedAt          time.Time     `json:"updated_at"`
    TransactionHistory []Transaction `json:"transaction_history,omitempty"`
}

type Transaction struct {
    ID          string    `json:"id"`
    UserID      int       `json:"user_id"`
    FromAccount string    `json:"from_account"`
    ToAccount   string    `json:"to_account"`
    Amount      float64   `json:"amount"`
    Currency    string    `json:"currency"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
}

type BankCredentials struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    BankName  string    `json:"bank_name"`
    APIKey    string    `json:"api_key"`
    Token     string    `json:"token"`
    IsActive  bool      `json:"is_active"`
    CreatedAt time.Time `json:"created_at"`
}

// Ответ от Тинькофф API
type TinkoffAccountResponse struct {
    Accounts []struct {
        ID   string `json:"id"`
        Name string `json:"name"`
        Type string `json:"type"`
    } `json:"accounts"`
}

type TinkoffBalanceResponse struct {
    Balance struct {
        Amount   float64 `json:"amount"`
        Currency string  `json:"currency"`
    } `json:"balance"`
}