package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zcubbs/pulse/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/health", controllers.HandleHealthCheck)
	route.Post("/event/gitlab", controllers.HandleGitlabPipelineEvent)
}
