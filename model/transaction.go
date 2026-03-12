package model

import "gorm.io/gorm"

type Transaction struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	SchoolID  string         `gorm:"column:school_id" json:"school_id"`
	BookTitle string         `gorm:"column:book_title" json:"book_title"`
	Status    string         `gorm:"column:status" json:"status"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// Pinipilit natin ang GORM na gamitin ang "transactions" table
func (Transaction) TableName() string {
	return "transactions"
}
