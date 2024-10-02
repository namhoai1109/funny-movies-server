package me

import (
	"encoding/json"
	dbutil "funnymovies/util/db"
	"funnymovies/util/testutil"
	"net/http/httptest"
	"testing"

	"funnymovies/internal/model"
	userrepository "funnymovies/internal/repository/user"

	userautho "funnymovies/internal/api/user/autho"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Test struct {
	route      *Route
	requestCtx RequestContext
}

type RequestContext interface {
	GetRequestContext() (echo.Context, *httptest.ResponseRecorder)
}

func NewTest() *Test {
	cfg := testutil.MockConfig
	db, err := dbutil.New(cfg.DbDsn, false)
	if err != nil {
		panic(err)
	}

	userRepository := userrepository.NewRepository()
	userAuthoService := userautho.New()

	service := &Me{
		db:             db,
		userRepository: userRepository,
	}

	r := &Route{service: service, autho: userAuthoService}
	return &Test{
		route:      r,
		requestCtx: testutil.New(),
	}
}

var test = NewTest()

func TestView(t *testing.T) {
	ctx, rec := test.requestCtx.GetRequestContext()
	// mock parsing token from Header
	ctx.Set("id", testutil.MockUser.ID)
	ctx.Set("email", testutil.MockUser.Email)

	err := test.route.view(ctx)

	// Assertions
	if assert.NoError(t, err) {
		assert.NotContains(t, rec.Body.String(), "password")

		response := &model.UserResponse{}
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			panic(err)
		}

		assert.NotEmpty(t, response)
		assert.Equal(t, testutil.MockUser.ID, response.ID)
		assert.Equal(t, testutil.MockUser.Email, response.Email)
	}
}

func TestNonAuthorizedView(t *testing.T) {
	ctx, _ := test.requestCtx.GetRequestContext()
	err := test.route.view(ctx)

	// Assertions
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "User not found")
	}
}
