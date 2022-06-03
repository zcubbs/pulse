package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zcubbs/pulse/configs"
	"github.com/zcubbs/pulse/routes"
	"github.com/zcubbs/pulse/utils"
	"os"
)

func main() {
	os.Setenv("TZ", "Europe/Paris")
	utils.CheckTimeZone()

	// Load yaml config
	utils.LoadYamlConfig()

	// Init Database
	utils.ConnectToPostgresDB()

	// Init Rabbitmq connection
	rabbitmq, channel := utils.ConnectRabbitmq()
	defer channel.Close()
	defer rabbitmq.Close()

	// Setup Git Webhooks
	utils.SetupGitlabWebhook()

	// Launch event worker routine
	utils.LaunchEventWorker()

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

	// Start server (with graceful shutdown).
	utils.StartServer(app)
}
