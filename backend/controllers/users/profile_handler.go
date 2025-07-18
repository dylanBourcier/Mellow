package users

import (
	"encoding/json"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
	"time"
)

func GetUserProfileHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		if r.Method != http.MethodGet {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
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

		data := map[string]interface{}{
			"user_id":       user.UserID,
			"email":         user.Email,
			"username":      user.Username,
			"firstname":     user.Firstname,
			"lastname":      user.Lastname,
			"birthdate":     user.Birthdate,
			"role":          user.Role,
			"image_url":     utils.GetFullImageURLAvatar(user.ImageURL),
			"creation_date": user.CreationDate,
			"description":   user.Description,
		}

		utils.RespondJSON(w, http.StatusOK, "User retrieved", data)
	}
}

func UpdateUserProfileHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/users/")
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil || uid.String() != id {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		var req struct {
			Username    *string `json:"username"`
			Password    *string `json:"password"`
			Firstname   *string `json:"firstname"`
			Lastname    *string `json:"lastname"`
			Birthdate   *string `json:"birthdate"`
			Description *string `json:"description"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid JSON", utils.ErrInvalidJSON)
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
		if req.Birthdate != nil && *req.Birthdate != "" {
			bd, err := time.Parse("2006-01-02", *req.Birthdate)
			if err != nil {
				utils.RespondError(w, http.StatusBadRequest, "Invalid birthdate format. Expected format: YYYY-MM-DD", utils.ErrInvalidInput)
				return
			}
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
		if r.Method != http.MethodDelete {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/users/")
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil || uid.String() != id {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if err := userService.DeleteUser(r.Context(), id); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to delete user", utils.ErrInternalServerError)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "User deleted", nil)
	}
}
