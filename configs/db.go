package configs

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	controllers "github.com/nilerajput91/gingonic/controllers"
)

//Connect func use to connect the db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "nilesh",
		Password: "nilesh",
		Addr:     "localhost:5432",
		Database: "nileshdb",
	}

	var db *pg.DB = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connect to db is successfull")

	controllers.CreateTodoTable(db)

	controllers.InitiateDB(db)

	return db

}
