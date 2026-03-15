package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	SchoolID  string `gorm:"column:school_id" json:"school_id"`
	BookTitle string `gorm:"column:book_title" json:"book_title"`
	Status    string `gorm:"column:status" json:"status"`

	PickupDate string `gorm:"column:pickup_date" json:"pickup_date"`
	ReturnDate string `gorm:"column:return_date" json:"return_date"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`

	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}
