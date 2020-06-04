package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/koddr/getopi/stores"
)

// ShowProjectByAlias ...
// show exists project by alias or 404 error
func ShowProjectByAlias(c *fiber.Ctx) {
	// Get alias from URL
	alias := c.Params("alias")

	// Create DB connection
	db, errConnectDB := stores.OpenStore()
	if errConnectDB != nil {
		// DB connection error
		c.Status(500).JSON(fiber.Map{"error": true, "msg": errConnectDB.Error()})
		return
	}

	// Find project by alias
	project, errFindProjectByAlias := db.FindProjectByAlias(alias)
	if errFindProjectByAlias != nil {
		// Project not found
		c.Status(404).JSON(fiber.Map{"error": true, "msg": "project not found", "project": nil})
		return
	}

	// Find project author by ID
	user, errFindUserByID := db.FindUserByID(project.AuthorID)
	if errFindUserByID != nil {
		// User not found
		c.Status(404).JSON(fiber.Map{
			"error": true, "msg": "project author not found", "project": nil, "author": nil,
		})
		return
	}

	// Rebuild author data
	author := fiber.Map{
		"username":   user.Username,
		"picture":    user.UserAttrs.Picture,
		"first_name": user.UserAttrs.FirstName,
		"last_name":  user.UserAttrs.LastName,
	}

	c.JSON(fiber.Map{"error": false, "msg": nil, "project": project, "author": author})
}
