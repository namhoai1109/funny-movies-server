package user

import (
	"context"
	"funnymovies/internal/model"
	"funnymovies/util/crypter"
	"funnymovies/util/server"
)

func (s *AuthenUser) Login(ctx context.Context, req *CredentialRequest) (*model.AuthToken, error) {
	user := &model.User{}
	if err := s.userRepository.View(s.db, &user, "email = ?", req.Email); err != nil {
		return nil, server.NewHTTPInternalError("invalid email!")
	}

	if !crypter.CompareHashAndPassword(user.Password, req.Password) {
		return nil, server.NewHTTPInternalError("invalid password!")
	}

	return s.getToken(user)
}

func (s *AuthenUser) Register(ctx context.Context, req *CredentialRequest) (*model.AuthToken, error) {
	exist, err := s.userRepository.Exist(s.db, "email = ?", req.Email)
	if err != nil {
		return nil, server.NewHTTPInternalError(err.Error())
	}

	if exist {
		return nil, server.NewHTTPInternalError("email already exist!")
	}

	hashPassword, err := crypter.HashPassword(req.Password)
	if err != nil {
		return nil, server.NewHTTPInternalError("Error when hash password")
	}

	user := &model.User{
		Email:    req.Email,
		Password: hashPassword,
	}
	if err := s.userRepository.Create(s.db, &user); err != nil {
		return nil, server.NewHTTPInternalError(err.Error())
	}

	return s.getToken(user)
}
