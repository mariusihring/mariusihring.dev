package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mariusihring.dev/services/db/databasemodels"
	"os"
	"time"
)



func NewDbConnection() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to db")
	}

	db.AutoMigrate(&databasemodels.Goal{}, &databasemodels.Entry{}, &databasemodels.User{}, &databasemodels.Session{})
	return db ,nil

}

