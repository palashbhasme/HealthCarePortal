package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	Clerk Role = "receptionist"
	Doc   Role = "doctor"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      Role      `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(User{}, Patient{})
}
