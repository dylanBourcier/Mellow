package posts

import (
	"fmt"
	"io"
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
		fmt.Println("Error:", r)
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}
		var image_url *string
		file, header, err := r.FormFile("image")
		if err == nil {
			defer file.Close()
			fileName := uuid.New().String() + filepath.Ext(header.Filename)
			savePath := filepath.Join("uploads", fileName)
			if err := os.MkdirAll(filepath.Dir(savePath), os.ModePerm); err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to create directory", err)
				return
			}
			dst, err := os.Create(savePath)
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to create file", err)
				return
			}
			defer dst.Close()
			if _, err := io.Copy(dst, file); err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to save file", err)
				return
			}
			image_url = &fileName
		} else {
			if err != http.ErrMissingFile {
				utils.RespondError(w, http.StatusBadRequest, "Invalid image file", err)
				return
			}
			image_url = nil // No image uploaded
		}

		groupIdStr := r.FormValue("group_id")
		var groupID *uuid.UUID
		if groupIdStr != "" {
			groupIDValue, err := uuid.Parse(groupIdStr)
			if err != nil {
				utils.RespondError(w, http.StatusBadRequest, "Invalid group ID", utils.ErrInvalidPayload)
				return
			}
			groupID = &groupIDValue
		}

		post := models.Post{
			Title:      r.FormValue("title"),
			Content:    r.FormValue("content"),
			Visibility: r.FormValue("visibility"),
			UserID:     userID,
			ImageURL:   image_url,
			GroupID:    groupID,
		}
		fmt.Println("Post to create:", post)

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
