package board

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// CreateBoard Create Board godoc
// @Summary      Create a Board
// @Description  create a board by user Id
// @Tags         board
// @Accept       json
// @Produce      json
// @Param        board   path      int  true  "Account ID"
// @Success      200  {object}  Board
// @Router       /board [post]
func CreateBoard(c *fiber.Ctx) error {

	log.Println("check")
	var body Board

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	board := Board{
		Id:      1,
		Title:   body.Title,
		Content: body.Content,
		UserId:  body.UserId,
	}
	return c.Status(fiber.StatusCreated).JSON(board)
}

// FindBoard Find a Board godoc
// @Summary      Find a Board
// @Description  Find a board by user Id
// @Tags         board
// @Accept       json
// @Produce      json
// @Param        boardId   path      int  true  "Board ID"
// @Success      200  {object}  Board
// @Router       /board/{boardId} [get]
func FindBoard(c *fiber.Ctx) error {
	board := Board{
		Id:      123,
		Title:   "Check Title",
		Content: "Check Content",
		UserId:  "Check UserId",
	}
	return c.Status(fiber.StatusCreated).JSON(board)
}

func FindAllBoard(c *fiber.Ctx) error {
	return nil
}

func UpdateBoard(c *fiber.Ctx) error {
	return nil
}

func DeleteBoard(c *fiber.Ctx) error {
	return nil
}
