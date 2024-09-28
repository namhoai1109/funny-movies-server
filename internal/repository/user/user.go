package user

import (
	"funnymovies/internal/model"
	dbutil "funnymovies/util/db"
)

func NewRepository() *DB {
	return &DB{dbutil.NewDB(&model.User{})}
}

type DB struct {
	*dbutil.DB
}
