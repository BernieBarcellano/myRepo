package routes

import (
	"github.com/Ejil/studen_database/handler"
	"github.com/gofiber/fiber/v2"
)

func Routes(r *fiber.App) {
	r.Post("/newuser", handler.SignUp)
	r.Put("/updateuser", handler.UpdateUser)//update student
	r.Delete("/deleteuser/:id", handler.DeleteUser)//delete data student
	r.Get("/getAllUser", handler.GetRegisteredUser)//show list of student
	r.Post("/loginuser", handler.Login)
}
