-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET TIMEZONE="Europe/Moscow";

-- Create users table
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    email VARCHAR (254) NOT NULL UNIQUE,
    password_hash VARCHAR NOT NULL,
    username VARCHAR (13) NOT NULL UNIQUE,
    user_status INT NOT NULL,
    user_attrs JSONB NOT NULL
);

-- Create projects table
CREATE TABLE projects (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    author_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    alias VARCHAR (17) NOT NULL UNIQUE,
    project_status INT NOT NULL,
    project_attrs JSONB NOT NULL
);

-- Create tasks table
CREATE TABLE tasks (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    project_id UUID NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    task_status INT NOT NULL,
    task_attrs JSONB NOT NULL
);

-- Create tokens table
CREATE TABLE tokens (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    expired_at TIMESTAMP WITH TIME ZONE NOT NULL,
    access_token VARCHAR NOT NULL
);

-- Add indexes
CREATE INDEX active_users ON users (username) WHERE user_status = 1;
CREATE INDEX active_projects ON projects (alias) WHERE project_status = 1;
CREATE INDEX active_tasks ON tasks (id) WHERE task_status = 1;
