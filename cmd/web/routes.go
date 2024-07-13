package web

import (
	"goth/internals/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, homeController *controllers.HomeController, countController *controllers.CountController, componentController *controllers.ComponentController) {
    e.GET("/", homeController.Index)
    e.GET("/count", countController.Index)
    e.POST("/count", countController.IncrementCount)
    e.GET("/component", componentController.Index)
}
