package me

import (
	"context"
	"funnymovies/internal/model"
	"funnymovies/util/server"
)

func (s *Me) View(ctx context.Context, authoUser *model.AuthoUser) (*MeResponse, error) {
	user := &model.User{}
	if err := s.userRepository.View(s.db, &user, "id = ? AND email = ?", authoUser.ID, authoUser.Email); err != nil {
		return nil, server.NewHTTPInternalError("User not found")
	}

	return &MeResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}
