package link

import (
	"funnymovies/internal/model"
	dbutil "funnymovies/util/db"
)

func NewRepository() *DB {
	return &DB{dbutil.NewDB(&model.Link{})}
}

type DB struct {
	*dbutil.DB
}
