package routes

import (
	"github.com/Ejil/studen_database/handler"
	"github.com/gofiber/fiber/v2"
)

func Routes(r *fiber.App){
	r.Post("/newuser",handler.SignUp)
}