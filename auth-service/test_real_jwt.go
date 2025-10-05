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

	// Get JWT secret exactly like the services
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "simple-secret-12345"
	}

	fmt.Printf("🔐 JWT Secret used: %s\n", jwtSecret)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatal("Token signing error: ", err)
	}

	fmt.Printf("Generated token: %s\n", tokenString)

	// Now validate it
	tokenParsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
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