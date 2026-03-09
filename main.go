package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"SmartLib_Likod/database"
	"SmartLib_Likod/handler"
	"SmartLib_Likod/middleware"
	"SmartLib_Likod/model"
	"SmartLib_Likod/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	database.DB.AutoMigrate(&model.User{})

	app := fiber.New()

	middleware.SetupCORS(app)

	routes.Setup(app)

	// just a health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "API is running"})
	})

	_ = handler.Register

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
