package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

func main() {
	// Get a real token from auth service
	loginReq := LoginRequest{
		Username: "testuser",
		Password: "password123",
	}
	
	jsonData, err := json.Marshal(loginReq)
	if err != nil {
		log.Fatalf("Error marshaling login request: %v", err)
	}
	
	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error making login request: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Login request failed with status: %d", resp.StatusCode)
	}
	
	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		log.Fatalf("Error decoding login response: %v", err)
	}
	
	fmt.Printf("Got token: %s\n", loginResp.Token)
	
	// Validate the token
	jwtSecret := "simple-secret-12345"
	parsedToken, err := jwt.Parse(loginResp.Token, func(token *jwt.Token) (interface{}, error) {
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
	exp := int64(claims["exp"].(float64))
	
	fmt.Printf("Valid token, user ID: %d\n", userID)
	fmt.Printf("Token expires at: %s\n", time.Unix(exp, 0).Format("2006-01-02 15:04:05"))
}