package router

import (
	"github.com/WalkingMileageService/BMS/board"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func SetupRoutes(app *echo.Echo) {
	// echo middleware func
	app.Use(middleware.Logger())                             //Setting logger
	app.Use(middleware.Recover())                            //Recover from panics anywhere in the chain
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //CORS Middleware
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// debug 모드로 사용하기 위해서는 디버그 설정을 true로 변경
	//app.Debug = true

	app.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy!")
	})

	app.GET("/board/:boardId", board.FindBoard)
	app.GET("/boards", board.FindAllBoard)
	app.POST("/board", board.CreateBoard)
	app.PUT("/board", board.UpdateBoard)
	app.DELETE("/board", board.DeleteBoard)
}
