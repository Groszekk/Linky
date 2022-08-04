package scripts

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DB-HOST")
	port     = 5432
	user     = os.Getenv("DB-USER")
	password = os.Getenv("DB-PASS")
	database = os.Getenv("DB-NAME")
)

func Connect() *sql.DB {
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	db, err := sql.Open("postgres", psql)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
