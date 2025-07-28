package users

import (
	"mellow/controllers/posts"
	ctr "mellow/controllers/users"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux, userService services.UserService, authSvc services.AuthService, postSvc services.PostService, notificationSvc services.NotificationService) {
	// Profil utilisateur : GET, PUT, DELETE
	mux.Handle("/users/", utils.ChainHTTP(UserRouter(userService), middlewares.RequireAuthMiddleware(authSvc)))

	// Posts d'un utilisateur : GET
	mux.Handle("/users/posts/", utils.ChainHTTP(posts.GetUserPosts(postSvc), middlewares.RequireAuthMiddleware(authSvc)))

	// Follow / Unfollow
	mux.Handle("/users/follow/", utils.ChainHTTP(FollowRouter(userService, notificationSvc), middlewares.RequireAuthMiddleware(authSvc)))

	// Voir followers / following
	mux.Handle("/users/followers/", utils.ChainHTTP(ctr.GetFollowersHandler(userService), middlewares.OptionalAuthMiddleware(authSvc)))
	mux.Handle("/users/following/", utils.ChainHTTP(ctr.GetFollowingHandler(userService), middlewares.OptionalAuthMiddleware(authSvc)))

	// Report post / user
	mux.HandleFunc("/users/report/", ctr.ReportHandler)
}
