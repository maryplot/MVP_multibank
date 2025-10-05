#!/bin/bash

# Create logs directory if it doesn't exist
mkdir -p logs

# Start auth-service
echo "Starting auth-service..."
(cd auth-service && go run ./cmd/main.go > ../logs/auth.log 2>&1) &

# Start accounts-service
echo "Starting accounts-service..."
(cd accounts-service && go run ./cmd/main.go > ../logs/accounts.log 2>&1) &

# Start transfer-service
echo "Starting transfer-service..."
(cd transfer-service && go run ./cmd/main.go > ../logs/transfer.log 2>&1) &

# Start frontend
echo "Starting frontend..."
(cd frontend && npm run dev > ../logs/frontend.log 2>&1) &

echo "All services started. Check logs in the logs directory."