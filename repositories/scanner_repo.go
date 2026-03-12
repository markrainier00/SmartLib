package repositories

import (
	"SmartLib_Likod/database"
	"SmartLib_Likod/model"
)

func GetPendingTransaction(schoolID string) (*model.Transaction, error) {
	var transaction model.Transaction
	err := database.DB.Where("school_id = ? AND status = ?", schoolID, "Pending").First(&transaction).Error
	return &transaction, err
}

func GetUnpaidPenalty(schoolID string) (float64, error) {
	var penalties []model.Penalty
	err := database.DB.Where("school_id = ? AND is_paid = ?", schoolID, false).Find(&penalties).Error
	if err != nil {
		return 0, err
	}
	total := 0.0
	for _, p := range penalties {
		total += p.Amount
	}
	return total, nil
}
