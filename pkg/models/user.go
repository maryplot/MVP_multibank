package models

import "time"

type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username" binding:"required"`
    Email     string    `json:"email" binding:"required"`
    Password  string    `json:"password" binding:"required"`
    CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}