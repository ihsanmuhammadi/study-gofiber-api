package route

import (
	"fiber-api/config"
	"fiber-api/handler"
	"fiber-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App)  {
	// Assets
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	// Login needs
	r.Post("/login", handler.LoginHandler)

	// CRUD
	r.Get("/user", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)
}
