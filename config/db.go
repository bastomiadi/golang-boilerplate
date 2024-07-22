package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Connect initializes the database connection
func Connect() {
	once.Do(func() {
		// Load environment variables from .env file
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		var err error
		dsn := os.Getenv("DB_DSN")
		driver := os.Getenv("DB_DRIVER")

		if driver == "" {
			driver = "mysql" // default to MySQL if not specified
		}

		// Initialize the database connection
		switch driver {
		case "postgres":
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		case "mysql":
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		default:
			log.Fatalf("Unsupported database driver: %s", driver)
		}

		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}

		// Optional: Test the connection
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Could not get raw database object: %v", err)
		}

		if err = sqlDB.Ping(); err != nil {
			log.Fatalf("Could not ping the database: %v", err)
		}

		log.Println("Connected to the database")
	})
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
