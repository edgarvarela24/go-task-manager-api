# Task Manager API

This is a RESTful API for managing tasks built with Go and PostgreSQL. Using the native Go http library and GORM.

## Prerequisites

- Go (version 1.22.1 or later)
- PostgreSQL

## Setup

- Clone the repository:
    ``` 
    git clone https://github.com/your-username/task-manager-api.git
    ```
- Change into the project directory:
    ```
    cd task-manager-api
    ```
  
- Create a `config.yml` file in the project root directory with the following content:
    ```yaml
    database:
      host: localhost
      port: 5432
      user: your_username
      password: your_password
      name: your_database
    ```
  Replace `your_username`, `your_password`, and `your_database` with your specific PostgreSQL database credentials.


- Set up the database:
  - Install PostgreSQL and ensure it is running.
  - Create a new database with the name you specified in the config.yml file.
  - Run the setup.sql script to create the required table and insert initial data:
    ```
    psql -U your_username -d your_database -f setup.sql
    ```

    Replace `your_username` and `your_database` with the values from the config.yml file.

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