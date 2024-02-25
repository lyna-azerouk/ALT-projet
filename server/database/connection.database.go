package database

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
	const format = "postgresql://%s:%s@%s:%d/%s?sslmode=require"
	psqlInfo := fmt.Sprintf(format, GetUser(), GetPassword(), GetHost(), GetPort(), GetDBName())
	return sql.Open("postgres", psqlInfo)
}
