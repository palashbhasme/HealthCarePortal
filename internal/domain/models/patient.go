package models

import (
	"time"
)

type Patient struct {
	ID             uint      `gorm:"primaryKey"`
	FirstName      string    `gorm:"type:varchar(255);not null"`
	LastName       string    `gorm:"type:varchar(255);not null"`
	DOB            time.Time `gorm:"not null"`
	Email          *string   `gorm:"unique"`
	Gender         string    `gorm:"type:varchar(255);not null"`
	PhoneNumber    string    `gorm:"not null"`
	Address        string    `gorm:"not null"`
	MedicalHistory string    `gorm:"type:text; not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
