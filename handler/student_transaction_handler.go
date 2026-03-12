package handler

import (
	"SmartLib_Likod/repositories"
	"SmartLib_Likod/services"

	"github.com/gofiber/fiber/v2"
)

func BorrowBook(c *fiber.Ctx) error {
	var input services.BorrowInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input format"})
	}

	if err := services.BorrowBookService(input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error(), "isSuccess": false})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Request sent to Staff!", "isSuccess": true})
}

func GetAllPending(c *fiber.Ctx) error {
	requests, err := repositories.GetAllPendingRequests()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Database error", "isSuccess": false})
	}
	return c.JSON(fiber.Map{"isSuccess": true, "data": requests})
}

func GetDashboardStats(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"isSuccess": true,
		"data": fiber.Map{
			"pendingRegistrations": repositories.GetPendingRegCount(),
			"borrowRequests":       repositories.GetPendingBorrowCount(),
			"activeBorrows":        repositories.GetActiveBorrowCount(),
		},
	})
}

func GetStudentHistory(c *fiber.Ctx) error {
	schoolID := c.Query("school_id")
	if schoolID == "" {
		return c.Status(400).JSON(fiber.Map{"message": "School ID required", "isSuccess": false})
	}
	history, err := repositories.GetTransactionHistory(schoolID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "History fetch failed", "isSuccess": false})
	}
	return c.JSON(fiber.Map{"isSuccess": true, "data": history})
}

func ReleaseBook(c *fiber.Ctx) error {
	type Req struct {
		SchoolID string `json:"school_id"`
	}
	var body Req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request", "isSuccess": false})
	}
	if err := services.ReleaseBookService(body.SchoolID); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error(), "isSuccess": false})
	}
	return c.JSON(fiber.Map{"message": "Book released!", "isSuccess": true})
}
