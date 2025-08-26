package storage

import (
    "database/sql"
    "log"
    
    "github.com/ErzhanBersagurov/MVP_multibank/auth-service/models"
    "golang.org/x/crypto/bcrypt"
)

type Storage struct {
    db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
    return &Storage{db: db}
}

func (s *Storage) CreateUser(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("BCrypt error: %v", err)
        return err
    }
    
    err = s.db.QueryRow(
        "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id", 
        user.Username, user.Email, string(hashedPassword),
    ).Scan(&user.ID)
    
    if err != nil {
        log.Printf("Database error: %v", err)
    } else {
        log.Printf("User created successfully: ID=%d, Username=%s", user.ID, user.Username)
    }
    
    return err
}

func (s *Storage) FindUserByUsername(username string) (*models.User, error) {
    user := &models.User{}
    err := s.db.QueryRow(
        "SELECT id, username, email, password FROM users WHERE username = $1", 
        username,
    ).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
    
    if err != nil {
        log.Printf("Find user error: %v", err)
    }
    
    return user, err
}