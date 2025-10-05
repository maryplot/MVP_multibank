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

	fmt.Printf("Generated token length: %d\n", len(tokenString))

	// Now validate the token exactly like the accounts service does
	authHeader := "Bearer " + tokenString
	tokenStringFromHeader := authHeader[7:] // Remove "Bearer "

	jwtSecretForValidation := os.Getenv("JWT_SECRET")
	if jwtSecretForValidation == "" {
		jwtSecretForValidation = "simple-secret-12345"
	}

	tokenParsed, err := jwt.Parse(tokenStringFromHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretForValidation), nil
	})

	if err != nil {
		fmt.Printf("JWT Parse error: %v\n", err)
		return
	}

	if !tokenParsed.Valid {
		fmt.Printf("Token is invalid\n")
		return
	}

	claims := tokenParsed.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))
	fmt.Printf("Token validation successful for user ID: %d\n", userID)
}