package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User Struct
type User struct {
	gorm.Model
	Name  string
	Email string
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// migrate the schema
	db.AutoMigrate(&User{})
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", allUsers)
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	fmt.Println("Go ORM Tutorial")

	// Add the call to our new initialMigration function
	initialMigration()
	handleRequests()

}
