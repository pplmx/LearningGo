package handlers

import (
	"LearningGo/fiber/models"
	"github.com/gofiber/fiber/v2"
)

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func GetUserByID(c *fiber.Ctx) error {
	uid := c.Params("uid")
	user := new(models.User)
	user.UID = uid
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get a user by id.",
		Data:    user,
	})
}

func GetUserByName(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get a user by name.",
		Data:    user,
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get all users.",
		Data:    users,
	})
}

func DeleteUserByID(c *fiber.Ctx) error {
	uid := c.Params("uid")
	var user models.User
	user.UID = uid
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success delete a user by id.",
		Data:    user,
	})
}
