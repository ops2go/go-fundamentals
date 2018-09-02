package basic

import (
	"database/sql"
	"log"
)

func init() {
	database.db, err = sql.Open("sqlite3", "./tasks.db")
	taskStatus = map[string]int{"COMPLETE": 1, "PENDING": 2, "DELETED": 3}
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	database.db.Close()
}
