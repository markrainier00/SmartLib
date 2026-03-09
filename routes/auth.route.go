package routes

import (
	"SmartLib_Likod/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/signin", handler.Signin)
}
