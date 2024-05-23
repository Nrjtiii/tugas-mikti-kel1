package app

import (
	"github.com/haerul-umam/tugas-mikti-kel1/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Validator = helper.NewValidator()
	e.HTTPErrorHandler = helper.BindAndValidate

	return e
}