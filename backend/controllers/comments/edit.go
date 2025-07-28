package comments

import (
	"encoding/json"
	"errors"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

// UpdateComment retourne un handler permettant de mettre à jour un commentaire.
func UpdateComment(commentService services.CommentService, commentID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var payload struct {
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid JSON", utils.ErrInvalidJSON)
			return
		}
		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		err = commentService.UpdateComment(r.Context(), commentID, userID.String(), payload.Content)
		if err != nil {
			switch {
			case errors.Is(err, utils.ErrCommentNotFound):
				utils.RespondError(w, http.StatusNotFound, "Comment not found", err)
			case errors.Is(err, utils.ErrForbidden):
				utils.RespondError(w, http.StatusForbidden, "Forbidden", err)
			case errors.Is(err, utils.ErrContentTooLong), errors.Is(err, utils.ErrContentTooShort), errors.Is(err, utils.ErrInvalidPayload):
				utils.RespondError(w, http.StatusBadRequest, "Invalid content", err)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to update comment", err)
			}
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Comment updated successfully", nil)
	}
}
