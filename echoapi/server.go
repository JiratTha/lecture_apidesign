package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"log" // Added missing import for the log package
	"net/http"
	"os"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")

}

type Logger struct {
	Handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("Server http middleware : %s %s %s %s", r.RemoteAddr, r.Method, r.URL, time.Since(start))
}

func getUsersHandler(c echo.Context) error {
	stmt, err := db.Prepare("SELECT id , name , age FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	users := []User{}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		users = append(users, u)
	}
	return c.JSON(http.StatusOK, users)
}

func createUsersHandler(c echo.Context) error {

	var u User
	err := c.Bind(&u)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	row := db.QueryRow("INSERT INTO users (name,age) values ($1 , $2) RETURNING id", u.Name, u.Age)

	err = row.Scan(&u.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, u)

}

var db *sql.DB

func main() {
	var err error
	url := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("connect fail", err)
	}
	defer db.Close()

	createTB := `CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY , name TEXT , age INT);`
	_, err = db.Exec(createTB)
	if err != nil {
		log.Fatal("Can't create TABLE", err)
	}
	{
		log.Println("OKAY")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	g := e.Group("/api")
	g.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == "apidesign" && password == "45678" {
			return true, nil
		}
		return false, nil
	}))
	g.GET("/users", getUsersHandler)
	g.POST("/users", createUsersHandler)

	// Fixed: Moved the closing parenthesis for the anonymous function to the correct place.
	log.Fatal(e.Start(":2565"))
}
