package routes

import (
	"SmartLib_Likod/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Base Group for API
	api := app.Group("/api")

	// 🔐 Authentication
	// Tatawagin nito ang SetupAuthRoutes mula sa auth_route.go
	SetupAuthRoutes(api)

	// 💸 Transactions Group
	// Handler: student_transaction_handler.go
	transactions := api.Group("/transactions")
	transactions.Post("/borrow", handler.BorrowBook)
	transactions.Get("/history", handler.GetStudentHistory)
	transactions.Get("/pending-all", handler.GetAllPending)
	transactions.Put("/release", handler.ReleaseBook)

	// 📊 Dashboard Stats
	// Handler: student_transaction_handler.go
	api.Get("/admin/stats", handler.GetDashboardStats)

	// 📚 Books Group
	// Handler: book_handler.go
	books := api.Group("/books")
	books.Get("/", handler.GetAllBooks) // GET /api/books/
	books.Post("/", handler.AddBook)    // POST /api/books/

	// ==========================================
	// 🚀 BAGONG DAGDAG: DELETE ROUTE
	// ==========================================
	books.Delete("/:id", handler.DeleteBook) // DELETE /api/books/:id

	// 🔍 Smart Scanner
	// Handler: admin_scanner_handler.go
	scanner := api.Group("/scanner")
	scanner.Get("/:school_id", handler.GetStudentScannerData) // GET /api/scanner/:school_id
}
