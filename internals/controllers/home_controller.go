package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeController struct{}

func NewHomeController() *HomeController {
    return &HomeController{}
}

func (ctl *HomeController) Index(c echo.Context) error {
    log.Println("Rendering index.html template for / route")
    return c.Render(http.StatusOK, "index", nil) // Use "index.html" explicitly
}
