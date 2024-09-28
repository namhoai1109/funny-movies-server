package user

import (
	"funnymovies/internal/model"
	"funnymovies/util/server"
	structutil "funnymovies/util/struct"
	"time"
)

func (s *AuthenUser) getToken(user *model.User) (*model.AuthToken, error) {
	userTokenClaims := &model.UserTokenClaims{
		ID:    user.ID,
		Email: user.Email,
	}
	claims := structutil.ToMap(userTokenClaims)

	// default expire time is 24 hours
	timeExpire := time.Time.Add(time.Now(), 24*time.Hour)
	token, expiresIn, err := s.jwt.GenerateToken(claims, &timeExpire)
	if err != nil {
		return nil, server.NewHTTPInternalError("Error when generate token")
	}

	return &model.AuthToken{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
	}, nil
}
