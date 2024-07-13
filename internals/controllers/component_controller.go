package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ComponentController struct {}

func NewComponentController() *ComponentController {
	return &ComponentController{}
}

func (ctl *ComponentController) Index(c echo.Context) error {
	log.Println("TEst render")
	return c.Render(http.StatusOK, "componentIndex", nil)
}
