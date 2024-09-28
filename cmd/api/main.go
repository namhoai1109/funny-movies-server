package main

import (
	"funnymovies/config"
	dbutil "funnymovies/util/db"
	"funnymovies/util/server"

	authenuser "funnymovies/internal/api/authen/user"
	userrepository "funnymovies/internal/repository/user"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := dbutil.New(cfg.DbDsn, true)
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	// * Initialize HTTP server
	e := server.New(&server.Config{
		Port: cfg.Port,
	})

	// --- repository
	userRepository := userrepository.NewRepository()

	// -- controller
	authenUserController := authenuser.New(db, userRepository)

	// --route
	authenRouter := e.Group("/authen")
	authenuser.NewRoute(authenUserController, authenRouter.Group("/user"))

	server.Start(e)
}
