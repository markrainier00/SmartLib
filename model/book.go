package model

import (
	"time"
)

type Book struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Author      string    `gorm:"type:varchar(255);not null" json:"author"`
	Category    string    `gorm:"type:varchar(100)" json:"category"`
	Course      string    `gorm:"type:varchar(50)" json:"course"`
	Available   bool      `gorm:"default:true" json:"available"`
	Pages       int       `json:"pages"`
	Copies      int       `gorm:"default:1" json:"copies"`
	Description string    `gorm:"type:text" json:"description"`
	ActualImage string    `gorm:"type:text" json:"actualImage"`
	Cover       string    `gorm:"type:varchar(50);default:'#1a1a2e'" json:"cover"`
	Accent      string    `gorm:"type:varchar(50);default:'#c9a84c'" json:"accent"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
