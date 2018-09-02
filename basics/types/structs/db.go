package structs

import "database/sql"

type Database struct {
	db *sql.DB
}

var database Database
