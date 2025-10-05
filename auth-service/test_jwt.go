package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Generate a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 9,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtSecret := "simple-secret-12345"
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated token: %s\n", tokenString)

	// Validate the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if !parsedToken.Valid {
		log.Fatal("Token is invalid")
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))
	fmt.Printf("Valid token for user ID: %d\n", userID)
}