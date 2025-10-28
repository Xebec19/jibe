# Jibe API

A production-ready Go web server built with enterprise-grade design patterns and best practices, including clean architecture, dependency injection, resilience patterns, structured logging, and comprehensive middleware.

## Features

### Architecture & Design Patterns
- **Clean Architecture**: Clear separation of concerns with layers (domain, service, repository, handler)
- **Repository Pattern**: Abstraction layer for data access with interface-based design
- **Service Layer Pattern**: Business logic encapsulation separate from HTTP handlers
- **Dependency Injection**: Container-based DI for loose coupling and testability
- **Interface Segregation**: All dependencies rely on interfaces, not concrete implementations

### Resilience Patterns
- **Circuit Breaker**: Prevents cascading failures by stopping requests to failing services
- **Retry with Exponential Backoff**: Automatic retry logic with configurable backoff strategies
- **Timeout Handling**: Request-level timeouts to prevent resource exhaustion
- **Graceful Shutdown**: Proper cleanup and completion of in-flight requests

### Production Features
- **Structured Logging**: Using zerolog for efficient JSON logging with context
- **Environment Configuration**: Type-safe configuration with validation
- **Middleware Stack**:
  - Request ID tracking for distributed tracing
  - Request/response logging with metrics
  - Panic recovery with stack traces
  - CORS support
  - Security headers
- **RESTful API**: Gorilla Mux router with versioned endpoints
- **Health Checks**: Built-in health check endpoint for monitoring

## Project Structure

The project follows Clean Architecture principles with clear separation between layers:

```
api/
├── main.go                         # Application entry point & dependency wiring
├── go.mod                          # Go module definition
├── go.sum                          # Go module checksums
├── .env.example                    # Example environment variables
├── .gitignore                      # Git ignore rules
├── Makefile                        # Build and development commands
│
├── internal/                       # Private application code
│   ├── domain/                     # Domain layer (entities & business rules)
│   │   ├── user.go                # User domain model
│   │   └── errors.go              # Domain-specific errors
│   │
│   ├── repository/                 # Repository interfaces & implementations
│   │   ├── user_repository.go     # User repository interface
│   │   ├── memory/                # In-memory implementation (dev/testing)
│   │   │   └── user_repository.go
│   │   └── postgres/              # PostgreSQL implementation (production)
│   │       ├── user_repository.go
│   │       └── migrations.sql
│   │
│   ├── service/                    # Service layer (business logic)
│   │   └── user_service.go        # User service with business logic
│   │
│   ├── handlers/                   # HTTP handlers (presentation layer)
│   │   ├── handlers.go            # Common handlers (health, root)
│   │   └── user_handler.go        # User-specific handlers
│   │
│   ├── middleware/                 # HTTP middleware
│   │   └── middleware.go          # Request ID, logging, CORS, etc.
│   │
│   ├── container/                  # Dependency injection container
│   │   └── container.go           # Wires all dependencies together
│   │
│   └── config/                     # Configuration management
│       └── config.go              # Environment-based configuration
│
└── pkg/                            # Public/reusable code
    ├── logger/                     # Structured logging
    │   └── logger.go
    └── resilience/                 # Resilience patterns
        ├── circuit_breaker.go     # Circuit breaker implementation
        └── retry.go               # Retry with exponential backoff
```

### Architecture Layers

1. **Domain Layer** (`internal/domain/`)
   - Core business entities and rules
   - No dependencies on other layers
   - Defines domain errors and validation logic

2. **Repository Layer** (`internal/repository/`)
   - Data access abstraction
   - Interfaces define contracts, implementations handle storage
   - Supports multiple backends (memory, PostgreSQL)

3. **Service Layer** (`internal/service/`)
   - Business logic and orchestration
   - Coordinates between repositories
   - Enforces business rules and validation

4. **Handler Layer** (`internal/handlers/`)
   - HTTP request/response handling
   - Input validation and serialization
   - Delegates business logic to services

5. **Infrastructure** (`pkg/`)
   - Cross-cutting concerns (logging, resilience)
   - Reusable utilities
   - No business logic

## Prerequisites

- Go 1.21 or higher
- Git

## Getting Started

### 1. Clone and Navigate

