# Go Gin Starter

A production-ready REST API starter template built with **Go**, **Gin Web Framework**, and **GORM ORM**. This project provides a solid foundation for building scalable web applications with authentication, user management, and clean architecture principles.

## 🎯 About

This is a comprehensive starter template that demonstrates best practices for building REST APIs in Go. It includes:

- **Authentication & Authorization**: JWT-based authentication with secure password hashing
- **User Management**: Complete CRUD operations for user management
- **Database Integration**: MySQL database with GORM ORM
- **Clean Architecture**: Well-organized code structure following Go conventions
- **Middleware Support**: JWT validation and request processing
- **Environment Configuration**: Flexible configuration management with `.env` support
- **Hot Reload**: Development setup with Air for automatic code reloading

## 📋 Features

- ✅ User registration and login with JWT tokens
- ✅ Password hashing with bcrypt
- ✅ Protected routes with JWT middleware
- ✅ User CRUD operations (Create, Read, Update, Delete)
- ✅ Admin role support
- ✅ Email verification tracking
- ✅ Comprehensive error handling
- ✅ RESTful API design
- ✅ Database migrations ready
- ✅ Development and production ready

## 🏗️ Project Structure

```
go-gin-starter/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── app/
│   │   └── app.go                  # Application initialization and setup
│   ├── config/
│   │   └── config.go               # Configuration management (env variables)
│   ├── database/
│   │   └── database.go             # Database connection and initialization
│   ├── dto/
│   │   ├── auth_request.go         # Authentication request DTOs
│   │   ├── auth_response.go        # Authentication response DTOs
│   │   ├── user_request.go         # User request DTOs
│   │   └── user_response.go        # User response DTOs
│   ├── handler/
│   │   ├── auth_handler.go         # Authentication endpoints (login, register, me)
│   │   ├── user_handler.go         # User management endpoints (CRUD)
│   │   ├── health_handler.go       # Health check endpoint
│   │   └── favicon_handler.go      # Favicon handler
│   ├── middleware/
│   │   └── middleware.go           # JWT authentication middleware
│   ├── model/
│   │   └── user.go                 # User data model with GORM tags
│   ├── repository/
│   │   └── user_repository.go      # Data access layer for users
│   ├── router/
│   │   └── router.go               # Route definitions and dependency injection
│   ├── security/
│   │   ├── jwt.go                  # JWT token generation and validation
│   │   └── password.go             # Password hashing and verification
│   └── service/
│       ├── auth_service.go         # Authentication business logic
│       └── user_service.go         # User management business logic
├── .env.example                    # Environment variables template
├── .gitignore                      # Git ignore rules
├── .air.toml                       # Air hot reload configuration
├── go.mod                          # Go module definition
├── go.sum                          # Go module checksums
└── README.md                       # This file
```

## 🔄 Architecture Overview

The application follows a **layered architecture** pattern:

```
HTTP Request
    ↓
Router (Route Definitions)
    ↓
Handler (HTTP Request/Response)
    ↓
Service (Business Logic)
    ↓
Repository (Data Access)
    ↓
Database (MySQL)
```

### Layer Responsibilities

- **Handler**: Receives HTTP requests, validates input, calls services, returns responses
- **Service**: Contains business logic, orchestrates repositories, handles domain rules
- **Repository**: Manages database operations, abstracts data access
- **Model**: Defines data structures with GORM tags for database mapping
- **DTO**: Data Transfer Objects for request/response serialization
- **Middleware**: Cross-cutting concerns like authentication
- **Security**: Cryptographic operations (JWT, password hashing)

## 🚀 Getting Started

### Prerequisites

- Go 1.25.11 or higher
- MySQL 5.7 or higher
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/muhusni/go-gin-starter.git
   cd go-gin-starter
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment variables**
   ```bash
   cp .env.example .env
   ```
   
   Edit `.env` with your configuration:
   ```env
   PORT=8080
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=your_password
   DB_NAME=dbapp
   JWT_SECRET=your_secret_key_here
   JWT_ISSUER=go-gin-starter
   ```

4. **Create database**
   ```sql
   CREATE DATABASE dbapp;
   ```

5. **Run the application**
   ```bash
   go run ./cmd/api/main.go
   ```

   Or with hot reload (requires Air):
   ```bash
   air
   ```

