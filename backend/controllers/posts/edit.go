package posts

import (
	"encoding/json"
	"errors"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

func UpdatePost(postService services.PostService, id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var payload models.UpdatePostPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid JSON", utils.ErrInvalidJSON)
			return
		}

		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		err = postService.UpdatePost(r.Context(), id, userID.String(), payload.Title, payload.Content)
		if err != nil {
			switch err {
			case utils.ErrPostNotFound:
				utils.RespondError(w, http.StatusNotFound, "Post not found", err)
			case utils.ErrForbidden:
				utils.RespondError(w, http.StatusForbidden, "Forbidden", err)
			case utils.ErrContentTooLong, utils.ErrContentTooShort, utils.ErrInvalidPayload:
				utils.RespondError(w, http.StatusBadRequest, "Invalid post", err)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to update post", err)
			}
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Post updated successfully", nil)
	}
}

func DeletePost(postService services.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}
		id := strings.TrimPrefix(r.URL.Path, "/posts/")
		if id == "" || strings.Contains(id, "/") {
			utils.RespondError(w, http.StatusNotFound, "Post not found", utils.ErrPostNotFound)
			return
		}

		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if err := postService.DeletePost(r.Context(), id, userID.String()); err != nil {
			if errors.Is(err, utils.ErrPostNotFound) {
				utils.RespondError(w, http.StatusNotFound, "Post not found", utils.ErrPostNotFound)
				return
			}
			if errors.Is(err, utils.ErrUnauthorized) {
				utils.RespondError(w, http.StatusForbidden, "Forbidden", utils.ErrUnauthorized)
				return
			}
			utils.RespondError(w, http.StatusInternalServerError, "Failed to delete post", err)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Post deleted successfully", nil)
	}
}
