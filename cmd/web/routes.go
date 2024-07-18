package web

import (
	"goth/internals/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
    e *echo.Echo,
    homeController *controllers.HomeController,
    countController *controllers.CountController,
    componentController *controllers.ComponentController,
    todoController *controllers.TodoController,
) {
    e.GET("/", homeController.Index)
    e.GET("/count", countController.Index)
    e.POST("/count", countController.IncrementCount)
    e.GET("/component", componentController.GetIndexView)
    e.GET("/todo", todoController.GetIndexView)
    e.POST("/todo", todoController.PostTodo)
    e.PUT("/todo", todoController.PutTodoStatus)
    e.DELETE("/todo", todoController.DeleteTodo)
}
