package databasemodels

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id uint `gorm:"primaryKey autoincrement"`
	LastName *string `gorm:"size:255"`
	FirstName *string `gorm:"size:255"`
	UserName string `gorm:"size:255"`
	Email string `gorm:"size:255 index:,unique"`
	PasswordHash string `gorm:"size:255"`
	Session []Session
	Goals []Goal
}

type Session struct {
	Id string `gorm:"primaryKey autoincrement"`
	ExpiresAt int64 `gorm:"not null"`
	UserId uint `gorm:"not null"`
}