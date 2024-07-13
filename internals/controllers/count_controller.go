package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Count struct {
    Count int
}

type CountController struct {
    count *Count
}

func NewCountController() *CountController {
    return &CountController{
        count: &Count{Count: 0},
    }
}

func (ctl *CountController) Index(c echo.Context) error {
    log.Println("Rendering count/index.html template on route /count")
    return c.Render(http.StatusOK, "count/index", ctl.count) // Use "count/index.html" explicitly
}

func (ctl *CountController) IncrementCount(c echo.Context) error {
    ctl.count.Count++
    log.Println("Rendering count/index.html template after increment")
    return c.Render(http.StatusOK, "count/increment", ctl.count) // Use "count/index.html" explicitly
}
