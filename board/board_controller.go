package board

import "github.com/gofiber/fiber/v2"

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func CreateBoard(c *fiber.Ctx) error {
	return nil
}

func FindBoard(c *fiber.Ctx) error {
	return nil
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
