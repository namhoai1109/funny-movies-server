package link

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
	Create(ctx context.Context, user *model.AuthoUser, url string) error
}

func NewRoute(service Service, autho model.Autho, echoGroup *echo.Group) {
	route := &Route{service, autho}
	echoGroup.POST("", route.create)
}

func (r *Route) create(c echo.Context) error {
	req := LinkCreationRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := r.service.Create(c.Request().Context(), r.autho.User(c), req.Url)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
