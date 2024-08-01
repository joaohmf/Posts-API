package initializers

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectToDatabase() {
	var err error
	const file string = "./db/posts.db"
	DB, err = sql.Open("sqlite3", file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = DB.Ping(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
