package main

import (
	"github.com/gofiber/compression"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	jwtware "github.com/gofiber/jwt"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"github.com/google/uuid"
	"github.com/koddr/getopi/controllers"
	"github.com/koddr/getopi/utils"
)

func main() {
	// Fiber app
	app := fiber.New()

	app.Get("/token", func(c *fiber.Ctx) {
		// Convert string to UUID
		uuid, _ := uuid.Parse("45cb822b-65f4-410c-9d9d-0ccd0a8ba73e")
		token, _ := utils.GenerateJWT("admin", uuid)
		c.JSON(fiber.Map{"token": token})
	})

	// Middlewares
	app.Use(
		cors.New(),
		helmet.New(),
		compression.New(),
		logger.New(logger.Config{
			Format:     "${time} [${status}] ${method} ${path} (${latency})\n",
			TimeFormat: "Mon, 2 Jan 2006 15:04:05 MST",
		}),
		recover.New(recover.Config{
			Handler: func(c *fiber.Ctx, err error) {
				c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
			},
		}),
		jwtware.New(jwtware.Config{
			SigningKey: []byte(utils.GetDotEnvValue("JWT_SECRET_TOKEN")),
			ErrorHandler: func(c *fiber.Ctx, err error) {
				c.Status(403).JSON(fiber.Map{"error": true, "msg": err.Error()})
			},
		}),
	)

	// API v1 group
	v1 := app.Group("/api/v1")

	// GET
	v1.Get("/user/:username", controllers.UserController)
	v1.Get("/users", controllers.UsersController)

	// POST
	v1.Post("/user", controllers.UserCreateController)

	// PATCH
	v1.Patch("/user", controllers.UserUpdateController)

	// DELETE
	v1.Delete("/user", controllers.UserDeleteController)

	// 404 Not Found
	app.Use(func(c *fiber.Ctx) {
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "endpoint not found"})
	})

	// Run server
	app.Listen(utils.GetDotEnvValue("SERVER_PORT"))
}
