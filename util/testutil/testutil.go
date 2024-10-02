package testutil

import (
	"funnymovies/config"
	"funnymovies/internal/model"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
)

type Test struct {
	echo *echo.Echo
}

var (
	MockConfig = &config.Configuration{
		JwtUserAlgo:     "HS256",
		JwtUserSecret:   "Ez2C3g33NWE2CRQp4C2eS8DA8Cr4CAqC",
		JwtUserDuration: 31536001,
		DbDsn:           "postgres://funnymovies:funnymovies123@localhost:5432/funnymovies?sslmode=disable&connect_timeout=5",
	}

	MockUser = &model.User{
		ID:       1,
		Email:    "testemail@gmail.com",
		Password: "password",
	}
)

func New() *Test {
	return &Test{echo: echo.New()}
}

func (s *Test) PostRequestContext(inputJSON string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(inputJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return s.echo.NewContext(req, rec), rec
}

func (s *Test) GetRequestContext() (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return s.echo.NewContext(req, rec), rec
}
