package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Step 1: Login to get a real token
	loginData := map[string]string{
		"username": "testuser",
		"password": "password123",
	}
	
	loginJSON, err := json.Marshal(loginData)
	if err != nil {
		fmt.Printf("Error marshaling login data: %v\n", err)
		return
	}
	
	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		fmt.Printf("Error making login request: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading login response: %v\n", err)
		return
	}
	
	var loginResponse map[string]interface{}
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		fmt.Printf("Error unmarshaling login response: %v\n", err)
		return
	}
	
	token, ok := loginResponse["token"].(string)
	if !ok {
		fmt.Printf("Error getting token from response\n")
		return
	}
	
	fmt.Printf("Got real token from auth service\n")
	
	// Step 2: Test the accounts service with this real token
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8081/accounts", nil)
	if err != nil {
		fmt.Printf("Error creating accounts request: %v\n", err)
		return
	}
	
	req.Header.Set("Authorization", "Bearer "+token)
	
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error making accounts request: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading accounts response: %v\n", err)
		return
	}
	
	fmt.Printf("Accounts service response: %s\n", string(body))
	fmt.Printf("Status code: %d\n", resp.StatusCode)
}