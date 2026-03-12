package model

import (
	"time"
)

// Book - Ito ang structure ng table mo sa Supabase
type Book struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Author      string    `gorm:"type:varchar(255);not null" json:"author"`
	Category    string    `gorm:"type:varchar(100)" json:"category"`
	Course      string    `gorm:"type:varchar(50)" json:"course"` // Hal. "BSCS", "BSIT", o "All"
	Available   bool      `gorm:"default:true" json:"available"`
	Pages       int       `json:"pages"`
	Copies      int       `gorm:"default:1" json:"copies"`
	Description string    `gorm:"type:text" json:"description"`
	ActualImage string    `gorm:"type:text" json:"actualImage"`                     // Dito ise-save ang image URL o Base64
	Cover       string    `gorm:"type:varchar(50);default:'#1a1a2e'" json:"cover"`  // Kulay ng cover design
	Accent      string    `gorm:"type:varchar(50);default:'#c9a84c'" json:"accent"` // Kulay ng accent
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
