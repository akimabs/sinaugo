package main

import (
	db "sinaugo/src/database"
	"sinaugo/src/database/seed"

	"github.com/sirupsen/logrus"
)

func main () {
	db.SeedConnection()
	conn := db.GetDB()

	for _, seed := range seed.All(){
		if err := seed.Run(conn); err != nil {
			logrus.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}