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

	stmt, err := db.Prepare("UPDATE users SET name=$2 where id=$1;")
	if err != nil {
		log.Fatal("cant prepare stmt update", err)
	}

	if _, err := stmt.Exec(2, "Mick"); err != nil {
		log.Fatal("cannot execute update", err)
	}
	fmt.Println("update success")
}
