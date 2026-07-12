# CloudLens

A production-oriented, cloud-native e-commerce platform built in Go with a clean-architecture foundation, container-first deployment, and service-oriented structure.

## What this repository includes

This initial milestone establishes the platform foundation:

- Go module and repository structure for a multi-service backend
- Shared configuration loading with Viper
- Structured logging with Zap
- Health checks for liveness and readiness
- An API gateway entrypoint
- Docker Compose setup for core infrastructure
- Container build support for the gateway service
- Make targets for testing and local workflows

## Architecture overview

The project is organized to support a modular microservice architecture:

- `cmd/` contains service entrypoints
- `internal/` contains application-layer implementation
- `pkg/` contains reusable infrastructure packages
- `configs/` stores configuration defaults and environment templates
- `deployments/` contains Docker and Kubernetes assets
- `migrations/` is reserved for database migrations
- `monitoring/` is reserved for observability configuration

## Technology stack

- Language: Go 1.26+
- Configuration: Viper
- Logging: Zap
- HTTP routing: net/http
- API Documentation: Swagger/OpenAPI
- Containerization: Docker / Docker Compose
- Database foundation: PostgreSQL
- Cache foundation: Redis
- Messaging foundation: Kafka

## Prerequisites

Before running the project locally, make sure you have:

- Go 1.26 or newer
- Docker and Docker Compose
- Make (optional, but recommended)

## Getting started

### 1. Clone and enter the repository

```bash
git clone <repository-url>
cd cloud-native-ecommerce
```

### 2. Run tests

```bash
go test ./...
```

### 3. Start the API gateway locally

```bash
go run ./cmd/api-gateway
```

The service will listen on port `8080` by default.

### 4. Verify health endpoints

```bash
curl http://localhost:8080/health/live
curl http://localhost:8080/health/ready
```

## API Documentation

The API is fully documented with Swagger/OpenAPI. Once the server is running, access the interactive documentation at:

```
http://localhost:8080/swagger/
```

This provides a searchable, interactive interface where you can:
- View all available endpoints
- See request/response formats
- Test endpoints directly
- View status codes and error handling

### Available endpoints

- `GET /` - Welcome message
- `GET /health/live` - Liveness probe (returns "ok")
- `GET /health/ready` - Readiness probe (returns "ready")
- `GET /swagger/` - Interactive Swagger UI
- `GET /swagger/doc.json` - OpenAPI specification (JSON)

## Running with Docker Compose

```bash
docker compose up --build
```

This starts the core infrastructure services and the API gateway container.

## Development workflow

### Building and running scripts (Windows/PowerShell)

- `.\build.bat` - Build the binary
- `.\test.bat` - Run all tests
- `.\run.bat` - Build and run the API gateway
- `.\docker-up.bat` - Start Docker services
- `.\docker-down.bat` - Stop Docker services

### Updating API Documentation

When you add new endpoints or modify existing ones:

1. Add Swagger documentation comments to your handler functions:

```go
// handleExample godoc
// @Summary Brief description
// @Description Longer description
// @Produce json
// @Success 200 {object} YourType
// @Router /path [get]
func handleExample(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

2. Regenerate the Swagger documentation:

```bash
swag init -g cmd/api-gateway/main.go
```

3. Rebuild and test:

```bash
.\build.bat
.\test.bat
```

The Swagger UI at `/swagger/` will automatically reflect the changes.

## Useful Make targets

```bash
make test
make build
docker compose up --build
docker compose down -v
```

## Current status

This repository currently provides the foundation for the platform. The next development milestones will add:

- Authentication and user management
- Product and catalog services
- Cart, order, and payment flows
- Event-driven integration and observability

## Contributing

Contributions should keep the codebase aligned with the principles of:

- Clean architecture
- SOLID design
- Testability
- Secure configuration
- Production-ready deployment practices

