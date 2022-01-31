package main

import (
	"github.com/WalkingMileageService/BMS/router"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	// Side Effect import for auto-generated swagger documentation
	//_ "medium_go_fiber_swagger/docs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @contact.name API Support
// @contact.email youremail@provider.com
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))
	router.SetupRoutes(app)
	app.Listen(":3000")
}
