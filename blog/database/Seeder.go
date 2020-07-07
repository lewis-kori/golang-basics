package database

import "github.com/lewis-kori/todoapp/models"

// SeedToDo creates the todo table
func SeedToDo() {

	db := Init()

	db.DropTable(&models.Todo{})
	db.AutoMigrate(&models.Todo{})

	todo := models.Todo{
		Title:     "Creating todo app with golang",
		Completed: false,
	}

	db.Create(&todo)
}
