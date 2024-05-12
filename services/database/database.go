package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mariusihring.dev/services/database/databasemodels"
)



func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&databasemodels.Goal{}, &databasemodels.Entry{})

}