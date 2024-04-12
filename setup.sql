-- Connect to the database
\c task_manager;

-- Create the tasks table
CREATE TABLE IF NOT EXISTS tasks (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR(255) NOT NULL,
                        description TEXT,
                        completed BOOLEAN DEFAULT false
);

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
                        id SERIAL PRIMARY KEY,
                        username VARCHAR(255) NOT NULL UNIQUE,
                        email VARCHAR(255) NOT NULL UNIQUE,
                        password VARCHAR(255) NOT NULL,
                        role VARCHAR(50) DEFAULT 'user'
);