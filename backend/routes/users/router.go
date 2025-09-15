package users

import (
	"fmt"
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

func SearchUsersHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		groupId := r.URL.Query().Get("groupId")
		viewerID, err := utils.GetUserIDFromContext(r.Context())
		fmt.Println(viewerID)
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", err)
			return
		}
		excludeGroupMembers := r.URL.Query().Get("excludeGroupMembers") == "true"
		if query == "" {
			utils.RespondError(w, http.StatusBadRequest, "Query parameter 'q' is required", utils.ErrBadRequest)
			return
		}

		users, err := userService.SearchUsers(r.Context(), viewerID.String(), query, groupId, excludeGroupMembers)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, err.Error(), err)
			return
		}
		if len(users) == 0 {
			utils.RespondJSON(w, http.StatusOK, "No users found", nil)
			return
		}
		for _, user := range users {
			if user.ImageURL != nil {
				user.ImageURL = utils.GetFullImageURLAvatar(user.ImageURL)
			}
		}
		utils.RespondJSON(w, http.StatusOK, "Users retrieved successfully", users)

	}
}
