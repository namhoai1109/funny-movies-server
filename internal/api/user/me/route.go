package me

import (
	"context"
	"funnymovies/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Route struct {
	service Service
	autho   model.Autho
}

type Service interface {
	View(ctx context.Context, authoUser *model.AuthoUser) (*model.UserResponse, error)
}

func NewRoute(service Service, autho model.Autho, echoGroup *echo.Group) {
	route := &Route{service, autho}
	echoGroup.GET("", route.view)
}

func (r *Route) view(c echo.Context) error {
	res, err := r.service.View(c.Request().Context(), r.autho.User(c))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
