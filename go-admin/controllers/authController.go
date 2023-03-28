package controllers

import (
	"go-admin/database"
	"go-admin/models"
	"go-admin/util"

	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// Parse data from the request and return error if there is one
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Retutn error if password and password confirm do not match
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match.",
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

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	// Parse data from the request and return error if there is one
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	// If user is not found...
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User not found.",
		})
	}

	// Compare the entered password with the attempting password and throw error if the passwords don't match
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect password.",
		})
	}

	// Create claims using JWT
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Store the token into a cookie

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success!",
	})
}

type Claims struct {
	jwt.StandardClaims
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout success.",
	})
}

// Below are methods to change own profile info
func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	// Parse data from the request and return error if there is one
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	// Parse data from the request and return error if there is one
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Retutn error if password and password confirm do not match
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match.",
		})
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)
	user := models.User{
		Id: uint(userId),
	}

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}
