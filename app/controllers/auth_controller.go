package controllers

import (
	"book-restapi/app/models"
	"book-restapi/pkg/utils"
	"book-restapi/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func UserSignUp(c fiber.Ctx) error {
	//Create a new user auth struct
	signUp := &models.SignUp{}

	//Checking received data from JSON Body
	if err := c.BodyParser(signUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//Create a new validator for a User Model
	validate := utils.NewValidator()

	//Validate signup fields
	if err := validate.Struct(signUp); err != nil {
		//Return, if some fields are not valid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidationErrors(err),
		})
	}

	// Create database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//Checking role from sign up data
	role, err := utils.VerifyRole(signUp.UserRole)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//Create a new user struct
	user := &models.User{}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.Email = signUp.Email
	user.PasswordHash = utils.GeneratePassword(signUp.Password)
	user.UserStatus = 1
	user.UserRole = role

	//Validate user fields
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidationErrors(err),
		})
	}

	//Create a new user with validated data
	if err := db.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user.PasswordHash = ""

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}
