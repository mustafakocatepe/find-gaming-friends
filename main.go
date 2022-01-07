package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mustafakocatepe/find-gaming-friends/controller"
	"github.com/mustafakocatepe/find-gaming-friends/db"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	a := App{}

	// Initialize storage
	a.initialize()

	// Initialize routes
	a.routes()

	// Start server
	a.run(os.Getenv("HTTP_PORT"))
}

func (a *App) initialize() {
	var err error

	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	maxOpenConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	fmt.Println(os.Getenv("CONSOLE_L"))

	a.DB, err = db.ConnectDB(driver, host, database, user, password, port, maxOpenConnections)

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) run(addr string) {
	fmt.Printf("Server started at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) routes() {
	controller := controller.Controller{}
	a.Router.HandleFunc("/api/v1/users", controller.GetUsers(a.DB)).Methods("GET")
	a.Router.HandleFunc("/api/v1/users", controller.AddUser(a.DB)).Methods("POST")
	a.Router.HandleFunc("/api/v1/users/{id}", controller.UpdateUser(a.DB)).Methods("PATCH")
	a.Router.HandleFunc("/api/v1/users/{id}", controller.RemoveUser(a.DB)).Methods("DELETE")

	a.Router.HandleFunc("/api/v1/login", controller.LoginControl(a.DB)).Methods("POST")

	a.Router.HandleFunc("/api/v1/games", controller.GetGames(a.DB)).Methods("GET")
	a.Router.HandleFunc("/api/v1/games", controller.AddGame(a.DB)).Methods("POST")

	a.Router.HandleFunc("/api/v1/usergames/{id}", controller.GetUserGames(a.DB)).Methods("GET")
}
