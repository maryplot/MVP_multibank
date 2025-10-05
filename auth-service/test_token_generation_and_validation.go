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

	fmt.Printf("Generated token: [REDACTED - token is long]\n")
	fmt.Printf("Token length: %d characters\n", len(tokenString))

	// Now validate the token immediately
	tokenParsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
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
	fmt.Printf("Token is valid for user ID: %d\n", userID)
}