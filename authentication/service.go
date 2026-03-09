package authentication

import (
	"SmartLib_Likod/model"
	"SmartLib_Likod/model/status"
	"SmartLib_Likod/utils"
	"errors"
)

type RegisterInput struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	SchoolID  string `json:"school_id"`
	Program   string `json:"program"`
	Year      string `json:"year"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

type SigninInput struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

func RegisterUser(input RegisterInput) (*model.User, error) {
	existing, err := FindUserByEmailOrSchoolID(input.Email, input.SchoolID)
	if err == nil && existing.ID != 0 {
		return nil, errors.New("email or school ID already registered")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, errors.New("failed to process password")
	}

	user := &model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		SchoolID:  input.SchoolID,
		Program:   input.Program,
		Year:      input.Year,
		Status:    status.UserStatusNew,
		Password:  hashedPassword,
	}

	if err := CreateUser(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func SigninUser(input SigninInput) (*model.User, error) {
	user, err := FindUserByEmailOrSchoolID(input.Identifier, input.Identifier)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return nil, errors.New("Invalid credentials")
	}

	if user.Status == status.UserStatusNew {
		return nil, errors.New("Your account is not yet approved by the admin.")
	} else if user.Status == status.UserStatusLocked {
		return nil, errors.New("Your account has been locked, please contact the admin.")
	} else if user.Status != status.UserStatusActive {
		return nil, errors.New("Your account status is invalid, please contact the admin.")
	}

	return user, nil
}
