#!/bin/bash

# Define PostgreSQL connection parameters
DB_HOST="frost-hippo-13790.8nj.gcp-europe-west1.cockroachlabs.cloud"
DB_PORT="26257"
DB_NAME="defaultdb"
DB_USER="bouffluence"
DB_PASSWORD="NKi9yHEPNbAY-_MrwE8IRw"

SQL_QUERY_TYPE_CLIENT="CREATE TYPE USER_ROLE AS ENUM ('CLIENT', 'RESTAURANT', 'ADMIN');"
SQL_QUERY_TYPE_ORDER="CREATE TYPE ORDER_STATUS AS ENUM ('PENDING', 'IN_PROGRESS', 'COMPLETED', 'DECLINED');"
SQL_QUERY_TYPE_AFLUENCE="CREATE TYPE AFFLUENCE_LEVEL AS ENUM ('LOW', 'MODERATE', 'HIGH', 'VERY HIGH');"

CREATE_TABLE1="CREATE TABLE \"User\" (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(300) NOT NULL,
    user_role USER_ROLE NOT NULL
);"


CREATE_TABLE2="CREATE TABLE \"Restaurant\" (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(300) NOT NULL,
  affluence AFFLUENCE_LEVEL NOT NULL,
  cuisine_type VARCHAR(300)
);"


CREATE_TABLE3="CREATE TABLE \"Menus\" (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL,
  restaurant_id INT,
  price INT NOT NULL,
  description VARCHAR(255),
  image VARCHAR(255),
  FOREIGN KEY (restaurant_id) REFERENCES Restaurant(id)
);"


CREATE_TABLE4="CREATE TABLE \"Order\" (
  id SERIAL PRIMARY KEY,
  restaurant_id INT,
  client_id INT,
  price INT NOT NULL,
  FOREIGN KEY (restaurant_id) REFERENCES Restaurant(id),
  FOREIGN KEY (client_id) REFERENCES USER(id)
);"



# Insert data into the table
INSERT_DATA="INSERT INTO \"USER\" (email, password, user_role) VALUES
    ('example1@example.com', 'password1', 'CLIENT');"

# Connect to CockroachDB and execute SQL commands
 cockroach sql --host="$DB_HOST" --user="$DB_USER" --database="$DB_NAME" -e "$SQL_QUERY_TYPE_CLIENT"
 cockroach sql --host="$DB_HOST" --user="$DB_USER" --database="$DB_NAME" -e "$SQL_QUERY_TYPE_ORDER"
 cockroach sql --host="$DB_HOST" --user="$DB_USER" --database="$DB_NAME" -e "$SQL_QUERY_TYPE_AFLUENCE"
 cockroach sql --host="$DB_HOST" --user="$DB_USER" --database="$DB_NAME" -e "$CREATE_TABLE1"
 cockroach sql --host="$DB_HOST" --user="$DB_USER" --database="$DB_NAME" -e "$CREATE_TABLE2"
 cockroach sql --host="$DB_HOST" --user="$DB_USER" --database="$DB_NAME" -e "$CREATE_TABLE3"
 cockroach sql --host="$DB_HOST" --user="$DB_USER" --database="$DB_NAME" -e "$CREATE_TABLE4"
# password: NKi9yHEPNbAY-_MrwE8IRw