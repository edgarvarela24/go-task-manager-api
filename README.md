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
  ```
  $ curl http://localhost:8080/tasks
  
  [{"id":1,"title":"Buy groceries","description":"Milk, bread, eggs","completed":false},{"id":2,"title":"Clean the house","description":"Vacuum, dust, mop","completed":false},{"id":3,"ti
  tle":"Pay bills","description":"Electricity, water, rent","completed":true}]
  ```
- `GET /tasks/{id}` - Retrieve a specific task by ID
- `POST /tasks` - Create a new task
  ```
  $ curl -X POST -H "Content-Type: application/json" -d '{"title":"Learn Go","description":"Practice Go programming"}' http://localhost:8080/tasks
  {"id":4,"title":"Learn Go","description":"Practice Go programming","completed":false}
  ```
- `PUT /tasks/{id}` - Update a task
- `DELETE /tasks/{id}` - Delete a task