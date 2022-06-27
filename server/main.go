package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zcubbs/pulse/server/configs"
	"github.com/zcubbs/pulse/server/routes"
	"github.com/zcubbs/pulse/server/utils"
	"os"
)

func main() {
	// TODO: get from env variable
	os.Setenv("TZ", "Europe/Paris")
	utils.CheckTimeZone()

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	configs.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.WsRoutes(app)      // Register ws routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Run Websocket Hub
	go utils.RunHub()

	// Start gRPC client
	defer utils.StartGrpcClient()()

	// Start server (with graceful shutdown).
	utils.StartServer(app)
}
