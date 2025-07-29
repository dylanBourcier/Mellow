package users

import (
	"encoding/json"
	"fmt"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

func GetUserProfileHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := strings.TrimPrefix(r.URL.Path, "/users/")
		if id == "" || len(id) < 36 { // Assuming UUID length
			utils.RespondError(w, http.StatusNotFound, "User not found", utils.ErrUserNotFound)
			return
		}
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		user, err := userService.GetUserProfileData(r.Context(), userID.String(), id)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get user", utils.ErrInternalServerError)
			return
		}
		if user == nil {
			utils.RespondError(w, http.StatusNotFound, "User not found", utils.ErrUserNotFound)
			return
		}
		fmt.Println("User Profile Data:", user)
		//Check if the user can view the profile
		if user.Privacy == "private" && userID.String() != id {
			// Check if the user is following
			isFollowing, err := userService.IsFollowing(r.Context(), userID.String(), id)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to check following status", utils.ErrInternalServerError)
				return
			}
			if !isFollowing {
				var limitedProfile models.UserProfileData
				limitedProfile.UserID = user.UserID
				limitedProfile.Username = user.Username
				limitedProfile.ImageURL = user.ImageURL
				limitedProfile.FollowStatus = user.FollowStatus
				limitedProfile.Privacy = user.Privacy
				description := "This profile is private. Follow to see more."
				limitedProfile.Description = &description
				utils.RespondJSON(w, http.StatusOK, "Limited", limitedProfile)
				return
			}

		}
		utils.RespondJSON(w, http.StatusOK, "User retrieved", user)
	}
}

func UpdateUserProfileHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := strings.TrimPrefix(r.URL.Path, "/users/")
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Failed to get user from context", utils.ErrUnauthorized)
			return
		}

		if uid.String() != id {
			utils.RespondError(w, http.StatusNotFound, "You are not authorized to update this user", utils.ErrUserNotFound)
			return
		}

		var req models.UpdateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondError(w, http.StatusBadRequest,
				"Invalid JSON", utils.ErrInvalidJSON)
			return
		}

		user, err := userService.GetUserByID(r.Context(), id)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get user", utils.ErrInternalServerError)
			return
		}
		if user == nil {
			utils.RespondError(w, http.StatusNotFound, "User not found", utils.ErrUserNotFound)
			return
		}

		if req.Username != nil {
			user.Username = *req.Username
		}
		if req.Firstname != nil {
			user.Firstname = *req.Firstname
		}
		if req.Lastname != nil {
			user.Lastname = *req.Lastname
		}
		if req.Description != nil {
			user.Description = req.Description
		}
		if req.Password != nil {
			hashed, err := utils.HashPassword(*req.Password)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to hash password", utils.ErrInternalServerError)
				return
			}
			user.Password = hashed
		}
		if req.Birthdate != nil {
			bd := *req.Birthdate
			user.Birthdate = bd
		}

		if err := userService.UpdateUser(r.Context(), user); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to update user", utils.ErrInternalServerError)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "User updated", nil)
	}
}

func DeleteUserHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := strings.TrimPrefix(r.URL.Path, "/users/")
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Failed to get user from context", utils.ErrUnauthorized)
			return
		}

		if uid.String() != id {
			utils.RespondError(w, http.StatusNotFound, "You are not authorized to delete this user", utils.ErrUserNotFound)
			return
		}

		if err := userService.DeleteUser(r.Context(), id); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to delete user", utils.ErrInternalServerError)
			return
		}
		defer r.Body.Close()

		utils.RespondJSON(w, http.StatusOK, "User deleted", nil)
	}
}
