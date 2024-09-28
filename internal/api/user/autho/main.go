package autho

import (
	"encoding/json"
	"fmt"
	"funnymovies/internal/model"
	structutil "funnymovies/util/struct"

	"github.com/labstack/echo/v4"
)

func (s *Autho) User(c echo.Context) *model.AuthoUser {
	tokenClaims := structutil.ToMap(&model.UserTokenClaims{})
	for k := range tokenClaims {
		tokenClaims[k] = c.Get(k)
	}

	tokenMarshal, err := json.Marshal(tokenClaims)
	if err != nil {
		fmt.Println("Failed to marshal token claims")
		return nil
	}

	authoUser := &model.AuthoUser{}
	if err := json.Unmarshal(tokenMarshal, authoUser); err != nil {
		fmt.Println("Failed to unmarshal token claims")
		return nil
	}

	return authoUser
}
