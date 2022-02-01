package router

import (
	"github.com/WalkingMileageService/BMS/board"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/board/:boardId", board.FindBoard)
	app.Get("/boards", board.FindAllBoard)
	app.Post("/board", board.CreateBoard)
	app.Put("/board", board.UpdateBoard)
	app.Delete("/board", board.DeleteBoard)
}
