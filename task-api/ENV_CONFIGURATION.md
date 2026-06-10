# Environment Variables Configuration (.env)

This project uses environment variables to configure the application in a flexible way. Variables can be defined through a `.env` file located in the project's root directory.

## .env File

A sample `.env` file is included. During development, you can customize the values according to your needs.

### Available Variables

#### Database (PostgreSQL)

* `DB_HOST` - PostgreSQL server host (default: `db`)
* `DB_PORT` - PostgreSQL server port (default: `5432`)
* `DB_USER` - PostgreSQL username (default: `postgres`)
* `DB_PASSWORD` - PostgreSQL password (default: `postgres`)
* `DB_NAME` - Database name (default: `tasks`)
* `DB_SSLMODE` - SSL mode for the database connection (default: `disable`)

#### JWT Authentication

* `JWT_SECRET` - Secret key used to sign JWT tokens (default: `my-secret-key`)

  * **⚠️ Important:** Use a strong and secure secret in production environments.

#### Server

* `SERVER_PORT` - Port where the API will run (default: `8080`)

#### Authentication

* `ADMIN_USERNAME` - Administrator username used for login (default: `admin`)
* `ADMIN_PASSWORD` - Administrator password used for login (default: `admin`)

## Using Docker Compose

When using Docker Compose, the `.env` file is automatically loaded and its variables are passed to the API container.

```bash
docker compose up --build
```

## Running Locally (Without Docker)

The application automatically loads the `.env` file using the `godotenv` package. Make sure the `.env` file is located in the project's root directory (same level as `cmd/`, `internal/`, etc.).

```bash
go run ./cmd/server/main.go
```

## Example Custom Configuration

If you want to use different settings, edit the `.env` file:

```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=my_tasks_db
DB_SSLMODE=require

# JWT
JWT_SECRET=my-very-strong-secret-key

# Server
SERVER_PORT=3000

# Authentication
ADMIN_USERNAME=myuser
ADMIN_PASSWORD=mypassword
```

## Default Values

If a variable is not specified in the `.env` file, the application will use its default value. This makes the `.env` file optional for local development.

## Implementation Reference

Environment variables are loaded in `internal/config/config.go` through the `Load()` function. This package centralizes all application configuration.

### Example Usage

```go
package main

import "task-api/internal/config"

func main() {
    cfg := config.Load()

    // Access configuration values
    host := cfg.Database.Host
    port := cfg.Database.Port
    jwtSecret := cfg.JWT.Secret
    serverPort := cfg.Server.Port
}
```

## Security

⚠️ **Recommendations**

1. Never commit a `.env` file containing real credentials (the file is already included in `.gitignore`).
2. In production, always use strong and secure values for:

   * `JWT_SECRET`
   * `DB_PASSWORD`
   * `ADMIN_PASSWORD`
3. Consider using dedicated secret management solutions such as HashiCorp Vault, AWS Secrets Manager, or similar tools for production environments.
