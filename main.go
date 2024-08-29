package main

import (
	"fiber-api/database"
	"fiber-api/migration"
	"fiber-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	// Initial database
	database.DatabaseInit()

	// Migrasi database
	migration.RunMigration()

	// Memulai server
	app := fiber.New()

	// Initial route
	route.RouteInit(app)

	app.Listen(":9000")
}
