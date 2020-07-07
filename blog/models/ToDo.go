package models

import (
	"github.com/jinzhu/gorm"
)

// Todo represents a todo instance
type Todo struct {
	gorm.Model
	Title     string
	Completed bool
}
