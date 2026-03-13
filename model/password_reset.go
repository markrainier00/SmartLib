package model

import "time"

type PasswordReset struct {
	ID        uint      `json:"iid" gorm:"primarykey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Token     string    `json:"token" gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `json:"expires_at"`
	Used      bool      `json:"used" gorm:"default:false"`
}
