package main

import (
	db "sinaugo/src/database"
	"sinaugo/src/database/migration"
)
func main() {
	db.AppConnection()
	conn := db.GetDB()
	db, err := conn.DB()
	if err != nil {
		defer db.Close()
	}

	migration.CreateTodo(conn)
}