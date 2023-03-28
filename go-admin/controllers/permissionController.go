package controllers

import (
	"go-admin/database"
	"go-admin/middlewares"
	"go-admin/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "permissions"); err != nil {
		return err
	}

	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
