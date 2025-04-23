# base_go_be - Go Backend Service
# base_go_be Project Documentation

This document provides an overview of the project structure and functionality to aid in maintenance and development.

## Table of Contents

- [Project Overview](#project-overview)
- [Directory Structure](#directory-structure)
- [Key Components](#key-components)
- [API Endpoints](#api-endpoints)
- [Setup & Installation](#setup--installation)
- [Development Workflow](#development-workflow)

## Project Overview

This project is a modular Go backend application built using the Gin web framework. It implements a clean architecture with organized components for routing, business logic, and data access. The project is designed to be scalable, maintainable, and easy to extend.

## Directory Structure

The following is the directory structure of the project along with a description of each directory's purpose:

```
base_go_be/
├── cmd/                # Entry points for the application (e.g., main.go)
├── config/             # Configuration files and utilities
├── internal/           # Internal application code (not exposed as a library)
│   ├── routes/         # API route definitions
│   ├── handlers/       # HTTP request handlers
│   ├── services/       # Business logic and service layer
│   ├── repositories/   # Data access layer (e.g., database queries)
│   ├── models/         # Data models and structures
│   └── utils/          # Utility functions and helpers
├── pkg/                # Shared packages that can be reused across projects
├── docs/               # Documentation files (e.g., API specs, design docs)
├── migrations/         # Database migration files
├── scripts/            # Helper scripts for automation (e.g., build, deploy)
├── test/               # Test cases and test utilities
└── vendor/             # Vendor dependencies (if vendoring is used)
```

### Directory Details

- **cmd/**: Contains the main entry point(s) for the application. For example, `main.go` initializes the application and starts the server.
- **config/**: Stores configuration files and utilities for loading and managing environment-specific settings.
- **internal/**: Contains the core application logic. This directory is further divided into:
  - **routes/**: Defines API routes and maps them to handlers.
  - **handlers/**: Implements HTTP request handlers for processing incoming requests.
  - **services/**: Contains the business logic and service layer, separating it from the HTTP layer.
  - **repositories/**: Handles data access, such as database queries or external API calls.
  - **models/**: Defines data models and structures used throughout the application.
  - **utils/**: Provides utility functions and helpers for common tasks.
- **pkg/**: Contains reusable packages that can be shared across projects or modules.
- **docs/**: Includes project documentation, such as API specifications, architecture diagrams, and design documents.
- **migrations/**: Stores database migration files for schema changes.
- **scripts/**: Contains scripts for automating tasks like building, testing, and deploying the application.
- **test/**: Includes test cases and utilities for unit, integration, and end-to-end testing.
- **vendor/**: Stores third-party dependencies if vendoring is used instead of Go modules.

## Key Components

- **Gin Framework**: Used for routing and handling HTTP requests.
- **Clean Architecture**: Ensures separation of concerns between layers (e.g., routing, business logic, and data access).
- **Database Integration**: Supports database operations through the repository pattern.
- **Configuration Management**: Handles environment-specific settings using configuration files.

## API Endpoints

The API endpoints are defined in the `routes/` directory and implemented in the `handlers/` directory. Refer to the `docs/` folder for detailed API specifications.

## Setup & Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/base_go_be.git
   cd base_go_be
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
   - Copy the `.env.example` file to `.env` and update the values as needed.

4. Run database migrations:
   ```bash
   go run scripts/migrate.go
   ```

5. Start the application:
   ```bash
   go run cmd/main.go
   ```

## Development Workflow

1. **Branching**: Use feature branches for new features and bug fixes.
2. **Code Style**: Follow Go best practices and lint your code using `golangci-lint`.
3. **Testing**: Write unit tests for all new code and run tests using `go test ./...`.
4. **Code Review**: Submit pull requests for review before merging into the main branch.
5. **Deployment**: Use CI/CD pipelines for automated testing and deployment.

