package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lewis-kori/todoapp/database"
	"github.com/lewis-kori/todoapp/models"
)

// TodoList retrieves all the todo items
func TodoList(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("views/list.html")
	if err != nil {
		log.Fatal("Template not found ")
	}

	db := database.Init()
	var todos []models.Todo

	query := db.Find(&todos)

	defer query.Close()

	view.Execute(w, todos)
}

// TodoAdd creates a new record on the db
func TodoAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		todoForm(w, r)
	case "POST":
		processTodoForm(w, r)
	}
}

func todoForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/add.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

func processTodoForm(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	if title == "" {
		http.Redirect(w, r, "/", 302)
	}

	db := database.Init()

	todo := models.Todo{
		Title:     title,
		Completed: false,
	}

	db.Create(&todo)
	http.Redirect(w, r, "/", 302)
}

// TodoDelete erases a to do item from the database
func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Redirect(w, r, "/", 302)
	}

	db := database.Init()
	var todos []models.Todo

	db.Where("id = ?", id).Delete(todos)
	http.Redirect(w, r, "/", 302)
}
