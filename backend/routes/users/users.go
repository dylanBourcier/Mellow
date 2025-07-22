package users

import (
	ctr "mellow/controllers/users"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux, userService services.UserService, authSvc services.AuthService) {
	// Profil utilisateur : GET, PUT, DELETE
	mux.Handle("/users/", utils.ChainHTTP(UserRouter(userService), middlewares.OptionalAuthMiddleware(authSvc)))

	// Follow / Unfollow
	mux.Handle("/users/follow/", utils.ChainHTTP(FollowRouter(userService), middlewares.RequireAuthMiddleware(authSvc)))

	// Voir followers / following
	mux.Handle("/users/followers/", utils.ChainHTTP(ctr.GetFollowersHandler(userService), middlewares.OptionalAuthMiddleware(authSvc)))
	mux.Handle("/users/following/", utils.ChainHTTP(ctr.GetFollowingHandler(userService), middlewares.OptionalAuthMiddleware(authSvc)))

	// Report post / user
	mux.HandleFunc("/users/report/", ctr.ReportHandler)
}