The API will be available at `http://localhost:8080`

## 📡 API Endpoints

### Health Check
- `GET /ping` - Simple ping endpoint
- `GET /health` - Health check endpoint

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login and get JWT token
- `GET /api/v1/me` - Get current user info (requires auth)

### User Management (Protected Routes)
- `GET /api/v1/users` - List all users
- `POST /api/v1/users` - Create new user
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

## 🔐 Authentication

The API uses **JWT (JSON Web Tokens)** for authentication:

1. **Register**: Create a new user account
   ```bash
   curl -X POST http://localhost:8080/api/v1/auth/register \
     -H "Content-Type: application/json" \
     -d '{
       "name": "John Doe",
       "email": "john@example.com",
       "password": "password123"
     }'
   ```

2. **Login**: Get JWT token
   ```bash
   curl -X POST http://localhost:8080/api/v1/auth/login \
     -H "Content-Type: application/json" \
     -d '{
       "email": "john@example.com",
       "password": "password123"
     }'
   ```

3. **Use Token**: Include in Authorization header
   ```bash
   curl -X GET http://localhost:8080/api/v1/me \
     -H "Authorization: Bearer YOUR_JWT_TOKEN"
   ```

## 📦 Dependencies

- **[Gin Web Framework](https://github.com/gin-gonic/gin)** - HTTP web framework
- **[GORM](https://gorm.io/)** - Object-Relational Mapping
- **[MySQL Driver](https://github.com/go-sql-driver/mysql)** - MySQL database driver
- **[JWT-Go](https://github.com/golang-jwt/jwt)** - JWT token handling
- **[Crypto](https://golang.org/x/crypto)** - Password hashing (bcrypt)
- **[Godotenv](https://github.com/joho/godotenv)** - Environment variable loading

## 🛠️ Development

### Hot Reload with Air

The project includes Air configuration for automatic code reloading during development:

```bash
air
```

Configuration is in `.air.toml`

### Running Tests

```bash
go test ./...
```

### Building for Production

```bash
go build -o api ./cmd/api/main.go
```

## 📝 Key Components

### User Model
```go
type User struct {
    ID              uint64     // Primary key
    Name            string     // User's full name
    Email           string     // User's email (unique)
    EmailVerifiedAt *time.Time // Email verification timestamp
    Password        string     // Hashed password
    RememberToken   string     // Remember me token
    IsAdmin         bool       // Admin flag
    CreatedAt       time.Time  // Creation timestamp
    UpdatedAt       time.Time  // Last update timestamp
}
```

### JWT Configuration
- **Secret**: Used to sign and verify tokens
- **Issuer**: Token issuer identifier
- **Token Claims**: Includes user ID and email

### Password Security
- Passwords are hashed using bcrypt
- Never stored in plain text
- Verified during login

## 🔧 Configuration

All configuration is managed through environment variables in `.env`:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | 8080 |
| `DB_HOST` | Database host | localhost |
| `DB_PORT` | Database port | 3306 |
| `DB_USER` | Database user | root |
| `DB_PASSWORD` | Database password | - |
| `DB_NAME` | Database name | dbapp |
| `JWT_SECRET` | JWT signing secret | - |
| `JWT_ISSUER` | JWT issuer name | go-gin-starter |

## 📚 Best Practices Implemented

- ✅ **Separation of Concerns**: Clear layer separation (handler, service, repository)
- ✅ **Dependency Injection**: Dependencies passed through constructors
- ✅ **Error Handling**: Comprehensive error handling throughout
- ✅ **Security**: Password hashing, JWT validation, protected routes
- ✅ **Configuration Management**: Environment-based configuration
- ✅ **Code Organization**: Following Go project layout conventions
- ✅ **RESTful Design**: Proper HTTP methods and status codes
- ✅ **Data Transfer Objects**: DTOs for request/response handling

## 🚦 Status Codes

- `200 OK` - Successful request
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request data
- `401 Unauthorized` - Missing or invalid authentication
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## 📄 License

This project is open source and available under the MIT License.

## 👤 Author

Created by [muhusni](https://github.com/muhusni)

## 🤝 Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

## 📞 Support

For issues, questions, or suggestions, please open an issue on the GitHub repository.

---

**Happy coding! 🎉**
