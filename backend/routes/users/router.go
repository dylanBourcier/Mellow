package users

import (
	"mellow/controllers/users"
	"net/http"
	"strings"
)

func UserRouter(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		users.GetUserProfileHandler(w, r, id)
	case http.MethodPut:
		users.UpdateUserProfileHandler(w, r, id)
	case http.MethodDelete:
		users.DeleteUserHandler(w, r, id)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func FollowRouter(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/follow/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodPost:
		users.FollowUser(w, r, id)
	case http.MethodDelete:
		users.UnfollowUser(w, r, id)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
