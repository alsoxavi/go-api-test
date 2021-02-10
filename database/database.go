package database

import (
	"fmt"
	"strconv"

	"github.com/alsoxavi/go-api-test/config"
	"github.com/alsoxavi/go-api-test/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConn is global for database access
var DBConn *gorm.DB

// InitDatabase starts database connection
func InitDatabase() error {
	var err error
	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
		return err
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s DB.name=%s port=%d sslmode=disable TimeZone=UTC", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database")
		return err
	}

	DBConn.AutoMigrate(&model.Product{})

	return nil
}
