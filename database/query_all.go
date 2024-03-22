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

	stmt, err := db.Prepare("SELECT id , name , age FROM users")
	if err != nil {
		log.Fatal("cant prepare query all", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("cannot query rows all users", err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("cannot scan rows", err)
		}
		fmt.Println(id, name, age)
	}
	fmt.Println("Query completed")
}
