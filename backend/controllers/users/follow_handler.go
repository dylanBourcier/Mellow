package users

import (
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

func FollowUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		targetID := strings.TrimPrefix(r.URL.Path, "/users/follow/")
		followerID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if err := userService.FollowUser(r.Context(), followerID.String(), targetID); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to follow user", utils.ErrInternalServerError)
			return
		}
		utils.RespondJSON(w, http.StatusOK, "User followed", nil)
	}
}

func UnfollowUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		targetID := strings.TrimPrefix(r.URL.Path, "/users/unfollow/")
		followerID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if err := userService.UnfollowUser(r.Context(), followerID.String(), targetID); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to unfollow user", utils.ErrInternalServerError)
		}
		utils.RespondJSON(w, http.StatusOK, "User unfollowed", nil)
	}
}

func GetFollowersHandler(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/users/followers/")
		followers, err := userService.GetFollowers(r.Context(), id)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get followers", utils.ErrInternalServerError)
			return
		}
		utils.RespondJSON(w, http.StatusOK, "Followers", followers)
	}
}

func GetFollowingHandler(userSerice services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/users/following/")
		following, err := userSerice.GetFollowing(r.Context(), id)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get following", utils.ErrInternalServerError)
			return
		}
		utils.RespondJSON(w, http.StatusOK, "Following", following)
	}
}
