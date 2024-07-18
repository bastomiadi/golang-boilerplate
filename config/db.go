package config

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

// Connect initializes the database connection
func Connect() {
	once.Do(func() {
		var err error
		dsn := "root:@/golang-boilerplate"
		db, err = sql.Open("mysql", dsn)
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
