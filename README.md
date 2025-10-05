# MVP Multibank Application

This is a demo application for a multi-bank service.

## macOS Setup

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

## Ubuntu 24.04 Setup

### 1. Install Dependencies

```bash
# Update package lists
sudo apt update

# Install Go
sudo apt install -y golang-go

# Install Node.js and npm
sudo apt install -y nodejs npm

# Install PostgreSQL
sudo apt install -y postgresql postgresql-contrib

# Install build essentials (for node-gyp)
sudo apt install -y build-essential
```

### 2. Database Setup

1.  **Start PostgreSQL service:**
    ```bash
    sudo service postgresql start
    ```

2.  **Create database user:**
    ```bash
    sudo -u postgres createuser --superuser $USER
    ```

3.  **Create the database:**
    ```bash
    createdb auth_db
    ```

4.  **Set password for your user (optional but recommended):**
    ```bash
    sudo -u postgres psql -c "ALTER USER $USER WITH PASSWORD 'your_password';"
    ```

### 3. Backend Setup

Same as macOS setup - run each service in a separate terminal.

### 4. Frontend Setup

Same as macOS setup.

### 5. Seeding the Database

Same as macOS setup.

## Troubleshooting

### Port Conflicts
If a service fails to start due to port conflict:
```bash
# Find process using port
sudo netstat -tulpn | grep :8082

# Kill the process
kill -9 <PID>
```

### PostgreSQL Authentication Issues
If you see "password authentication failed for user":

#### Option 1: Use Peer Authentication with Your System User (Recommended)
1. Update connection string in `auth-service/database/database.go` to use your system username:
   ```go
   connStr = "host=localhost user=your_username dbname=auth_db sslmode=disable"
   ```
   Replace `your_username` with your actual system username (e.g., 'maryplot')

2. Edit PostgreSQL config:
   ```bash
   sudo nano /etc/postgresql/14/main/pg_hba.conf
   ```

3. Add a line for your user at the top:
   ```conf
   local   auth_db       your_username                              peer
   ```
   Replace `your_username` with your actual system username

4. Restart PostgreSQL:
   ```bash
   sudo service postgresql restart
   ```

#### Option 2: Use Password Authentication
1. Set password for postgres user:
   ```bash
   sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD 'new_password';"
   ```

2. Update connection string in `auth-service/database/database.go`:
   ```go
   connStr = "host=localhost user=postgres password=new_password dbname=auth_db sslmode=disable"
   ```

#### Verify Connection
Test your connection:
```bash
psql -h localhost -U your_username -d auth_db
```

### Frontend Node.js Version Issues
If you see errors about Node.js version being too old for Vite:
```bash
# Install Node Version Manager (nvm)
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
source ~/.bashrc

# Install Node.js 20 (LTS)
nvm install 20
nvm use 20

# Verify Node.js version
node --version  # Should be 20.x or higher

# Reinstall dependencies
cd frontend
rm -rf node_modules package-lock.json
npm install

# Restart the development server
npm run dev
```

### Expired Tokens
If you see `JWT Parse error: token is expired`:
1. Log out and log back in to get a new token
2. Or restart the auth-service to reset token expiration

## Making the Service Publicly Accessible

To expose your service to the public internet:

### 1. Configure Services to Listen on All Interfaces
For each service (`auth-service`, `accounts-service`, `transfer-service`), update the server binding:

```go
// In each service's main.go file
r.Run("0.0.0.0:" + port)
```

### 2. Configure Frontend
Update the frontend's API base URL in `frontend/src/services/api.js`:
```js
const API_BASE_URL = "http://YOUR_SERVER_IP:8080";
```

### 3. Open Firewall Ports
Allow traffic to the necessary ports:
```bash
sudo ufw allow 8080/tcp  # auth-service
sudo ufw allow 8081/tcp  # accounts-service
sudo ufw allow 8082/tcp  # transfer-service
sudo ufw allow 5173/tcp  # frontend
```

### 4. Start Services with Public Binding
Run each service with the updated binding:
```bash
cd auth-service && go run ./cmd/main.go
cd accounts-service && go run ./cmd/main.go
cd transfer-service && go run ./cmd/main.go
cd frontend && npm run dev
```

### 5. Access the Application
Your application will be available at:
```
http://YOUR_SERVER_IP:5173
```

### Alternative: Production Deployment with Nginx
For better security and performance:
1. Install Nginx:
```bash
sudo apt install nginx
```

2. Create a config file `/etc/nginx/sites-available/multibank`:
```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:5173;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    location /api/auth {
        proxy_pass http://localhost:8080;
    }

    location /api/accounts {
        proxy_pass http://localhost:8081;
    }

    location /api/transfer {
        proxy_pass http://localhost:8082;
    }
}
```

3. Enable the site:
```bash
sudo ln -s /etc/nginx/sites-available/multibank /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
```

## New Features

### Transaction History
The application includes a transaction history feature that displays all transfers made between accounts, showing:
- Transaction ID
- Source account
- Destination account
- Transfer amount and currency
- Transaction status
- Timestamp