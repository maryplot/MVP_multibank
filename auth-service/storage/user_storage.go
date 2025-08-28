package storage

import (
    "database/sql"

    
    "github.com/ErzhanBersagurov/MVP_multibank/auth-service/models"
    "golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
    db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
    return &UserStorage{db: db}
}

func (s *UserStorage) CreateUser(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    
    err = s.db.QueryRow(
        "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id", 
        user.Username, user.Email, string(hashedPassword),
    ).Scan(&user.ID)
    
    return err
}

func (s *UserStorage) FindUserByUsername(username string) (*models.User, error) {
    user := &models.User{}
    err := s.db.QueryRow(
        "SELECT id, username, email, password FROM users WHERE username = $1", 
        username,
    ).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
    
    return user, err
}