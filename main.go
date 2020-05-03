package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/compression"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	jwtware "github.com/gofiber/jwt"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"github.com/koddr/getopi/controllers"
)

func main() {
	// Fiber app
	app := fiber.New()

	app.Get("/token", func(c *fiber.Ctx) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "John Doe"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
			return
		}

		c.JSON(fiber.Map{"token": t})
	})

	// Middlewares
	app.Use(
		cors.New(),
		helmet.New(),
		compression.New(),
		logger.New(logger.Config{
			Format:     "${time} [${status}] ${method} ${path}\n",
			TimeFormat: "Mon, 2 Jan 2006 15:04:05 MST",
		}),
		recover.New(recover.Config{
			Handler: func(c *fiber.Ctx, err error) {
				c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
			},
		}),
		jwtware.New(jwtware.Config{
			SigningKey: []byte("secret"),
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

	app.Listen(":3000")
}
