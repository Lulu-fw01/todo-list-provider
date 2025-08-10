# todo-list-provider
Project idea from Grok.

## Todo List API with User Authentication
Description: Build a RESTful API for a todo list application where users can create accounts, log in, manage their tasks (create, read, update, delete), and share tasks with others. This project focuses on core backend concepts like API design, authentication, and database interactions. Add features like task deadlines, priorities, and email notifications for overdue tasks to make it more challenging.

## Technologies Used
- **Golang**: Use the Gin framework for routing and middleware.
- **Database**: PostgreSQL for structured data storage (users and tasks tables with relationships).
- **Authentication**: JWT (JSON Web Tokens) for secure user sessions.
- **No ORM**: Raw SQL queries for database operations.

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
- Make (optional, for using Makefile commands)

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

3. **Run the application**
   ```bash
   # Using Makefile
   make run
   
   # Or manually
   go run cmd/server/main.go
   ```

4. **Run with PostgreSQL**
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

### Todos
- `GET /api/v1/todos` - Get all todos
- `GET /api/v1/todos/:id` - Get specific todo
- `POST /api/v1/todos` - Create new todo
- `PUT /api/v1/todos/:id` - Update todo
- `DELETE /api/v1/todos/:id` - Delete todo

## Database Schema

The application uses PostgreSQL with the following main tables:
- `users` - User accounts and authentication
- `todos` - Todo items with priorities and deadlines
- `shared_todos` - Task sharing between users

## Development

### Available Make Commands
- `make help` - Show available commands
- `make build` - Build the application
- `make run` - Run locally
- `make test` - Run tests
- `make clean` - Clean build artifacts
- `make deps` - Download dependencies
- `make docker-build` - Build Docker image
- `make docker-run` - Run with Docker Compose
- `make docker-stop` - Stop Docker services

### Adding New Features
1. Create models in `internal/models/`
2. Add handlers in `internal/handlers/`
3. Update database schema in `scripts/init.sql`
4. Add tests for new functionality

## Environment Variables

The application can be configured using environment variables:
- `PORT` - Server port (default: 8080)
- `HOST` - Server host (default: 0.0.0.0)
- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_NAME` - Database name (default: todo_db)
- `DB_USER` - Database user (default: todo_user)
- `DB_PASSWORD` - Database password (default: todo_password)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).
