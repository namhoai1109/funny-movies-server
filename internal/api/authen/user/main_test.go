package user

import (
	"fmt"
	dbutil "funnymovies/util/db"
	jwtutil "funnymovies/util/jwt"
	testutil "funnymovies/util/testutil"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	userrepository "funnymovies/internal/repository/user"

	"github.com/goombaio/namegenerator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Test struct {
	route      *Route
	requestCtx RequestContext
}

type RequestContext interface {
	PostRequestContext(inputJSON string) (echo.Context, *httptest.ResponseRecorder)
}

func NewTest() *Test {
	cfg := testutil.MockConfig
	db, err := dbutil.New(cfg.DbDsn, false)
	if err != nil {
		panic(err)
	}

	userRepository := userrepository.NewRepository()
	jwtUserService := jwtutil.New(cfg.JwtUserAlgo, cfg.JwtUserSecret, cfg.JwtUserDuration)
	service := &AuthenUser{
		db:             db,
		userRepository: userRepository,
		jwt:            jwtUserService,
	}

	r := &Route{service: service}
	return &Test{
		route:      r,
		requestCtx: testutil.New(),
	}
}

func generateCredentialInputJSON() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	name := nameGenerator.Generate()

	return fmt.Sprintf(`{"email":"%s@gmail.com","password":"%s"}`, name, name)
}

var test = NewTest()

// TestLogin test login
func TestLogin(t *testing.T) {
	inputJSON := fmt.Sprintf(`{"email":"%s","password":"%s"}`, testutil.MockUser.Email, testutil.MockUser.Password)
	ctx, rec := test.requestCtx.PostRequestContext(inputJSON)
	err := test.route.login(ctx)

	// Assertions
	if assert.NoError(t, err) {
		assert.True(t, strings.Contains(rec.Body.String(), "access_token"))
	}

}

// TestLoginFailed test login with invalid email
func TestLoginFailed(t *testing.T) {
	inputJSON := generateCredentialInputJSON()
	ctx, _ := test.requestCtx.PostRequestContext(inputJSON)
	err := test.route.login(ctx)

	// Assertions
	if assert.Error(t, err) {
		assert.ErrorContains(t, err, "invalid email")
	}
}

func TestRegister(t *testing.T) {
	inputJSON := generateCredentialInputJSON()
	ctx, rec := test.requestCtx.PostRequestContext(inputJSON)
	err := test.route.register(ctx)

	// Assertions
	if assert.NoError(t, err) {
		assert.True(t, strings.Contains(rec.Body.String(), "access_token"))
	}
}
