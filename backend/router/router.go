package router

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	api := app.Group("/api")

	// Guest
	product := api.Group("/product")
	product.Get("/")
	product.Get("/:id")

	//
	auth := api.Group("/auth")
	auth.Post("/login")
	auth.Post("/register")
	auth.Post("/logout")
}
