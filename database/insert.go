package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	log.Printf(url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("connect fail", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO users (name,age) values ($1 , $2) RETURNING id", "PAEFRIENDMICK", 25)
	var id int
	err = row.Scan(&id)

	if err != nil {
		log.Fatal("Can't Insert data", err)
	}
	fmt.Println("INSERT COMPLETED", id)
}
