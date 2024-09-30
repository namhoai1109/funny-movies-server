package link

import (
	"context"
	"funnymovies/internal/model"
	dbutil "funnymovies/util/db"
	httputil "funnymovies/util/http"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Route struct {
	service Service
}

type Service interface {
	List(ctx context.Context, lq *dbutil.ListQueryCondition) ([]*model.LinkResponse, error)
	Total(ctx context.Context) (int64, error)
}

func NewRoute(service Service, echoGroup *echo.Group) {
	route := &Route{service}
	echoGroup.GET("", route.list)
	echoGroup.GET("/total", route.total)
}

func (r *Route) list(c echo.Context) error {
	lq, err := httputil.ReqListQuery(c)
	if err != nil {
		return err
	}

	res, err := r.service.List(c.Request().Context(), lq)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (r *Route) total(c echo.Context) error {
	res, err := r.service.Total(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, TotalLinkResponse{Total: res})

}
