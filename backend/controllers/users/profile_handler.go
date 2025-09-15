package users

import (
	"fmt"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
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

		defer r.Body.Close()
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Failed to get user from context", utils.ErrUnauthorized)
			return
		}

		userInfo, err := userService.GetUserByID(r.Context(), uid.String())

		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Failed to get user from context", utils.ErrUnauthorized)
			return
		}

		if uid.String() != id {
			utils.RespondError(w, http.StatusNotFound, "You are not authorized to update this user", utils.ErrUserNotFound)
			return
		}

		if r.Method != http.MethodPut {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		// Limite taille max (par ex 10 Mo)
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Failed to parse multipart form: "+err.Error(), utils.ErrInvalidPayload)
			return
		}

		// Récupère les champs texte
		user := models.User{}
		user.UserID = uid

		username := r.FormValue("username")
		if username != "" {
			user.Username = username
		} else {
			user.Username = ""
		}
		fmt.Println("Username:", user.Username)

		firstname := r.FormValue("firstname")
		if firstname != "" {
			user.Firstname = firstname
		} else {
			user.Firstname = ""
		}
		fmt.Println("Firstname:", user.Firstname)

		lastname := r.FormValue("lastname")
		if lastname != "" {
			user.Lastname = lastname
		} else {
			user.Lastname = ""
		}
		fmt.Println("Lastname:", user.Lastname)

		privacy := r.FormValue("privacy")
		if privacy != "" {
			user.Privacy = privacy
		} else {
			user.Privacy = ""
		}
		fmt.Println("Privacy:", user.Privacy)
		user.Description = nil
		if desc := r.FormValue("description"); desc != "" {
			user.Description = &desc
		}
		fmt.Println("Description:", user.Description)

		// Parse birthdate (exemple format "2006-01-02")
		bdStr := r.FormValue("birthdate")
		if bdStr != "" {
			birthdate, err := time.Parse("2006-01-02", bdStr)
			if err != nil {
				utils.RespondError(w, http.StatusBadRequest, "Invalid birthdate format", utils.ErrInvalidPayload)
				return
			}
			user.Birthdate = birthdate
		}
		fmt.Println("Birthdate:", user.Birthdate)
		//password ne change pas, il n'est pas modifiable dans le form
		user.Password = ""

		file, header, err := r.FormFile("avatar")
		var image_url *string
		if err == nil {
			image_url, err = utils.HandleImageUpload(header, file, []string{"uploads", "avatars"})
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to upload image", err)
				return
			}
			user.ImageURL = image_url
			// Supprimer l'ancienne image si elle existe
			if userInfo.ImageURL != nil {
				os.Remove(filepath.Join("uploads", "avatars", *userInfo.ImageURL))
			}
		}
		// Appelle le service (il hash le mdp, valide, insert)
		if err := userService.UpdateUser(r.Context(), &user); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "User creation failed: "+err.Error(), utils.ErrInvalidPayload)
			// Supprimer le fichier si l'utilisateur n'a pas été créé
			if user.ImageURL != nil {
				os.Remove(filepath.Join("uploads", "avatars", *user.ImageURL))
			}
			return
		}

		utils.RespondJSON(w, http.StatusCreated, "User created successfully", nil)
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
