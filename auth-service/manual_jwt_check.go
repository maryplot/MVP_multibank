package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// This is how the auth service generates tokens
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTk1MjI2OTEsInVzZXJfaWQiOjl9.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	
	fmt.Printf("Testing token: %s\n", tokenString)
	
	// This is how the accounts service validates tokens
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "simple-secret-12345"
	}
	
	fmt.Printf("JWT Secret: %s\n", jwtSecret)
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	
	if err != nil {
		fmt.Printf("JWT Parse error: %v\n", err)
		return
	}
	
	if !token.Valid {
		fmt.Printf("Token is invalid\n")
		return
	}
	
	fmt.Printf("Token is valid\n")
}