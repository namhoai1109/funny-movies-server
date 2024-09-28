package user

import (
	dbutil "funnymovies/util/db"

	"gorm.io/gorm"
)

type AuthenUser struct {
	db             *gorm.DB
	userRepository UserRepository
}

func New(db *gorm.DB, userRepository UserRepository) *AuthenUser {
	return &AuthenUser{
		db:             db,
		userRepository: userRepository,
	}
}

type UserRepository interface {
	dbutil.Intf
}
