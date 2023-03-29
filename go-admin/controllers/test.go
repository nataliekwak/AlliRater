package controllers

import (
	"fmt"
	"go-admin/database"
	"go-admin/middlewares"
	"go-admin/models"

	"github.com/gofiber/fiber/v2"
)

func RegisterTest(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	if user.FirstName != "" && user.LastName != "" && user.Email != "" {
		fmt.Println("User has been added successfully")
	} else {
		fmt.Println("User has NOT been added successfully")
	}

	return nil
}

func CreateUserTest(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	if user != nil {
		fmt.Println("Did not create user")
	} else {
		fmt.Println("Successfully created user")
	}

	return c.JSON(user)
}
