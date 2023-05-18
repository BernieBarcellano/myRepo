// set up ng pagcoconnect kay DB
package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Ejil/studen_database/config"
	"github.com/Ejil/studen_database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatalf("Failed to parse DB_PORT: %v", err)
	}

	//Connection URL to connect to Postgres DB
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	//Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Print("Successfully connected to database")

	//Migrate the database
	DB.AutoMigrate(&models.SignUp{})
	fmt.Println("Database Migrated")
}
