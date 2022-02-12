package main

import (
	"github.com/WalkingMileageService/BMS/router"
	"github.com/labstack/echo"
)

func main() {
	app := echo.New()
	router.SetupRoutes(app)

	app.Start(":8080")
}
