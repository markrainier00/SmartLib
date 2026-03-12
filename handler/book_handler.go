package handler

import (
	"SmartLib_Likod/database"
	"SmartLib_Likod/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetAllBooks - Kinukuha ang mga libro sa Supabase
func GetAllBooks(c *fiber.Ctx) error {
	var books []model.Book

	// 🚀 Kukunin sa database
	if err := database.DB.Find(&books).Error; err != nil {
		fmt.Println("🚨 Error sa pagkuha ng books:", err)
		return c.Status(500).JSON(fiber.Map{
			"isSuccess": false,
			"message":   "Failed to fetch books",
		})
	}

	return c.JSON(fiber.Map{
		"isSuccess": true,
		"data":      books,
	})
}

// AddBook - Nagse-save ng bagong libro mula sa Admin papunta sa Supabase
func AddBook(c *fiber.Ctx) error {
	book := new(model.Book)

	// 1. Basahin ang pinadala ng Frontend
	if err := c.BodyParser(book); err != nil {
		fmt.Println("🚨 Error sa Body Parser:", err)
		return c.Status(400).JSON(fiber.Map{
			"isSuccess": false,
			"message":   "Invalid input data",
		})
	}

	// 👁️ CCTV: I-print sa terminal kung ano yung natanggap mula sa Next.js
	fmt.Printf("📦 TANGKANG I-SAVE NA LIBRO: %+v\n", book)

	// 2. 🚀 I-SAVE SA SUPABASE
	result := database.DB.Create(&book)
	if result.Error != nil {
		// Kapag nag-error ang Supabase, i-print sa terminal ang dahilan!
		fmt.Println("🚨 SUPABASE SAVE ERROR:", result.Error)
		return c.Status(500).JSON(fiber.Map{
			"isSuccess": false,
			"message":   "Failed to save to database",
			"error":     result.Error.Error(),
		})
	}

	// 3. Success!
	fmt.Println("✅ SUCCESS! PUMASOK SA SUPABASE ANG LIBRO. ID:", book.ID)

	return c.JSON(fiber.Map{
		"isSuccess": true,
		"message":   "Book added successfully",
		"data":      book,
	})
}

// ==========================================
// 🚀 BAGONG DAGDAG: DELETE BOOK FUNCTION
// ==========================================

// DeleteBook - Binubura ang libro sa Supabase gamit ang ID
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id") // Kukunin ang ID mula sa URL (halimbawa: /api/books/1)

	// 👁️ CCTV: Tingnan natin kung anong ID ang gustong burahin
	fmt.Println("🗑️ TANGKANG BURAHIN ANG LIBRO. ID:", id)

	// Uutusan ang GORM na burahin ang record sa database
	result := database.DB.Delete(&model.Book{}, id)

	if result.Error != nil {
		fmt.Println("🚨 SUPABASE DELETE ERROR:", result.Error)
		return c.Status(500).JSON(fiber.Map{
			"isSuccess": false,
			"message":   "Failed to delete book",
		})
	}

	// Success!
	fmt.Println("✅ SUCCESS! NABURA NA ANG LIBRO SA SUPABASE. ID:", id)

	return c.JSON(fiber.Map{
		"isSuccess": true,
		"message":   "Book successfully deleted!",
	})
}
