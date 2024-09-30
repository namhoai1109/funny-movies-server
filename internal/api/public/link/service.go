package link

import (
	dbutil "funnymovies/util/db"

	"gorm.io/gorm"
)

type Link struct {
	db             *gorm.DB
	linkRepository LinkRepository
}

func New(
	db *gorm.DB,
	linkRepository LinkRepository,
) *Link {
	return &Link{
		db:             db,
		linkRepository: linkRepository,
	}
}

type LinkRepository interface {
	dbutil.Intf
}
