# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Jibe is a **monorepo** containing three main components:
- **API** (`api/`): Production-ready Go backend with clean architecture
- **Web** (`web/`): Next.js 16 frontend application
- **Hardhat** (`hardhat/`): Solidity smart contracts development environment

## Architecture

### API (Go Backend)

The API follows **Clean Architecture** with strict layer separation:

```
Handler Layer → Service Layer → Repository Layer → Domain Layer
```

**Key Principles:**
- Dependencies flow inward (Handler → Service → Repository → Domain)
- Domain layer has zero external dependencies
- All inter-layer communication uses interfaces
- Repository pattern allows swapping data stores (in-memory for dev, PostgreSQL for production)

**Directory Structure:**
- `internal/domain/` - Core entities and business rules (User, errors)
- `internal/repository/` - Data access interfaces with memory/ and postgres/ implementations
- `internal/service/` - Business logic and orchestration
- `internal/handlers/` - HTTP request/response handling
- `internal/middleware/` - Request ID, logging, CORS, panic recovery
- `internal/container/` - Dependency injection container that wires everything together
- `internal/config/` - Environment-based configuration
- `pkg/logger/` - Structured logging with zerolog
- `pkg/resilience/` - Circuit breaker and retry patterns
- `docs/` - Auto-generated Swagger documentation

**Important Files:**
- `main.go` - Application entry point, dependency wiring, route setup
- `ARCHITECTURE.md` - Detailed architectural documentation
- `README.md` - Comprehensive setup and API documentation

### Web (Next.js Frontend)

- Next.js 16 with App Router
- React 19 with TypeScript
- Tailwind CSS v4 for styling
- Directory: `web/app/` contains pages and routes

### Hardhat (Smart Contracts)

- Hardhat 3 Beta for Ethereum development
- Solidity 0.8.28
- Contracts in `hardhat/contracts/`
- Uses viem for Ethereum interactions
- Supports Foundry-compatible tests

## Common Development Commands

### API (from `api/` directory)

**Development:**
```bash
make run              # Run the application directly
make dev              # Run with hot reload (requires air)
make install          # Install dependencies (go mod download + tidy)
```

**Building:**
```bash
make build            # Build binary to bin/api
```

**Testing:**
```bash
make test             # Run all tests
make test-coverage    # Run tests with coverage report
```

**Code Quality:**
```bash
make fmt              # Format code with go fmt
make lint             # Run go vet
```

**Swagger Documentation:**
After modifying Swagger annotations in handlers, regenerate docs:
```bash
~/go/bin/swag init    # Regenerates docs/ directory
```

**Docker:**
```bash
make docker-build     # Build Docker image
make docker-run       # Run Docker container
```

### Web (from `web/` directory)

```bash
pnpm dev              # Start development server on localhost:3000
pnpm build            # Build for production
pnpm start            # Start production server
pnpm lint             # Run ESLint
```

### Hardhat (from `hardhat/` directory)

```bash
npx hardhat test                        # Run all tests
npx hardhat test solidity               # Run Solidity tests only
npx hardhat test nodejs                 # Run Node.js integration tests only
npx hardhat ignition deploy ignition/modules/Counter.ts  # Deploy locally
npx hardhat ignition deploy --network sepolia ignition/modules/Counter.ts  # Deploy to Sepolia
```

## Key Design Patterns (API)

### Dependency Injection
All dependencies are wired in `internal/container/container.go`. The Container struct holds all services, repositories, and infrastructure components. This is the single source of truth for dependency wiring.

### Repository Pattern
- Interface: `internal/repository/user_repository.go`
- Implementations: `memory/user_repository.go` (dev) and `postgres/user_repository.go` (production)
- Switch implementations in the container based on environment

### Service Layer
Services coordinate business logic and enforce business rules. They sit between handlers and repositories, ensuring handlers never directly access data stores.

### Middleware Stack
Request flow: RequestID → Logger → Recoverer → CORS → Router → Handler

Each request gets a unique ID for distributed tracing. All logs include this ID for correlation.

### Resilience Patterns
- **Circuit Breaker** (`pkg/resilience/circuit_breaker.go`): States are Closed → Open → Half-Open
- **Retry with Backoff** (`pkg/resilience/retry.go`): Exponential backoff with configurable max retries

## API Response Format

All API responses follow this structure:
```json
{
  "success": true/false,
  "data": { ... },
  "error": "error message if success is false"
}
```

## API Versioning

Current version: `v1`
Base path: `/api/v1`
All routes are prefixed with this base path.

## Error Handling (API)

Domain errors are defined in `internal/domain/errors.go`. Handlers translate these to HTTP status codes:
- `ErrNotFound` → 404
- `ErrAlreadyExists` → 409
- `ErrInvalidInput` → 400
- Unknown errors → 500 (logged, generic message to client)

## Configuration

### API Configuration
Environment variables loaded from `.env` file (copy from `.env.example`). Key variables:
- `ENVIRONMENT` - development/production
- `PORT` - Server port (default 8080)
- `DATABASE_URL` - PostgreSQL connection string
- `JWT_SECRET` - JWT signing key
- Timeout settings: `READ_TIMEOUT`, `WRITE_TIMEOUT`, `IDLE_TIMEOUT`, `SHUTDOWN_TIMEOUT`

### Hardhat Configuration
Configuration variables for deployment:
- `SEPOLIA_RPC_URL` - Sepolia RPC endpoint
- `SEPOLIA_PRIVATE_KEY` - Private key for deployment account

Set using: `npx hardhat keystore set SEPOLIA_PRIVATE_KEY`

## Testing Strategy (API)

- **Unit Tests**: Test services with mock repositories, test handlers with mock services
- **Integration Tests**: Test repositories with real database connections
- **E2E Tests**: Test full HTTP request flow

Run tests with `go test ./...` from the `api/` directory.

## Adding New Features (API)

When adding a new domain entity (e.g., "Product"):

1. **Domain**: Create `internal/domain/product.go` with the entity struct and validation
2. **Repository**: Create interface in `internal/repository/product_repository.go` and implementations in `memory/` and `postgres/`
3. **Service**: Create `internal/service/product_service.go` with business logic
4. **Handler**: Create `internal/handlers/product_handler.go` with HTTP handlers
5. **Wire Dependencies**: Update `internal/container/container.go` to instantiate new components
6. **Routes**: Add routes in `main.go` under the API v1 router
7. **Swagger**: Add Swagger annotations to handlers and run `~/go/bin/swag init`

## Swagger Documentation

Swagger is available at `/swagger/index.html` when the API is running. Annotations are in handler methods using swaggo format. After modifying annotations, regenerate with `~/go/bin/swag init`.

## Logging

- Uses `zerolog` for structured JSON logging
- Development: pretty-printed console output
- Production: JSON format
- Every log includes request ID for tracing
- Log levels: Debug, Info, Warn, Error, Fatal

## Package Manager

- **API**: Go modules (go.mod)
- **Web**: pnpm (pnpm-lock.yaml)
- **Hardhat**: pnpm (pnpm-lock.yaml)

## Go Version

Go 1.24.0 (toolchain 1.24.5) - specified in `api/go.mod`
