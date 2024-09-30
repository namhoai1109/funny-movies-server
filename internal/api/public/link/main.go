package link

import (
	"context"
	"fmt"
	"funnymovies/internal/model"
	dbutil "funnymovies/util/db"
	"funnymovies/util/server"
)

func (s *Link) List(ctx context.Context, lq *dbutil.ListQueryCondition) ([]*model.LinkResponse, error) {
	links := []*model.Link{}
	if err := s.linkRepository.List(s.db.Preload("User"), &links, lq, nil); err != nil {
		return nil, server.NewHTTPInternalError(fmt.Sprintf("Error when list link: %v", err.Error()))
	}

	res := []*model.LinkResponse{}
	for _, link := range links {
		res = append(res, link.ToResponse())
	}

	return res, nil
}

func (s *Link) Total(ctx context.Context) (int64, error) {
	count := int64(0)
	if err := s.db.Limit(-1).Offset(-1).Model(&model.Link{}).Count(&count).Error; err != nil {
		return 0, server.NewHTTPInternalError(fmt.Sprintf("Error when count link: %v", err.Error()))
	}

	return count, nil
}
