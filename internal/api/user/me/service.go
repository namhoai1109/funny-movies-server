package me

import (
	dbutil "funnymovies/util/db"

	"gorm.io/gorm"
)

type Me struct {
	db             *gorm.DB
	userRepository UserRepository
}

func New(
	db *gorm.DB,
	userRepository UserRepository,
) *Me {
	return &Me{
		db:             db,
		userRepository: userRepository,
	}
}

type UserRepository interface {
	dbutil.Intf
}
