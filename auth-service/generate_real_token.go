package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Generate a real token with valid signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 9,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "simple-secret-12345"
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		fmt.Printf("Error signing token: %v\n", err)
		return
	}

	fmt.Printf("Real token with valid signature: %s\n", tokenString)
}