// naghahandle ng mga data na iniinput ni user sa sign up text fields
package handler

import (
	"fmt"

	"github.com/Ejil/studen_database/Database"
	"github.com/Ejil/studen_database/models"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	db := database.DB
	newuser := new(models.SignUp)

	//Store body in the SignUp, return error if failed
	if err := c.BodyParser(newuser); err != nil {
		fmt.Println("err:", err)
		return c.Status(500).JSON(fiber.Map{"message": "Review your input"})
	}

	fmt.Println(newuser)

	//Create the new user
	if err := db.Create(&newuser).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "Error!", "message": "Could not create new user", "data": err})
	}

	//Return the created user
	return c.Status(201).JSON(newuser)

}
