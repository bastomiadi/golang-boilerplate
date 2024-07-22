package config

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
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

		db, err = sql.Open(driver, dsn)
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Could not ping the database: %v", err)
		}

		log.Println("Connected to the database")
	})
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
