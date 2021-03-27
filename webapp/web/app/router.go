package app

import (
	"github.com/labstack/echo"
	"github.com/s-aska/favicon/webapp/web/app/controllers/top"
)

func registerRoutes(e *echo.Echo) {
	e.GET("/", top.Root)
	e.POST("/manual", top.Manual)
	e.POST("/auto", top.Auto)
}
