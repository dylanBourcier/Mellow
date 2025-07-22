package comments

import (
	"mellow/services"
	"mellow/utils"
	"net/http"
)

// DeleteComment supprime un commentaire identifié par commentID.
// Seul l'auteur du commentaire peut le supprimer.
func DeleteComment(commentService services.CommentService, commentID string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		userID, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if err := commentService.DeleteComment(r.Context(), commentID, userID.String()); err != nil {
			switch err {
			case utils.ErrCommentNotFound:
				utils.RespondError(w, http.StatusNotFound, "Comment not found", err)
			case utils.ErrForbidden:
				utils.RespondError(w, http.StatusForbidden, "Forbidden", err)
			case utils.ErrInvalidPayload:
				utils.RespondError(w, http.StatusBadRequest, "Invalid comment ID", err)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to delete comment", err)
			}
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Comment deleted successfully", nil)
	}
}
