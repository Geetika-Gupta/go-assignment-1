package migration

import (
	"fmt"

	"github.com/go-pg/pg"
)

//Connect creates a database connection.
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "localhost:5432",
	}

	db := pg.Connect(opts)

	if db == nil {
		fmt.Println("Database Connection failed")
		return nil
	}

	fmt.Println("Database Connection Success")
	return db
}

//Close database connection.
func Close(db *pg.DB) {
	err := db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Database Connection Closed Success")
}
