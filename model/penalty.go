package model

import "gorm.io/gorm"

type Penalty struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	SchoolID  string         `json:"school_id" gorm:"index"`
	Amount    float64        `json:"amount"`
	IsPaid    bool           `json:"is_paid" gorm:"default:false"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
