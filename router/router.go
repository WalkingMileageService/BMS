package router

import (
	"github.com/WalkingMileageService/BMS/board"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"os"
)

func SetupRoutes() {
	app := echo.New()

	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format:           "${time_rfc3339_nano} INFO T[${tid}] U[${userId}] - [ - ] ${method} ${uri} ${status} (${latency_human})\n",
			CustomTimeFormat: "2006-01-02 15:04:05.00000",
			Output:           os.Stdout,
		}))
	app.Use(middleware.Recover())

	// debug 모드로 사용하기 위해서는 디버그 설정을 true로 변경
	//app.Debug = true

	root := app.Group("/bms")
	{
		api := root.Group("/api")
		{
			v1 := api.Group("/v1")
			{
				v1.GET("/board/:boardId", board.FindBoard)
				v1.GET("/boards", board.FindAllBoard)
				v1.POST("/board", board.CreateBoard)
				v1.PUT("/board", board.UpdateBoard)
				v1.DELETE("/board", board.DeleteBoard)
			}
		}

		actuator := root.Group("")
		{
			actuator.GET("/health", nil)
			actuator.GET("/shutdown", nil)
		}

		swagger := root.Group("/swagger")
		{
			swagger.GET("/*", echoSwagger.WrapHandler)
		}
	}

	log.Fatal(app.Start(":3000"))
}
