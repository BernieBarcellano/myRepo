package main

import (
	"github.com/Ejil/studen_database/Database"
	"github.com/Ejil/studen_database/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//Start a new fiber app
	app := fiber.New()
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//Connect to the Database
	database.ConnectDB()

	//INITIAL ROUTE
	routes.Routes(app)

	app.Listen(":8080")
}
