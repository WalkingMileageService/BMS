package board

import "github.com/gofiber/fiber/v2"

func CreateBoardService(c *fiber.Ctx) (board Board, statusCode int) {

	var body Board

	err := c.BodyParser(&body)
	if err != nil {
		//TODO make return struct
		return body, fiber.StatusBadRequest
	}

	board = Board{
		Id:      1,
		Title:   body.Title,
		Content: body.Content,
		UserId:  body.UserId,
	}
	return board, fiber.StatusCreated
}
