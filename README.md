# Project Documentation

## 1. Deploy Locally

### - Prerequisites in your local computer
- Git
- Docker
- Go (Optional)

### - Steps
1. Clone this repository to your local machine.
2. Start the Docker engine.
3. Navigate to the source code directory.
4. Run the following command to start all services.
```bash
   docker compose up -d
```

### - Services
1. API App – API built with Go (Gin framework)
2. Postgres – Database
3. Adminer – Admin interface to view data
4. Nginx – Reverse proxy for the API

### - Using the API
- This API provides a single endpoint to check strong password steps:
```bash
   /api/strong_password_steps
```

1. Access the API at http://localhost/api/strong_password_steps.
2. Set the password you want to check in the request body in JSON format:
```json
{
    "init_password": "testpassword"
}
```
3. You will receive a response indicating the number of steps needed to make the password strong:
```json
{
    "num_of_steps": 2
}
```

---

## 2. Run Unit Tests
Unit tests can be executed in two ways:

1. During the build process of the Go Docker image.
2. Manually on your local machine, ensure that Go is installed, and then use the following command:
```bash
    go mod tidy
    go test -v
```