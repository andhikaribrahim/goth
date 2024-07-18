package controllers

import (
	"goth/internals/utils"
	"goth/views/components"

	"github.com/labstack/echo/v4"
)

type ComponentController struct {}

func NewComponentController() *ComponentController {
	return &ComponentController{}
}

func (cc *ComponentController) GetIndexView(c echo.Context) error {
	indexView := components.Index()

	return utils.RenderView(c, indexView)
}
