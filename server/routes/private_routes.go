package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zcubbs/pulse/controllers"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	//route.Get("/projects", middleware.JWTProtected(), controllers.HandleGetProjects)
	route.Get("/projects", controllers.HandleGetProjects)
	route.Get("/watch/:projectId/latest", controllers.HandleGetLatestWatchReport)
	route.Get("/watch/entries", controllers.HandleGetWatchReports)
	route.Post("/watch/wsRefresh", controllers.HandleNotifyRefresh)
}
