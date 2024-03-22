package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	url := "postgres://wpczwjti:LP1X19Q9vf7Nq0aPprKfYhdvK6Rfkb5w@rain.db.elephantsql.com/wpczwjti"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("connect fail", err)
	}
	defer db.Close()
	createTB := `CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY , name TEXT , age INT);`
	_, err = db.Exec(createTB)
	if err != nil {
		log.Fatal("Can't create TABLE", err)
	}
	log.Println("OKAY")
}
