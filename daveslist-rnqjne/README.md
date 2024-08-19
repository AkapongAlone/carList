### Brief

After a few awkward experiences trying to sell your car on Craigslist, including that time you received a strange offer from a complete stranger which was totally unrelated to your old Toyota Camry, you decide to that there has to be a better way. So, you build it: *Daveslist - for used cars.*

Now it's time to build the API. In the application, users should be able to interact with each other by creating listings, replying to listings, and sending private messages. Users should also have different permissions within the app depending on their status (e.g. anonymous site visitor, registered user, moderator, and admin.) 

### Tasks

- Implement assignment using:
    - Language: *Go*
    - Framework: *any framework*
- Every site visitor should be able to see public listings in public categories.
- Registered users should be able to see all categories (both public and private) and all listings.
- Registered users should be able to create a listing, with the following options:
    - Select a category from existing categories
    - Select whether the listing should be public or private (visible to registered users only)
- Each listing contains title, content, and pictures (5000 characters limit for content, up to 10 pictures per listing)
- Replies can be text only (255 chars max)
- Registered users should be able to edit or delete their own listings.
- Registered users should be able to reply to any listing unless the original post is older than 1 year (to prevent "necrobumping")
- Registered user can send a private message to other users
- Moderators should be able to temporarily hide (but not edit or delete) any listing.
- Moderators can create or delete categories (on deleting category all listings in that category should not be permanently deleted, but rather moved to "trash bin")
- Admin (superuser) can do everything the moderator can do, plus they can edit and delete any listing.

Your task is to build an HTTP API that will provide the functionality above. 
You should write unit tests for business logic. 
You are expected to design any other required models and routes for your API.

 ### Evaluation Criteria

 - *Go* best practices
 - Completeness: Did you include all features?
 - Correctness: Does the solution work in sensible, thought-out ways?
 - Maintainability: Is the code written in a clean, maintainable way?
 - Testing: Is the solution adequately tested?
 - Documentation: Is the API well-documented?

### CodeSubmit


---

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


