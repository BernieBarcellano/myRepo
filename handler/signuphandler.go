package handler

import (
	"fmt"
	"net/http"

	database "github.com/Ejil/studen_database/Database"
	"github.com/Ejil/studen_database/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GenerateJWT(userID uint) (string, error) {
	// Define the claims for the JWT
	claims := jwt.MapClaims{
		"user_id": userID,
		// Add other desired claims...
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the token string
	jwtSecret := []byte("your_secret_key") // Replace with your secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Login(context *fiber.Ctx) error {
	db := database.DB
	newUser := new(models.SignUp)

	err := context.BodyParser(newUser)
	if err != nil {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "request failed",
		})
	}

	// Perform authentication logic
	// Example:
	user := &models.SignUp{}
	err = db.Where("email = ? AND password = ?", newUser.Email_Address, newUser.Password).First(user).Error
	if err != nil {
		return context.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "success login",
		})
	}

	// Assuming authentication is successful, generate a token
	// Example:
	token, err := GenerateJWT(user.ID)
	if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "failed to generate token",
		})
	}

	return context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login successful",
		"token":   token,
	})
}

func SignUp(c *fiber.Ctx) error {
	db := database.DB
	newUser := new(models.SignUp)

	// Store body in the newUser, return error if failed
	if err := c.BodyParser(newUser); err != nil {
		fmt.Println("err:", err)
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Done creating student",
		})
	}

	fmt.Println(newUser)

	// Create the new user
	if err := db.Create(newUser).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "Error!",
			"message": "Could not create new user",
			"data":    err,
		})
	}

	// Return the created user
	return c.Status(http.StatusCreated).JSON(newUser)
}

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
	return c.JSON(&fiber.Map{
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

func DeleteUser(c *fiber.Ctx) error {
	var deleteUser models.SignUp

	// Parse the request body and bind it to the deleteUser struct
	if err := c.BodyParser(&deleteUser); err != nil {
		return err
	}

	// Perform the delete operation in the database
	result := database.DB.Model(&models.SignUp{}).Where("id = ?", deleteUser.ID).Delete(&deleteUser)
	if result.Error != nil {
		return result.Error
	}

	// Check if any rows were affected by the delete operation
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	// Return the deleteUser info
	return c.JSON(&fiber.Map{
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

func GetRegisteredUser(c *fiber.Ctx) error {
	db := database.DB
	var registeredUsers []models.SignUp

	// Find all users
	db.Find(&registeredUsers)

	// If no user found
	if len(registeredUsers) == 0 {
		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
			"status":  "Error",
			"message": "No User Found",
			"data":    nil,
		})
	}

	// Else return the user list
	return c.JSON(&fiber.Map{
		"data": registeredUsers,
	})
}
func AttendanceUser(c *fiber.Ctx) error {
	var AttendanceUser models.Attendance

	// Parse the request body and bind it to the AttendanceUser struct
	if err := c.BodyParser(&AttendanceUser); err != nil {
		return err
	}

	// Perform the attendance operation in the database
	result := database.DB.Model(&models.SignUp{}).Where("id = ?", AttendanceUser.ID).Updates(&AttendanceUser)
	if result.Error != nil {
		return result.Error
	}

	// Check if any rows were affected by the attendance operation
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	// To show AttendanceUser info
	return c.JSON(&fiber.Map{
		"id":           AttendanceUser.ID,
		"full_name":    AttendanceUser.Full_Name,
		"subject":      AttendanceUser.Subject,
		"block_no":     AttendanceUser.Block_No,
		"today_date":   AttendanceUser.Today_Date,
		"current_time": AttendanceUser.Current_Time,
		"message":      "attendance successfull",
	})
}
