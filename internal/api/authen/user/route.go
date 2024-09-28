package user

import (
	"context"
	"fmt"
	"funnymovies/internal/model"
	"funnymovies/util/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Route struct {
	service Service
}

type Service interface {
	Login(ctx context.Context, req *CredentialRequest) (*model.AuthToken, error)
	Register(ctx context.Context, req *CredentialRequest) (*model.AuthToken, error)
}

func NewRoute(service Service, echoGroup *echo.Group) {
	route := &Route{service}
	echoGroup.POST("/login", route.login)
	echoGroup.POST("/register", route.register)
}

func (r *Route) login(c echo.Context) error {
	req := CredentialRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := r.service.Login(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (r *Route) register(c echo.Context) error {
	req := CredentialRequest{}
	if err := c.Bind(&req); err != nil {
		return server.NewHTTPInternalError(fmt.Sprintf("Error when bind request: %v", err.Error()))
	}

	res, err := r.service.Register(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
