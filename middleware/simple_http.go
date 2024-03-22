package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log" // Added missing import for the log package
	"net/http"
	"time"
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

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
		return
	}
	if r.Method == "POST" {

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
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

//func logMiddleware(next http.HandlerFunc) http.HandlerFunc { // สามารถนำมาเลือกใช้กับบาง function ได้
//	return func(writer http.ResponseWriter, request *http.Request) {
//		start := time.Now()
//		next.ServeHTTP(writer, request)
//		log.Printf("Server http middleware : %s %s %s %s", request.RemoteAddr, request.Method, request.URL, time.Since(start))
//	}
//}

type Logger struct {
	Handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("Server http middleware : %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(401)
			w.Write([]byte(`can't parse the basic auth`))
			return
		}
		if u != "apidesign" || p != "45678" {
			w.WriteHeader(401)
			w.Write([]byte(`Username/Password incorrect.`))
			return
		}
		fmt.Println("Auth passed.")
		next(w, r)
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", AuthMiddleware(userHandler))
	mux.HandleFunc("/health", healthHandler)
	logMux := Logger{Handler: mux}
	srv := http.Server{
		Addr:    ":2565",
		Handler: logMux,
	}
	// Fixed: Moved the closing parenthesis for the anonymous function to the correct place.
	log.Fatal(srv.ListenAndServe())
}
