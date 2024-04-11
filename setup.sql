-- Connect to the database
\c task_manager;

-- Create the tasks table
CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       description TEXT,
                       completed BOOLEAN DEFAULT false
);

-- Insert initial data
INSERT INTO tasks (title, description) VALUES ('Buy groceries', 'Milk, bread, eggs');
INSERT INTO tasks (title, description) VALUES ('Clean the house', 'Vacuum, dust, mop');
INSERT INTO tasks (title, description, completed) VALUES ('Pay bills', 'Electricity, water, rent', true);