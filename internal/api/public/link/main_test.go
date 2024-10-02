package link

import (
	"encoding/json"
	dbutil "funnymovies/util/db"
	testutil "funnymovies/util/testutil"
	"net/http/httptest"
	"testing"

	"funnymovies/internal/model"
	linkrepository "funnymovies/internal/repository/link"

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

	linkRepository := linkrepository.NewRepository()
	service := &Link{
		db:             db,
		linkRepository: linkRepository,
	}

	r := &Route{service: service}
	return &Test{
		route:      r,
		requestCtx: testutil.New(),
	}
}

var test = NewTest()

func TestList(t *testing.T) {
	ctx, rec := test.requestCtx.GetRequestContext()
	err := test.route.list(ctx)

	//Assertions
	if assert.NoError(t, err) {
		response := []*model.LinkResponse{}
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			panic(err)
		}
		assert.NotEmpty(t, response)
	}
}

func TestTotal(t *testing.T) {
	ctx, rec := test.requestCtx.GetRequestContext()
	err := test.route.total(ctx)

	//Assertions
	if assert.NoError(t, err) {
		response := TotalLinkResponse{}
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			panic(err)
		}
		assert.NotEmpty(t, response)

		ctx, rec := test.requestCtx.GetRequestContext()
		err := test.route.list(ctx)
		if assert.NoError(t, err) {
			listResponse := []*model.LinkResponse{}
			if err := json.Unmarshal(rec.Body.Bytes(), &listResponse); err != nil {
				panic(err)
			}
			assert.NotEmpty(t, listResponse)
			assert.Equal(t, len(listResponse), int(response.Total))
		}

	}
}
