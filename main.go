package main

import (
	"github.com/WalkingMileageService/BMS/router"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	_ "github.com/WalkingMileageService/BMS/docs"
)

// @title Board Management Server API
// @version 1.0
// @description This is a board managemet server api for WalkingMileageService
// @contact.name jyson
// @contact.email jysohn0825@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /
func main() {
	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// TODO DB Conncet, error exception handler + api sender setting

	router.SetupRoutes(app)

	app.Listen(":3001")
}
