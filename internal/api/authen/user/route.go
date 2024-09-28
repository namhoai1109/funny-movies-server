package user

import (
	"github.com/labstack/echo/v4"
)

type Route struct {
	controller Controller
}

type Controller interface{}

func NewRoute(controller Controller, echoGroup *echo.Group) {
	route := &Route{controller}
	echoGroup.POST("/login", route.login)
}

func (r *Route) login(c echo.Context) error {
	return c.JSON(200, "OK")
}
