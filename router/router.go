package router

import (
	"github.com/WalkingMileageService/BMS/board"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/board", board.FindBoard)
	app.Get("/api/boards", board.FindAllBoard)
	app.Post("/api/board", board.CreateBoard)
	app.Put("/api/board", board.UpdateBoard)
	app.Delete("/api/board", board.DeleteBoard)
}
