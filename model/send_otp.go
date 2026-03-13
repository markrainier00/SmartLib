package model

import "time"

type OTPCode struct {
	ID        uint      `json:"iid" gorm:"primarykey"`
	Email     string    `json:"email" gorm:"not null"`
	OTP       string    `json:"otp" gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `json:"expires_at"`
	Used      bool      `json:"used" gorm:"default:false"`
}
