package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `json:"id" gorm:"primarykey"`
	FirstName     string         `json:"firstname" gorm:"not null"`
	LastName      string         `json:"lastname" gorm:"not null"`
	Email         string         `json:"email" gorm:"uniqueIndex;not null"`
	SchoolID      string         `json:"school_id" gorm:"uniqueIndex;not null"`
	Program       string         `json:"program" gorm:"not null"`
	Year          string         `json:"year" gorm:"not null"`
	Status        string         `json:"status" gorm:"default:New"`
	Password      string         `json:"-" gorm:"not null"`
	SchoolIDImage string         `json:" school_id_image" gorm:"column:school_id_image"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
