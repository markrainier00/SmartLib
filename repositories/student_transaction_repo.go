package repositories

import (
	db "SmartLib_Likod/database"
	"SmartLib_Likod/model"
	"fmt"
	"time"
)

// HasActiveBorrow - Check kung may active borrow ang student (hindi na naka-block ang pending)
func HasActiveBorrow(schoolID string) bool {
	var count int64
	db.DB.Model(&model.Transaction{}).
		Where("school_id = ? AND status = ?", schoolID, "Borrowed").
		Count(&count)
	return count > 0
}

// HasPendingTransaction - Chine-check kung may active request pa ang student
func HasPendingTransaction(schoolID string) bool {
	var count int64
	// Ginamit natin ang 'db.DB' imbes na 'database.DB'
	db.DB.Model(&model.Transaction{}).Where("school_id = ? AND status = ?", schoolID, "Pending").Count(&count)
	return count > 0
}

// CreateTransaction - Nagse-save ng bagong borrow request
func CreateTransaction(tx *model.Transaction) error {
	result := db.DB.Debug().Create(tx)
	if result.Error != nil {
		fmt.Println("🚨 SUPABASE INSERT ERROR:", result.Error)
		return result.Error
	}
	return nil
}

// GetTransactionHistory - Para sa "Borrow History" ng student
func GetTransactionHistory(schoolID string) ([]model.Transaction, error) {
	var history []model.Transaction
	err := db.DB.Where("school_id = ?", schoolID).Order("id desc").Find(&history).Error
	return history, err
}

// GetAllPendingRequests - Para sa Admin "Pending Approvals" list
func GetAllPendingRequests() ([]model.Transaction, error) {
	var requests []model.Transaction
	err := db.DB.Where("status = ?", "Pending").Order("id desc").Find(&requests).Error
	return requests, err
}

// ReleaseBookStatus - Update status from 'Pending' to 'Borrowed' (Scanner Action)
// ✅ FIX: I-update din ang updated_at para lumabas ang Date Approved sa frontend
func ReleaseBookStatus(schoolID string) error {
	return db.DB.Model(&model.Transaction{}).
		Where("school_id = ? AND status = ?", schoolID, "Pending").
		Updates(map[string]interface{}{
			"status":     "Borrowed",
			"updated_at": time.Now(),
		}).Error
}

func GetAllRequests() ([]model.Transaction, error) {
	var requests []model.Transaction
	err := db.DB.Order("id desc").Find(&requests).Error
	return requests, err
}

func UpdateTransactionStatus(schoolID string, oldStatus string, newStatus string) error {
	return db.DB.Model(&model.Transaction{}).
		Where("school_id = ? AND status = ?", schoolID, oldStatus).
		Updates(map[string]interface{}{
			"status":     newStatus,
			"updated_at": time.Now(),
		}).Error
}

// --- DASHBOARD STATS QUERIES ---

func GetPendingRegCount() int64 {
	var count int64
	db.DB.Model(&model.User{}).Where("status = ?", "New").Count(&count)
	return count
}

func GetPendingBorrowCount() int64 {
	var count int64
	db.DB.Model(&model.Transaction{}).Where("status = ?", "Pending").Count(&count)
	return count
}

func GetActiveBorrowCount() int64 {
	var count int64
	db.DB.Model(&model.Transaction{}).Where("status = ?", "Borrowed").Count(&count)
	return count
}
