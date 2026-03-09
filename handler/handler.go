package handler

import (
	"SmartLib_Likod/authentication"
	errormodel "SmartLib_Likod/model/error"
	"SmartLib_Likod/model/response"
	"SmartLib_Likod/model/status"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var input authentication.RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errormodel.ErrorModel{
			Message:   status.RetCode404,
			IsSuccess: false,
			Error:     err,
		})
	}

	if input.FirstName == "" || input.LastName == "" || input.Email == "" || input.SchoolID == "" || input.Program == "" || input.Year == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(errormodel.ErrorModel{
			Message:   status.RetCode401,
			IsSuccess: false,
			Error:     nil,
		})
	}

	if len(input.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(errormodel.ErrorModel{
			Message:   "Password must be at least 8 characters",
			IsSuccess: false,
			Error:     nil,
		})
	}

	user, err := authentication.RegisterUser(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errormodel.ErrorModel{
			Message:   err.Error(),
			IsSuccess: false,
			Error:     err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.ResponseModel{
		RetCode: "201",
		Message: "Registration successful",
		Data:    user,
	})
}

func Signin(c *fiber.Ctx) error {
	var input authentication.SigninInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errormodel.ErrorModel{
			Message:   status.RetCode404,
			IsSuccess: false,
			Error:     err,
		})
	}

	if input.Identifier == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(errormodel.ErrorModel{
			Message:   status.RetCode401,
			IsSuccess: false,
			Error:     nil,
		})
	}

	user, err := authentication.SigninUser(input)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(errormodel.ErrorModel{
			Message:   err.Error(),
			IsSuccess: false,
			Error:     err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.ResponseModel{
		RetCode: "200",
		Message: "Sign in successfull",
		Data:    user,
	})
}
