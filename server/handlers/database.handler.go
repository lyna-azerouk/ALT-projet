package handlers

import (
	"database/sql"
	"fmt"
)

/**
 * ConnectDB
 * Connect to the database
 * should return two types
 * @return *sql.DB
 */
func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=require", USER_PROD, PASSWORD_PROD, HOST_PROD, PORT_PROD, DB_NAME_PROD)
	return sql.Open("postgres", psqlInfo)
}
