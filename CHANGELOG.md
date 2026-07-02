# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-07-02

### Added
- **Authentication System**: JWT-based authentication with secure token generation and validation
  - User registration endpoint (`POST /api/v1/auth/register`)
  - User login endpoint (`POST /api/v1/auth/login`)
  - Current user info endpoint (`GET /api/v1/me`)
  
- **User Management**: Complete CRUD operations for user management
  - List all users (`GET /api/v1/users`)
  - Create new user (`POST /api/v1/users`)
  - Get user by ID (`GET /api/v1/users/:id`)
  - Update user (`PUT /api/v1/users/:id`)
  - Delete user (`DELETE /api/v1/users/:id`)

- **Security Features**:
  - Password hashing with bcrypt
  - JWT middleware for protected routes
  - Secure password verification
  - Admin role support

- **Database Integration**:
  - MySQL database support with GORM ORM
  - User model with comprehensive fields (ID, Name, Email, Password, IsAdmin, etc.)
  - Email verification tracking
  - Timestamps for audit trail (CreatedAt, UpdatedAt)

- **API Infrastructure**:
  - Health check endpoints (`GET /ping`, `GET /health`)
  - Favicon handler
  - RESTful API design with proper HTTP status codes
  - Comprehensive error handling

- **Project Structure**:
  - Clean layered architecture (Handler → Service → Repository → Database)
  - Dependency injection pattern
  - Data Transfer Objects (DTOs) for request/response handling
  - Middleware support for cross-cutting concerns

- **Development Tools**:
  - Air configuration for hot reload during development
  - Environment variable configuration with `.env` support
  - Go module management with go.mod and go.sum

- **Documentation**:
  - Comprehensive README with setup instructions
  - API endpoint documentation
  - Architecture overview
  - Best practices implementation guide

### Features
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

### Technical Stack
- **Language**: Go 1.25.11+
- **Web Framework**: Gin Web Framework
- **ORM**: GORM
- **Database**: MySQL 5.7+
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Environment Management**: godotenv

### Project Structure
```
go-gin-starter/
├── cmd/api/                    # Application entry point
├── internal/
│   ├── app/                    # Application initialization
│   ├── config/                 # Configuration management
│   ├── database/               # Database connection
│   ├── dto/                    # Data Transfer Objects
│   ├── handler/                # HTTP handlers
│   ├── middleware/             # Middleware (JWT auth)
│   ├── model/                  # Data models
│   ├── repository/             # Data access layer
│   ├── router/                 # Route definitions
│   ├── security/               # JWT and password utilities
│   └── service/                # Business logic
├── .env.example                # Environment template
├── .gitignore                  # Git ignore rules
├── .air.toml                   # Air hot reload config
├── go.mod                      # Go module definition
├── go.sum                      # Module checksums
└── README.md                   # Documentation
```

### Getting Started
1. Clone the repository
2. Install dependencies: `go mod download`
3. Configure environment: `cp .env.example .env`
4. Create database: `CREATE DATABASE dbapp;`
5. Run application: `go run ./cmd/api/main.go` or `air`

### API Endpoints
- **Health**: `GET /ping`, `GET /health`
- **Auth**: `POST /api/v1/auth/register`, `POST /api/v1/auth/login`, `GET /api/v1/me`
- **Users**: `GET /api/v1/users`, `POST /api/v1/users`, `GET /api/v1/users/:id`, `PUT /api/v1/users/:id`, `DELETE /api/v1/users/:id`

### Known Limitations
- Database migrations are manual (ready for implementation)
- Email verification is tracked but not enforced
- Rate limiting not implemented

### Future Enhancements
- Automated database migrations
- Email verification workflow
- Rate limiting and request throttling
- API documentation with Swagger/OpenAPI
- Unit and integration tests
- Docker containerization
- CI/CD pipeline integration
- Refresh token implementation
- Two-factor authentication
- Role-based access control (RBAC)

---

For more information, see [README.md](README.md)
