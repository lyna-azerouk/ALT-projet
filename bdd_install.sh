#!/bin/bash

# Define PostgreSQL connection parameters
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="data_boufluence"
DB_USER="" #replace with ur db user_name
DB_PASSWORD=""

DROP_TABLE1="DROP TABLE IF EXISTS authentication;"

# Define SQL commands to create tables
CREATE_TABLE1="CREATE TABLE authentication (
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(300) NOT NULL
);"

# insert into bdd
INSERT_DATA="INSERT INTO authentication (email, password) VALUES
    ('example1@example.com', 'password1'),
    ('example2@example.com', 'password2');"

# Connect to PostgreSQL and execute SQL commands
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "$DROP_TABLE1"
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "$CREATE_TABLE1"
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "$INSERT_DATA"
