# todo-list-provider
Project idea from Grok.

This project demonstrates 
- routing using gin framework
- middleware
- jwt token usage
- crud sql operations.

## Todo List API with User Authentication
RESTful API for a todo list application where users can create accounts, log in, manage their tasks (create, read, update, delete). This project focuses on core backend concepts like API design, authentication, and database interactions.

## Project Structure
```
todo-list-provider/
├── cmd/
│   └── server/          # Main application entry point
├── internal/
│   ├── handlers/        # HTTP request handlers
│   ├── models/          # Data models and structs
│   ├── database/        # Database connection and operations
│   └── auth/            # Authentication middleware
├── pkg/
│   └── middleware/      # Custom middleware
├── configs/             # Configuration management
├── scripts/             # Database initialization scripts
├── Dockerfile           # Application container
├── docker-compose.yml   # PostgreSQL and app services
├── Makefile            # Build and run commands
└── go.mod              # Go module dependencies
```

## Quick Start

### Prerequisites
- Go 1.21 or later
- Docker and Docker Compose

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd todo-list-provider
   ```

2. **Install dependencies**
   ```bash
   go mod download
   go mod tidy
   ```

3. **Run the application locally**
   ```bash
   # Using Makefile
   make run
   
   # Or manually
   go run cmd/server/main.go
   ```

4. **Run with docker compose**
   ```bash
   # Start PostgreSQL
   docker-compose up -d postgres
   
   # Wait a few seconds for database to initialize
   sleep 5
   
   # Start the application
   docker-compose up -d app
   ```

### Docker Setup

1. **Build and run everything**
   ```bash
   make docker-run
   ```

2. **Stop services**
   ```bash
   make docker-stop
   ```

3. **View logs**
   ```bash
   make docker-logs
   ```

## API Endpoints

### Health Check
- `GET /api/v1/health` - Check API status

### Tasks
- `GET /api/v1/task` - Get all task
- `GET /api/v1/task/:id` - Get specific task
- `POST /api/v1/task` - Create new task
- `PUT /api/v1/task/:id` - Update task
- `DELETE /api/v1/task/:id` - Delete task

## Environment Variables

The application can be configured using environment variables:
- `PORT` - Server port (default: 8080)
- `HOST` - Server host (default: 0.0.0.0)
- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_NAME` - Database name (default: todo_db)
- `DB_USER` - Database user (default: todo_user)
- `DB_PASSWORD` - Database password (default: todo_password)

todo add other...

## License

This project is open source and available under the [MIT License](LICENSE).
