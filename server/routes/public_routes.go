package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zcubbs/pulse/server/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(app *fiber.App) {
	root := app.Group("/")
	route := app.Group("/api/v1")

	root.Get("/health", controllers.HandleHealthCheck)
	route.Get("/health", controllers.HandleHealthCheck)
	route.Post("/event/gitlab", controllers.HandleGitlabPipelineEvent)
}