```bash
cd api
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Environment

Copy the example environment file and adjust as needed:

```bash
cp .env.example .env
```

Edit `.env` with your configuration:

```env
ENVIRONMENT=development
PORT=8080
DATABASE_URL=postgresql://user:password@localhost:5432/jibe
JWT_SECRET=your-secret-key
```

### 4. Run the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

### 5. Build for Production

```bash
go build -o bin/api main.go
./bin/api
```

## API Endpoints

### Health Check

```bash
GET /health
```

Response:
```json
{
  "success": true,
  "data": {
    "status": "healthy"
  }
}
```

### Root

```bash
GET /
```

### API v1 Routes

#### User Management

```bash
GET    /api/v1/users         # List all users (with pagination)
POST   /api/v1/users         # Create a new user
GET    /api/v1/users/{id}    # Get a specific user
PUT    /api/v1/users/{id}    # Update a user
DELETE /api/v1/users/{id}    # Delete a user
```

#### Examples

**List users with pagination:**
```bash
curl "http://localhost:8080/api/v1/users?limit=10&offset=0"
```

**Create a user:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

**Get a user:**
```bash
curl http://localhost:8080/api/v1/users/1
```

**Update a user:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Smith","email":"john.smith@example.com"}'
```

**Delete a user:**
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Configuration

Configuration is managed through environment variables. See [.env.example](.env.example) for all available options.

### Key Configuration Options

| Variable | Description | Default |
|----------|-------------|---------|
| ENVIRONMENT | Runtime environment (development/production) | development |
| PORT | Server port | 8080 |
| READ_TIMEOUT | HTTP read timeout (seconds) | 15 |
| WRITE_TIMEOUT | HTTP write timeout (seconds) | 15 |
| IDLE_TIMEOUT | HTTP idle timeout (seconds) | 60 |
| SHUTDOWN_TIMEOUT | Graceful shutdown timeout (seconds) | 20 |
| ALLOWED_ORIGINS | CORS allowed origins | * |
| DATABASE_URL | Database connection string | - |
| JWT_SECRET | Secret key for JWT tokens | - |

## Development

### Running Tests

```bash
go test ./...
```

### Running with Hot Reload

Install air for hot reloading:

```bash
go install github.com/air-verse/air@latest
air
```

### Code Formatting

```bash
go fmt ./...
```

### Linting

```bash
go vet ./...
```

## Production Deployment

### Docker

Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

Build and run:

```bash
docker build -t jibe-api .
docker run -p 8080:8080 --env-file .env jibe-api
```

### Environment Variables

For production, ensure these variables are properly set:
- `ENVIRONMENT=production`
- `JWT_SECRET` (strong random string)
- `DATABASE_URL` (production database)
- `ALLOWED_ORIGINS` (specific domains, not *)

## Logging

The application uses structured logging with zerolog:

- **Development**: Pretty-printed console output
- **Production**: JSON formatted logs

Logs include:
- Request ID for tracing
- HTTP method, path, status code
- Response time
- Client IP address

## Design Patterns & Architecture

### Dependency Injection

The application uses a **Container pattern** for dependency injection ([internal/container/container.go](internal/container/container.go)):

```go
// Dependencies are injected through the container
container := container.New(cfg, log)

// Handlers receive their dependencies
userHandler := handlers.NewUserHandler(container.UserService, container.Logger)
```

Benefits:
- Loose coupling between components
- Easy to test with mock implementations
- Centralized dependency management
- Clear dependency graph

### Repository Pattern

Data access is abstracted through repository interfaces ([internal/repository/](internal/repository/)):

```go
type UserRepository interface {
    Create(ctx context.Context, user *domain.User) error
    GetByID(ctx context.Context, id int64) (*domain.User, error)
    // ... other methods
}
```

Implementations:
- **In-Memory** ([memory/](internal/repository/memory/)): For development and testing
- **PostgreSQL** ([postgres/](internal/repository/postgres/)): For production

This allows swapping implementations without changing business logic.

### Service Layer Pattern

Business logic is encapsulated in services ([internal/service/](internal/service/)):

```go
type UserService interface {
    CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error)
    GetUser(ctx context.Context, id int64) (*User, error)
    // ... other methods
}
```

Services:
- Coordinate between repositories
- Enforce business rules
- Handle validation
- Provide transaction boundaries

### Circuit Breaker Pattern

Prevents cascading failures ([pkg/resilience/circuit_breaker.go](pkg/resilience/circuit_breaker.go)):

```go
err := circuitBreaker.Execute(ctx, func() error {
    // Call external service
    return externalService.Call()
})
```

States:
- **Closed**: Normal operation, requests pass through
- **Open**: Too many failures, requests fail immediately
- **Half-Open**: Testing if service recovered

### Retry Pattern

Automatic retry with exponential backoff ([pkg/resilience/retry.go](pkg/resilience/retry.go)):

```go
err := retryPolicy.Execute(ctx, func() error {
    // Attempt operation
    return operation()
})
```

Features:
- Configurable max retries
- Exponential backoff
- Optional jitter to prevent thundering herd
- Context-aware cancellation

## Security

Built-in security features:
- Security headers (X-Frame-Options, X-Content-Type-Options, etc.)
- CORS configuration
- Panic recovery with stack trace logging
- Request timeouts at multiple levels
- Input validation in domain layer
- SQL injection prevention (parameterized queries)
- Unique request ID for tracing

## License

MIT

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
