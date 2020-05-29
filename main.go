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

	// Public API: for all users
	publicAPI := app.Group("/api/public")

	// GET
	publicAPI.Get("/users", controllers.ShowUsers)
	publicAPI.Get("/user/:username", controllers.ShowUserByUsername)
	// publicAPI.Get("/project/:alias", controllers.ShowProjectByAlias)
	// publicAPI.Get("/projects", controllers.ShowProjects)

	// POST
	publicAPI.Post("/auth", controllers.Authentication)
	publicAPI.Post("/register", controllers.CreateUser)
	publicAPI.Post("/forget-password", controllers.ForgetPasswordIssue)

	// DELETE
	publicAPI.Delete("/forget-password", controllers.ForgetPasswordCheckResetCode)

	// Private API: only for JWT authenticated users
	privateAPI := app.Group("/api/private", jwtware.New(jwtwareConfig))

	// POST
	privateAPI.Post("/refresh-token", controllers.RefreshToken)
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

	// Error 404 Not Found
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
