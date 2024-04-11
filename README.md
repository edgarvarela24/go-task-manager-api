# Task Manager API

This is a RESTful API for managing tasks built with Go and PostgreSQL. Using the native Go http library and GORM.

## Prerequisites

- Docker Compose

## Setup

- Clone the repository:
    ``` 
    git clone https://github.com/your-username/task-manager-api.git
    ```
- Change into the project directory:
    ```
    cd task-manager-api
    ```

## Running the Application

- Start the application:
    ```
    # Start services
    docker-compose up -d

    # Stop services
    docker-compose down
    ```

- The API will be accessible at `http://localhost:8080`.

## API Endpoints
- `GET /tasks` - Retrieve all tasks
- `GET /tasks/{id}` - Retrieve a specific task by ID
- `POST /tasks` - Create a new task
- `PUT /tasks/{id}` - Update a task
- `DELETE /tasks/{id}` - Delete a task