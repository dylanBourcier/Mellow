package posts

import (
	"fmt"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func CreatePost(postService services.PostService) http.HandlerFunc {
	// TODO: créer un post
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid form data", utils.ErrInvalidPayload)
			return
		}
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		file, header, err := r.FormFile("image")
		var image_url *string
		if err == nil {

			image_url, err = utils.HandleImageUpload(header, file, []string{"uploads"})
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to upload image", err)
				return
			}
		}
		groupIdStr := r.FormValue("postOn")
		var visibility string
		var groupID *uuid.UUID
		fmt.Println("Group ID String:", groupIdStr)
		if groupIdStr != "" {
			if groupIdStr == "everyone" {
				groupID = nil // Post visible to everyone
				visibility = r.FormValue("visibility")
			} else {
				groupIDValue, err := uuid.Parse(groupIdStr)
				if err != nil {
					utils.RespondError(w, http.StatusBadRequest, "Invalid group ID", utils.ErrInvalidPayload)
					return
				}
				groupID = &groupIDValue
				visibility = "public"

			}
		}
		var viewers []string
		if visibility == "private" {
			viewers = r.Form["selectedFollowers"]
		}

		post := models.Post{
			Title:      r.FormValue("title"),
			Content:    r.FormValue("content"),
			Visibility: visibility,
			UserID:     userID,
			ImageURL:   image_url,
			GroupID:    groupID,
			Viewers:    viewers,
		}
		fmt.Println(post)

		if err := postService.CreatePost(r.Context(), &post); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to create post", err)
			if image_url != nil {
				os.Remove(filepath.Join("uploads", *image_url))
			}
			return
		}
		utils.RespondJSON(w, http.StatusCreated, "Post created successfully", nil)
	}
}
