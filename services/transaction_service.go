package services

import (
	"SmartLib_Likod/model"
	"SmartLib_Likod/repositories"
	"errors"
)

type BorrowInput struct {
	SchoolID  string `json:"school_id"`
	BookTitle string `json:"book_title"`
}

func BorrowBookService(input BorrowInput) error {
	// Check if student already has a pending request
	if repositories.HasPendingTransaction(input.SchoolID) {
		return errors.New("you already have a pending book request")
	}

	tx := &model.Transaction{
		SchoolID:  input.SchoolID,
		BookTitle: input.BookTitle,
		Status:    "Pending",
	}

	return repositories.CreateTransaction(tx)
}

func ReleaseBookService(schoolID string) error {
	return repositories.ReleaseBookStatus(schoolID)
}
