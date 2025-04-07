package db

import (
	"fmt"
	"log"
	"observer-go/src/structs/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
	log  *log.Logger
}

var models = []interface{}{
	&model.Company{},
}
var dependableModels = []interface{}{
	&model.User{},
	&model.Client{},
	&model.Registration{},
}

func Init(logger *log.Logger) *Database {
	return &Database{
		log: logger,
	}
}

func (i *Database) Connect() {
	// Get environment variables
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	app_env := os.Getenv("APP_ENV")
	sslmode := "disable"
	timeZone := "UTC"
	if app_env == "test" {
		dbName = os.Getenv("POSTGRES_DB_TEST")
	} else if app_env != "production" && app_env != "dev" {
		i.log.Fatalf("Invalid APP_ENV: %s", app_env)
	}

	i.log.Printf("Running in %s environment. Database: %s\n", app_env, dbName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbName, port, sslmode, timeZone)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		i.log.Fatal("Failed to connect to the database: ", err)
	}

	i.Gorm = db
}
func (db *Database) Migrate() {
	db.Clear()
	for _, model := range models {
		log.Printf("Migrating: %T", model)
		if err := db.Gorm.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate %T: %v", model, err)
		}
	}
	for _, model := range dependableModels {
		log.Printf("Migrating: %T", model)
		if err := db.Gorm.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to migrate %T: %v", model, err)
		}
	}
	log.Println("Migration completed successfully")
}

func (db *Database) Disconnect() {
	sqlDB, err := db.Gorm.DB()
	if err != nil {
		db.log.Fatal("Failed to close the database: ", err)
	}
	sqlDB.Close()
}

// Clear the database. Only in test environment
func (db *Database) Clear() {
	if os.Getenv("APP_ENV") != "test" || os.Getenv("APP_ENV") != "dev" {
		return
	}
	query := `
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
	`
	if err := db.Gorm.Exec(query).Error; err != nil {
		db.log.Fatalf("Failed to clear database: %v", err)
	}
	db.log.Println("Database cleared successfully")
}
