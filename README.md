# MVP Multibank Application

This is a demo application for a multi-bank service.

## Running the project without Docker

### Prerequisites

*   Go (version 1.23 or newer)
*   Node.js (version 18 or newer)
*   PostgreSQL (version 14 or newer)
*   pyenv (optional, for Node.js version management)

### 1. Database Setup

1.  **Start PostgreSQL service:**
    ```bash
    brew services start postgresql@14
    ```

2.  **Create the database:**
    ```bash
    createdb auth_db
    ```

### 2. Backend Setup

The backend consists of three Go services. Run each in a separate terminal window.

**Important Notes:**
- The project uses local Go modules - no need to modify import paths
- Database connection is already configured for localhost

**For each service (`auth-service`, `accounts-service`, `transfer-service`):**

1.  **Open a new terminal**
2.  **Navigate to the service directory:**
    ```bash
    cd <service-name> 
    # e.g., cd auth-service
    ```
3.  **Tidy the modules:**
    ```bash
    go mod tidy
    ```
4.  **Run the service:**
    ```bash
    go run ./cmd/main.go
    ```

**Service ports:**
- `auth-service`: `http://localhost:8080`
- `accounts-service`: `http://localhost:8081`
- `transfer-service`: `http://localhost:8082`

### 3. Frontend Setup

1.  **Initialize pyenv (if needed):**
    ```bash
    eval "$(pyenv init - bash)"
    ```

2.  **Navigate to the frontend directory:**
    ```bash
    cd frontend
    ```

3.  **Install dependencies:**
    ```bash
    npm install
    ```

4.  **Start the development server:**
    ```bash
    npm run dev
    ```
The frontend will be available at `http://localhost:5173`.

### 4. Seeding the Database

1.  **Seed the test user:**
    ```bash
    psql -d auth_db -f init.sql
    ```

**Login credentials:**
- **Username:** `testuser`
- **Password:** `password123`

## Troubleshooting

### Port Conflicts
If a service fails to start due to port conflict:
```bash
# Find process using port
netstat -vanp tcp | grep 8082

# Kill the process
kill -9 <PID>
```

### Pyenv Issues
If you see `pyenv: shell integration not enabled`:
```bash
# Initialize pyenv
eval "$(pyenv init - bash)"

# Verify Node.js version
node --version
```

### Expired Tokens
If you see `JWT Parse error: token is expired`:
1. Log out and log back in to get a new token
2. Or restart the auth-service to reset token expiration

## New Features

### Transaction History
The application includes a transaction history feature that displays all transfers made between accounts, showing:
- Transaction ID
- Source account
- Destination account
- Transfer amount and currency
- Transaction status
- Timestamp