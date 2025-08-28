package models

import "time"

type Transaction struct {
    ID          string    `json:"id"`
    FromAccount string    `json:"from_account" binding:"required"`
    ToAccount   string    `json:"to_account" binding:"required"`
    Amount      float64   `json:"amount" binding:"required,gt=0"`
    Currency    string    `json:"currency" default:"RUB"`
    Status      string    `json:"status"` // pending, completed, failed
    CreatedAt   time.Time `json:"created_at"`
}

type TransferRequest struct {
    FromAccount string  `json:"from_account" binding:"required"`
    ToAccount   string  `json:"to_account" binding:"required"`
    Amount      float64 `json:"amount" binding:"required,gt=0"`
    Currency    string  `json:"currency" default:"RUB"`
}