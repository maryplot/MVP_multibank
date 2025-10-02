# MVP Multibank Application

This is a demo application for a multi-bank service.

## Running the project without Docker

### Prerequisites

*   Go (version 1.23 or newer)
*   Node.js (version 18 or newer)
*   PostgreSQL (version 14 or newer)

### 1. Database Setup

First, you need to set up a PostgreSQL database.

1.  **Install PostgreSQL.**
    If you are on macOS and using Homebrew, you can run:
    ```bash
    brew install postgresql@14
    ```

2.  **Start the PostgreSQL service.**
    ```bash
    brew services start postgresql@14
    ```

3.  **Create the database.**
    ```bash
    createdb auth_db
    ```

### 2. Backend Setup

The backend consists of three Go services. You will need to run each in a separate terminal window.

**Important Note:** The project was originally configured to use remote Go modules. To run it locally, you need to modify the `go.mod` and import paths for each service.

For each service (`auth-service`, `accounts-service`, `transfer-service`):

1.  Open the `go.mod` file and change the module name from `github.com/ErzhanBersagurov/MVP_multibank/<service-name>` to just `<service-name>`.
2.  Update all internal import paths in the `.go` files from `github.com/ErzhanBersagurov/MVP_multibank/<service-name>/...` to `<service-name>/...`.
3.  In the `auth-service/database/database.go` file, change the `connStr` from `host=postgres` to `host=localhost`.

After making these changes, follow these steps for each service:

1.  **Open a new terminal.**
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

The services will be running on the following ports:
*   `auth-service`: `http://localhost:8080`
*   `accounts-service`: `http://localhost:8081`
*   `transfer-service`: `http://localhost:8082`

### 3. Frontend Setup

1.  **Open a new terminal.**
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

To log in, you need to create a test user.

1.  Make sure the `auth-service` is **not** running.
2.  From the root of the project, run the following command to execute the `init.sql` script:
    ```bash
    psql -d auth_db -f init.sql
    ```
3.  Restart the `auth-service`.

You can now log in with the following credentials:
*   **Username:** `testuser`
*   **Password:** `password123`