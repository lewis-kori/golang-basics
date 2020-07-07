package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// import of the postgress orm driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Init initializes the database connection
func Init() *gorm.DB {

	host := os.Getenv("GO_DB_HOST")
	// port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port := 5432

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err.Error())
	}
	return db
}
