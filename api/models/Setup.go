package models

import (
	"fmt"
	"os"

	// production
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() *gorm.DB {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	databaseName := os.Getenv("DB_NAME")
	// production
	// postgres
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, user, databaseName, password)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        dsn,
	}))

	// dev
	// postgres
	// dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=5432 sslmode=disable password=%s", dbHost, user, databaseName, password)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Medicine{})

	return DB
}
