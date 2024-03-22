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
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("connect fail", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		log.Fatal("cant Delete", err)
	}

	if _, err := stmt.Exec(2); err != nil {
		log.Fatal("cannot execute delete", err)
	}
	fmt.Println("delete success")
}
