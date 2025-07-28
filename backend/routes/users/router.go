package users

import (
	ctr "mellow/controllers/users"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

func UserRouter(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		if id == "" || strings.Contains(id, "/") {
			utils.RespondError(w, http.StatusNotFound, "User not found", utils.ErrUserNotFound)
			return
		}

		switch r.Method {
		case http.MethodGet:
			ctr.GetUserProfileHandler(userService)(w, r)
		case http.MethodPut:
			ctr.UpdateUserProfileHandler(userService)(w, r)
		case http.MethodDelete:
			ctr.DeleteUserHandler(userService)(w, r)
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrBadRequest)
		}
	}
}

func FollowRouter(userService services.UserService, notificationSvc services.NotificationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/users/follow/")
		if id == "" || strings.Contains(id, "/") {
			utils.RespondError(w, http.StatusNotFound, "User not found", utils.ErrUserNotFound)
			return
		}

		switch r.Method {
		case http.MethodPost:
			ctr.SendFollowRequest(userService, notificationSvc)(w, r)
		case http.MethodDelete:
			ctr.UnfollowUser(userService)(w, r)
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrBadRequest)
		}
	}
}
