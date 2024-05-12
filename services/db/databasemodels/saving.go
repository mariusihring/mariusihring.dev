package databasemodels

import (
	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	Name string
	Amount float64
	Description *string
	Link *string
	Image *string
	Entries []Entry
	UserId uint
}


type Entry struct {
	Id uint `gorm:"primaryKey;autoIncrement"`
	Amount float64
	Comment *string
	Name *string
	GoalId uint
}