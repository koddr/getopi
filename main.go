package main

import (
	"github.com/gofiber/compression"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	jwtware "github.com/gofiber/jwt"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"github.com/koddr/getopi/controllers"
	"github.com/koddr/getopi/utils"
)

func main() {
	// Fiber app
	app := fiber.New()

	// Configs
	jwtwareConfig := jwtware.Config{
		SigningKey: []byte(utils.GetDotEnvValue("JWT_SECRET_TOKEN")),
		ErrorHandler: func(c *fiber.Ctx, err error) {
			c.Status(403).JSON(fiber.Map{"error": true, "msg": err.Error()})
		},
	}
	loggerConfig := logger.Config{
		Format:     "${time} [${status}] ${method} ${path} (${latency})\n",
		TimeFormat: "Mon, 2 Jan 2006 15:04:05 MST",
	}
	recoverConfig := recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			c.Status(500).JSON(fiber.Map{"error": true, "msg": err.Error()})
		},
	}

	// Middlewares
	app.Use(
		cors.New(),
		helmet.New(),
		compression.New(),
		recover.New(recoverConfig),
		logger.New(loggerConfig),
	)

	// Public API
	app.Post("/auth", controllers.Authentication)
	app.Post("/register", controllers.CreateUser)

	// Private API group
	privateAPI := app.Group("/api", jwtware.New(jwtwareConfig))

	// GET
	privateAPI.Get("/user/:username", controllers.ShowUserByUsername)
	privateAPI.Get("/users", controllers.ShowUsers)
	// privateAPI.Get("/project/:alias", controllers.ShowProjectByAlias)
	// privateAPI.Get("/projects", controllers.ShowProjects)

	// POST
	privateAPI.Post("/auth/refresh-token", controllers.RefreshToken)
	// privateAPI.Post("/project", controllers.CreateProject)
	// privateAPI.Post("/task", controllers.CreateTask)

	// PATCH
	privateAPI.Patch("/user", controllers.UpdateUser)
	// privateAPI.Patch("/project", controllers.UpdateProject)
	// privateAPI.Patch("/task", controllers.UpdateTask)

	// DELETE
	privateAPI.Delete("/user", controllers.DeleteUser)
	// privateAPI.Delete("/project", controllers.DeleteProject)
	// privateAPI.Delete("/task", controllers.DeleteTask)

	// 404 Not Found
	app.Use(func(c *fiber.Ctx) {
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "sorry, endpoint not found"})
	})

	// AutoCert
	// tls := utils.AutoCertSSLGenerator(
	// 	true, // TODO: Dry-run Let's Encrypt is now enabled
	// 	utils.GetDotEnvValue("EMAIL"),
	// 	[]string{
	// 		utils.GetDotEnvValue("DOMAIN_WITH_WWW"),
	// 		utils.GetDotEnvValue("DOMAIN_WITHOUT_WWW"),
	// 	},
	// )

	// Run server
	app.Listen(utils.GetDotEnvValue("SERVER_PORT"))
}
