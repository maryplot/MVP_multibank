package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// This is the token we're trying to validate
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTk1MjI1MjYsInVzZXJfaWQiOjl9.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

	fmt.Printf("🔐 Token string: %s\n", tokenString)

	// Get JWT secret exactly like the accounts service
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "simple-secret-12345"
	}

	fmt.Printf("🔐 JWT Secret from env: '%s'\n", os.Getenv("JWT_SECRET"))
	fmt.Printf("🔐 JWT Secret being used: %s\n", jwtSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		log.Fatal("🔐 JWT Parse error: ", err)
	}

	if !token.Valid {
		log.Fatal("🔐 Token is invalid")
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))
	fmt.Printf("🔐 Authenticated user ID: %d\n", userID)

	fmt.Println("Token validation successful!")
}