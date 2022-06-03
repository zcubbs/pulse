package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/zcubbs/pulse/controllers"
)

func WsRoutes(app *fiber.App) {
	route := app.Group("/ws")
	route.Get("/", websocket.New(controllers.HandleWS))
}
