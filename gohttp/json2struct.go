package main

import (
	"encoding/json"
	"fmt"
)

type Users struct {
	ID   int
	Name string
	Age  int
}

func main() {
	data := []byte(`{
	"id" : 2 , 
	"name" : "Tonson",
	"age" : 23
 
	}`)
	var u Users
	err := json.Unmarshal(data, &u)

	fmt.Println(err)
}
