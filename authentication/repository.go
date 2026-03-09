package authentication

import (
	"SmartLib_Likod/database"
	"SmartLib_Likod/model"
)

func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

func FindUserByEmailOrSchoolID(email, school_id string) (*model.User, error) {
	var user model.User
	result := database.DB.Where("email = ? OR school_id = ?", email, school_id).First(&user)
	return &user, result.Error
}
