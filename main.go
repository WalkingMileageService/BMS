package main

import (
	_ "github.com/WalkingMileageService/BMS/docs"
	"github.com/WalkingMileageService/BMS/router"
	"github.com/labstack/echo/v4"
)

// @title Board Management Server Swagger
// @version 1.0
// @host localhost:8080
// @BasePath /bms
func main() {
	app := echo.New()
	router.SetupRoutes(app)
	app.Start(":8080")
}
