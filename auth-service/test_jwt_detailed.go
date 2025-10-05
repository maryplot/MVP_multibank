package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Generate a token exactly like the auth service
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 9,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// Get JWT secret exactly like the auth service
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "simple-secret-12345"
	}

	fmt.Printf("🔐 JWT Secret used for token generation: %s\n", jwtSecret)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatal("Token signing error: ", err)
	}

	fmt.Printf("Generated token: %s\n", tokenString)

	// Validate the token exactly like the accounts service
	authHeader := "Bearer " + tokenString
	tokenStringFromHeader := authHeader[7:] // Remove "Bearer "

	fmt.Printf("🔐 Token string from header: %s\n", tokenStringFromHeader)

	// Get JWT secret exactly like the accounts service
	jwtSecretForValidation := os.Getenv("JWT_SECRET")
	if jwtSecretForValidation == "" {
		jwtSecretForValidation = "simple-secret-12345"
	}

	fmt.Printf("🔐 JWT Secret from env for validation: '%s'\n", os.Getenv("JWT_SECRET"))
	fmt.Printf("🔐 JWT Secret being used for validation: %s\n", jwtSecretForValidation)

	tokenParsed, err := jwt.Parse(tokenStringFromHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretForValidation), nil
	})

	if err != nil {
		log.Fatal("JWT Parse error: ", err)
	}

	if !tokenParsed.Valid {
		log.Fatal("Token is invalid")
	}

	claims := tokenParsed.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))
	fmt.Printf("🔐 Authenticated user ID: %d\n", userID)

	fmt.Println("Token validation successful!")
}