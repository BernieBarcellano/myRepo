// naghahandle ng mga data na iniinput ni user sa sign up text fields
package handler

import (
	"fmt"

	database "github.com/Ejil/studen_database/Database"
	"github.com/Ejil/studen_database/models"
	"github.com/gofiber/fiber/v2"
)

// Create New User
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

// Update the User Info
func UpdateUser(c *fiber.Ctx) error {
	var updateUser models.SignUp

	// Parse the request body and bind it to the updateUser struct
	if err := c.BodyParser(&updateUser); err != nil {
		return err
	}

	// Perform the update operation in the database
	result := database.DB.Model(&models.SignUp{}).Where("id = ?", updateUser.ID).Updates(updateUser)
	if result.Error != nil {
		return result.Error
	}

	// Check if any rows were affected by the update operation
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	// Return the updated user info
	return c.JSON(fiber.Map{
		"id":               updateUser.ID,
		"full_name":        updateUser.FullName,
		"student_number":   updateUser.Student_Number,
		"course":           updateUser.Course,
		"year_level":       updateUser.Year_Level,
		"email_address":    updateUser.Email_Address,
		"password":         updateUser.Password,
		"confirm_password": updateUser.Confirm_Password,
		"message":          "User updated successfully",
	})
}

// Delete the user information
func DeleteUser(c *fiber.Ctx) error {
	var deleteUser models.SignUp

	// Parse the request body and bind it to the deleteUser struct
	if err := c.BodyParser(&deleteUser); err != nil {
		return err
	}

	// Perform the delete operation in the database
	result := database.DB.Model(&models.SignUp{}).Where("id = ?", deleteUser.ID).Delete(deleteUser)
	if result.Error != nil {
		return result.Error
	}

	// Check if any rows were affected by the delete operation
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	// Return the deleteuser info
	return c.JSON(fiber.Map{
		"id":               deleteUser.ID,
		"full_name":        deleteUser.FullName,
		"student_number":   deleteUser.Student_Number,
		"course":           deleteUser.Course,
		"year_level":       deleteUser.Year_Level,
		"email_address":    deleteUser.Email_Address,
		"password":         deleteUser.Password,
		"confirm_password": deleteUser.Confirm_Password,
		"message":          "Delete successfully",
	})
}

// Get all User
func GetRegisteredUser(c *fiber.Ctx) error {
	db := database.DB
	var registeredusers []models.SignUp

	//find all users
	db.Find(&registeredusers)

	//if no user found
	if len(registeredusers) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "Error", "message": "No User Found", "data": nil})
	}

	//Else return the user list
	return c.JSON(fiber.Map{"data": registeredusers})
}
