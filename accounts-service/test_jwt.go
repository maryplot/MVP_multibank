package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Test JWT generation and validation
	jwtSecret := "simple-secret-12345"
	
	// Create a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 7,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatalf("Error signing token: %v", err)
	}
	
	fmt.Printf("Generated token: %s\n", tokenString)
	
	// Validate the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	
	if err != nil {
		log.Fatalf("Error parsing token: %v", err)
	}
	
	if !parsedToken.Valid {
		log.Fatalf("Token is invalid")
	}
	
	claims := parsedToken.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))
	fmt.Printf("Valid token, user ID: %d\n", userID)
}