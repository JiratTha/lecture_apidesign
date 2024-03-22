package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log" // Added missing import for the log package
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "Jirat", Age: 20},
	{ID: 2, Name: "Mick", Age: 20},
}

func main() {

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			log.Println("GET")
			b, err := json.Marshal(users)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(b)
			return
		}
		if r.Method == "POST" {
			log.Println("POST")
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			var u User
			err = json.Unmarshal(body, &u)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			users = append(users, u)

			fmt.Fprintf(w, "hello %s created users", "POST")
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})
	// Fixed: Moved the closing parenthesis for the anonymous function to the correct place.
	log.Fatal(http.ListenAndServe(":2565", nil))
}
