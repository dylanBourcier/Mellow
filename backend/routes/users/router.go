package users

import (
	"mellow/controllers/users"
	"mellow/utils"
	"net/http"
	"strings"
)

func UserRouter(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id == "" || strings.Contains(id, "/") {
		utils.RespondError(w, http.StatusNotFound, "Utilisateur introuvable", utils.ErrUserNotFound)
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
		utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
	}
}

func FollowRouter(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/follow/")
	if id == "" || strings.Contains(id, "/") {
		utils.RespondError(w, http.StatusNotFound, "Utilisateur introuvable", utils.ErrUserNotFound)
		return
	}

	switch r.Method {
	case http.MethodPost:
		users.FollowUser(w, r, id)
	case http.MethodDelete:
		users.UnfollowUser(w, r, id)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
	}
}
