package db

import (
	"log"
	"os"
	"time"

	"github.com/jasonDeichert/GoGrocery/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() *gorm.DB {
	// Setting up GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags), // log to standard output
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate all the schemas
	if err := DB.AutoMigrate(&model.Recipe{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return DB
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to close database connection: %v", err)
	}
	db.Close()
}
