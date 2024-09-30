package me

import (
	"context"
	"funnymovies/internal/model"
	"funnymovies/util/server"
)

func (s *Me) View(ctx context.Context, authoUser *model.AuthoUser) (*model.UserResponse, error) {
	user := &model.User{}
	if err := s.userRepository.View(s.db, &user, "id = ? AND email = ?", authoUser.ID, authoUser.Email); err != nil {
		return nil, server.NewHTTPInternalError("User not found")
	}

	return user.ToResponse(), nil
}
