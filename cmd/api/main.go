package main

import (
	"fmt"
	"funnymovies/config"
	dbutil "funnymovies/util/db"
	jwtutil "funnymovies/util/jwt"
	"funnymovies/util/server"

	authenuser "funnymovies/internal/api/authen/user"
	publiclink "funnymovies/internal/api/public/link"
	userautho "funnymovies/internal/api/user/autho"
	userlink "funnymovies/internal/api/user/link"
	userme "funnymovies/internal/api/user/me"
	linkrepository "funnymovies/internal/repository/link"
	userrepository "funnymovies/internal/repository/user"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println("Loaded config!")

	db, err := dbutil.New(cfg.DbDsn, true)
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	fmt.Println("db connected: " + db.Name())

	// * Initialize HTTP server
	e := server.New(&server.Config{
		Port: cfg.Port,
	})

	// --- authorization
	userAuthoService := userautho.New()

	// --- repository
	userRepository := userrepository.NewRepository()
	linkRepository := linkrepository.NewRepository()

	// -- service
	jwtUserService := jwtutil.New(cfg.JwtUserAlgo, cfg.JwtUserSecret, cfg.JwtUserDuration)
	authenUserService := authenuser.New(db, userRepository, jwtUserService)
	userMeService := userme.New(db, userRepository)
	userLinkService := userlink.New(db, linkRepository)
	publicLinkService := publiclink.New(db, linkRepository)

	// --route
	authenRouter := e.Group("/authen")
	authenuser.NewRoute(authenUserService, authenRouter.Group("/users"))

	userRouter := e.Group("/users")
	userRouter.Use(jwtUserService.MiddlewareFunction())
	userme.NewRoute(userMeService, userAuthoService, userRouter.Group("/me"))
	userlink.NewRoute(userLinkService, userAuthoService, userRouter.Group("/links"))

	publiclink.NewRoute(publicLinkService, e.Group("/links"))

	server.Start(e)
}
