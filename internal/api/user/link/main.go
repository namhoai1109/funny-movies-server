package link

import (
	"context"
	"fmt"
	"funnymovies/internal/model"
	"funnymovies/util/server"
)

func (s *Link) Create(ctx context.Context, user *model.AuthoUser, url string) error {
	if err := s.linkRepository.Create(s.db, &model.Link{Url: url, UserID: user.ID}); err != nil {
		server.NewHTTPInternalError(fmt.Sprintf("Error when create link: %v", err.Error()))
	}

	return nil
}
