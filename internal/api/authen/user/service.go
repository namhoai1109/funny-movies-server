package user

import (
	dbutil "funnymovies/util/db"
	"time"

	"gorm.io/gorm"
)

type AuthenUser struct {
	db             *gorm.DB
	userRepository UserRepository
	jwt            JWT
}

func New(db *gorm.DB, userRepository UserRepository, jwt JWT) *AuthenUser {
	return &AuthenUser{
		db:             db,
		userRepository: userRepository,
		jwt:            jwt,
	}
}

type UserRepository interface {
	dbutil.Intf
}

type JWT interface {
	GenerateToken(claims map[string]interface{}, expire *time.Time) (string, int, error)
}
