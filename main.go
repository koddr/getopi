package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/koddr/getopi/controllers"
)

func main() {
	// Fiber app
	app := fiber.New()

	// Settings
	loggerConfig := logger.Config{
		Format:     "${time} [${status}] ${method} ${path}\n",
		TimeFormat: "Mon, 2 Jan 2006 15:04:05 MST",
	}

	// Logger
	app.Use(logger.New(loggerConfig))

	// GET
	app.Get("/user/:uuid", controllers.UserController)
	app.Get("/users", controllers.UsersController)

	// POST
	app.Post("/user", controllers.UserCreateController)

	// PATCH
	// app.Patch("/user/:uuid", controllers.UserUpdateController)

	// 404 Not Found
	app.Use(func(c *fiber.Ctx) {
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "not found"})
	})

	app.Listen(":3000")
}
