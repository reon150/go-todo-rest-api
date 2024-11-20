package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/reon150/go-todo-rest-api/migrations"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func InitDatabase() {
	ensureDatabaseExists(AppConfig.SQLiteDBPath)

	var err error
	DB, err = gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        AppConfig.SQLiteDBPath,
	}, &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := migrations.RunMigrations(DB); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database connected successfully")
}

func ensureDatabaseExists(dbPath string) {
	const defaultFilePermissions = 0755
	dir := filepath.Dir(dbPath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, defaultFilePermissions)
		if err != nil {
			log.Fatalf("Failed to create directory for database: %v", err)
		}
		log.Printf("Created directory for database: %s", dir)
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatalf("Failed to create database file: %v", err)
		}
		file.Close()
		log.Printf("Created SQLite database file: %s", dbPath)
	}
}
