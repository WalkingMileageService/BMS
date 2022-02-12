package board

import (
	"github.com/labstack/echo/v4"
)

// @Summary Create Board
// @Description Create Board
// @Accept json
// @Produce json
// @Param board body BoardInfo true "BoardInfo"
// @Success 200 {object} Board
// @Router /board [post]
func CreateBoard(context echo.Context) error {
	return nil
}

// @Summary Find Board
// @Description Find Board
// @Accept json
// @Produce json
// @Param board param int64 true "BaordId"
// @Success 200 {object} Board
// @Router /board/{boardId} [get]
func FindBoard(context echo.Context) error {
	return nil
}

func FindAllBoard(context echo.Context) error {
	return nil
}

func UpdateBoard(context echo.Context) error {
	return nil
}

func DeleteBoard(context echo.Context) error {
	return nil
}
