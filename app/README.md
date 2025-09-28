# Todo List Service

A RESTful API service for managing todo lists built with Go, Gin, GORM, and PostgreSQL.

## Features

- User registration and authentication with JWT
- CRUD operations for todos
- Todo filtering and pagination
- User-specific todo management
- Docker containerization
- PostgreSQL database with GORM ORM

## Project Structure

```
app/
├── config/          # Database configuration
├── handlers/        # HTTP request handlers
├── middleware/      # Authentication middleware
├── models/          # Data models (User, Todo)
├── utils/           # Utility functions (auth, password hashing)
├── main.go          # Application entry point
├── go.mod           # Go module dependencies
├── Dockerfile       # Docker configuration
└── docker-compose.yml # Docker Compose configuration
```

## API Endpoints

### Authentication

- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login user

### Users

- `GET /api/v1/users/:id` - Get user profile (public)
- `GET /api/v1/me` - Get current user profile (protected)

### Todos

- `POST /api/v1/todos` - Create a new todo (protected)
- `GET /api/v1/todos` - Get user's todos with filtering and pagination (protected)
- `GET /api/v1/todos/:id` - Get specific todo (protected)
- `PUT /api/v1/todos/:id` - Update todo (protected)
- `DELETE /api/v1/todos/:id` - Delete todo (protected)
- `GET /api/v1/todos/stats` - Get todo statistics (protected)

### Health Check

- `GET /health` - Service health check

## Data Models

### User
```json
{
  "id": 1,
  "username": "john_doe",
  "email": "john@example.com",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Todo
```json
{
  "id": 1,
  "title": "Complete project",
  "description": "Finish the todo list API",
  "status": "pending",
  "priority": 3,
  "due_date": "2023-12-31T23:59:59Z",
  "user_id": 1,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

## Quick Start

### Using Docker Compose (Recommended)

1. Clone the repository and navigate to the app directory
2. Run the services:
   ```bash
   docker-compose up --build
   ```

The API will be available at `http://localhost:8080`

### Manual Setup

1. Install PostgreSQL and create a database named `todolist`
2. Set environment variables:
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=password
   export DB_NAME=todolist
   export DB_SSLMODE=disable
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_HOST | Database host | localhost |
| DB_PORT | Database port | 5432 |
| DB_USER | Database user | postgres |
| DB_PASSWORD | Database password | password |
| DB_NAME | Database name | todolist |
| DB_SSLMODE | SSL mode | disable |
| GIN_MODE | Gin mode (debug/release) | debug |

## API Usage Examples

### Register a new user
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create a todo (requires authentication)
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Complete project",
    "description": "Finish the todo list API",
    "priority": 3,
    "due_date": "2023-12-31T23:59:59Z"
  }'
```

### Get todos with filtering
```bash
curl -X GET "http://localhost:8080/api/v1/todos?status=pending&priority=3&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Update a todo
```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "status": "completed"
  }'
```

## Todo Status Values

- `pending` - Todo is not completed
- `completed` - Todo is completed

## Priority Levels

- `1` - Low priority
- `2` - Medium-low priority
- `3` - Medium priority (default)
- `4` - Medium-high priority
- `5` - High priority

## Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```

## Security Notes

- JWT secret key should be changed in production
- Use HTTPS in production
- Consider implementing rate limiting
- Validate and sanitize all user inputs
- Use environment variables for sensitive configuration

## License

This project is open source and available under the MIT License.
