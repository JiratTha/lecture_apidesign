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

	stmt, err := db.Prepare("SELECT id , name , age FROM users where id=$1")
	if err != nil {
		log.Fatal("cant prepare query one", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var id, age int
	var name string
	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal("cannot scan row", err)
	}
	fmt.Println(id, name, age)
}
