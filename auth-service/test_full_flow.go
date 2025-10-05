package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Step 1: Generate a real token
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
		log.Fatal("Token signing error: ", err)
	}

	fmt.Printf("Generated real token: %s\n", tokenString)

	// Step 2: Test the accounts service with this real token
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/accounts", nil)
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+tokenString)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response: ", err)
	}

	fmt.Printf("Accounts service response: %s\n", string(body))
	fmt.Printf("Status code: %d\n", resp.StatusCode)
}