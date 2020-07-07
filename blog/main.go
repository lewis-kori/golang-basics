package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lewis-kori/todoapp/controllers"
	"github.com/lewis-kori/todoapp/database"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	database.SeedToDo()

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.TodoList)
	router.HandleFunc("/add", controllers.TodoAdd)
	router.HandleFunc("/delete/{id}", controllers.TodoDelete)

	fmt.Println("connecting...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
