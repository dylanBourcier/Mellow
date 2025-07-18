package comments

import (
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func AddComment(commentService services.CommentService, postService services.PostService, postId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: ajouter un commentaire au post postId
		defer r.Body.Close()
		//s'assurer que le post existe
		if _, err := postService.IsPostExisting(r.Context(), postId); err != nil {
			utils.RespondError(w, http.StatusNotFound, "Post not found", utils.ErrPostNotFound)
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
		var imageURL *string
		if err == nil {
			imageURL, err = utils.HandleImageUpload(header, file, []string{"uploads"})
			if err != nil {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to upload image", err)
				return
			}
		}
		content := r.FormValue("content")
		if content == "" {
			utils.RespondError(w, http.StatusBadRequest, "Content cannot be empty", utils.ErrInvalidPayload)
			return
		}
		postIdValue, err := uuid.Parse(postId)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid post ID", utils.ErrInvalidPayload)
			return
		}

		comment := models.Comment{
			PostID:   postIdValue,
			UserID:   userID,
			Content:  &content,
			ImageURL: imageURL,
		}

		if err := commentService.CreateComment(r.Context(), &comment); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to add comment", err)
			if imageURL != nil {
				// Supprimer l'image si l'ajout du commentaire échoue
				os.Remove(filepath.Join("uploads", *imageURL))
			}
			return
		}
		if comment.ImageURL != nil {
			comment.ImageURL = utils.GetFullImageURL(comment.ImageURL) // Convert relative URL to full URL
		}
		utils.RespondJSON(w, http.StatusCreated, "Comment added successfully", comment)
	}
}
