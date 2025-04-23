# Base Go Backend

A modern Golang backend application built with clean architecture principles, providing a solid foundation for building scalable web services.

## Features

- **Clean Architecture**: Structured using domain-driven design principles
- **RESTful API**: Built with Gin web framework
- **Dependency Injection**: Uses Google Wire for dependency management
- **Database**: MySQL integration with GORM
- **Caching**: Redis for performance optimization
- **Message Broker**: Kafka integration (optional)
- **API Documentation**: Swagger for API documentation
- **Containerization**: Docker and Docker Compose support
- **Logging**: Structured logging with Zap

## Getting Started

### Prerequisites

- Go 1.23+
- Docker and Docker Compose
- Make (optional, for convenience commands)

### Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   cd base_go_be
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Start required services using Docker:
   ```
   docker-compose up -d mysql redis
   ```
   
   Or simply:
   ```
   make mysql
   ```

### Running the Application

Run the application locally:
```
make run
```

Or using Docker:
```
docker-compose up
```

### API Testing

Test a sample endpoint:
```
curl http://localhost:8386/v1/user/test
```

Visit Swagger documentation:
```
http://localhost:8386/docs/index.html
```

## Development

### Dependency Injection

This project uses Google Wire for dependency injection:

1. Install Wire if not already installed:
   ```
   go install github.com/google/wire/cmd/wire@latest
   ```

2. Generate dependency injection code:
   ```
   cd internal/wire
   wire
   ```

### API Documentation

Generate Swagger documentation:
```
make swag
```
or
```
swag init -g ./cmd/server/main.go -o docs
```

## Project Structure

- `cmd/`: Application entry points
- `config/`: Configuration files
- `docs/`: API documentation
- `internal/`: Application core code
  - `controller/`: HTTP handlers
  - `service/`: Business logic
  - `repo/`: Data access layer
  - `model/`: Domain models
  - `middlewares/`: HTTP middlewares
  - `routers/`: API routes
- `migrations/`: Database migrations
- `pkg/`: Reusable packages
- `tests/`: Test files

## Deployment

The application can be deployed using Docker:

```
docker build -t base_go_be .
docker run -p 8386:8386 base_go_be
```

Or using Docker Compose:

```
docker-compose up -d
```