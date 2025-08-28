package database

import (
    "database/sql"
    "log"
    "os"
    "time"
    
    _ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
    connStr := os.Getenv("DATABASE_URL")
    if connStr == "" {
        connStr = "host=postgres user=postgres password=postgres dbname=auth_db sslmode=disable"
    }
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // Ожидаем подключения к БД
    for i := 0; i < 5; i++ {
        if err = db.Ping(); err == nil {
            break
        }
        log.Printf("Waiting for database... Attempt %d", i+1)
        time.Sleep(2 * time.Second)
    }
    
    if err != nil {
        return nil, err
    }

    // Создаем таблицы
    if err := createTables(db); err != nil {
        return nil, err
    }

    return db, nil
}

func createTables(db *sql.DB) error {
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(50) UNIQUE NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    return err
}