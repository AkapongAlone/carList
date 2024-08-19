

# My Go Project

## Overview

This project is a basic Go application with Swagger documentation support. It includes a local database setup and some pre-populated master data.

## Getting Started

### 1. Install Required Packages

First, install the necessary packages for Swagger and other dependencies:

- **Install the Swag CLI tool:**

  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

- **Install the Swaggo library:**

  ```bash
  go get -u github.com/swaggo/swag
  ```

### 2. Generate Swagger API Documentation

After installing the necessary tools, generate the Swagger documentation:

```bash
swag init
```

### 3. Run the Project

Navigate to your project directory and start the Go server:

```bash
cd my-go-project
go run .
```

### 4. Access Swagger UI

Once the server is running, you can access the Swagger UI in your web browser:

[http://localhost:8080/docs/index.html#/](http://localhost:8080/docs/index.html#/)

### 5. Setup Environment Variables

Copy the `.env.example` file to `.env` and fill in the required keys and values.

### 6. Database Setup

The project uses a local database with some pre-populated master data. If you are using a different database, you should:

- Check the default master data for roles in `const.go`.
- Insert role data according to the values in `const.go`.
- Alternatively, you can run the `MockRole()` function in the `databases` package to populate the roles.

## Troubleshooting

### Error Importing Packages

If you encounter any issues with importing packages, try running:

```bash
go mod tidy
```

This command cleans up and synchronizes your `go.mod` file.

---


