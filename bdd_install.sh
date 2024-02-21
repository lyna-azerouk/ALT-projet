#!/bin/bash

# Define PostgreSQL connection parameters
DB_HOST="bouffluence-4322.g95.gcp-us-west2.cockroachlabs.cloud"
DB_PORT="26257"
DB_NAME="bouffluence"
DB_USER="bouffluence"
DB_PASSWORD="gTsPKkviQpqV3wl6JYeiOw"


SQL_QUERY_TYPE="CREATE TYPE USER_ROLE AS ENUM ('CLIENT', 'RESTAURANT', 'ADMIN');"

CREATE_TABLE1="CREATE TABLE \"USER\" (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(300) NOT NULL,
    user_role USER_ROLE NOT NULL
);"

# Insert data into the table
INSERT_DATA="INSERT INTO \"USER\" (email, password, user_role) VALUES
    ('example1@example.com', 'password1', 'CLIENT');"

# Connect to CockroachDB and execute SQL commands
cockroach sql --host="$DB_HOST" --user="$DB_USER" --database="$DB_NAME" -e "$SQL_QUERY_TYPE"
cockroach sql --host="$DB_HOST" --user="$DB_USER" --database="$DB_NAME" -e "$CREATE_TABLE1"
cockroach sql --host="$DB_HOST" --user="$DB_USER" --database="$DB_NAME" -e "$INSERT_DATA"