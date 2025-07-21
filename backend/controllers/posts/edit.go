package posts

import (
	"encoding/json"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"
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

func DeletePost(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: supprimer post
}
