package users

import (
	"mellow/controllers/users"
	"mellow/services"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux, userService services.UserService) {
	// Profil utilisateur : GET, PUT, DELETE
	mux.HandleFunc("/users/", UserRouter)

	// Follow / Unfollow
	mux.HandleFunc("/users/follow/", FollowRouter)

	// Voir followers / following
	mux.HandleFunc("/users/followers/", users.GetFollowersHandler)
	mux.HandleFunc("/users/following/", users.GetFollowingHandler)

	// Report post / user
	mux.HandleFunc("/users/report/", users.ReportHandler)
}
