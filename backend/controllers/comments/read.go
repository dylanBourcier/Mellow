package comments

import (
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func GetComments(commentSvc services.CommentService, postSvc services.PostService, postId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Vérifier que le postId est valide
		if postId == "" {
			utils.RespondError(w, http.StatusBadRequest, "Invalid post ID", utils.ErrInvalidPayload)
			return
		}

		// Récupérer les commentaires pour le post
		comments, err := commentSvc.GetCommentsByPostID(r.Context(), postId)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Error while geting comments", err)
			return
		}
		// Répondre avec les commentaires récupérés
		utils.RespondJSON(w, http.StatusOK, "comments retrieved successfully", comments)

	}
}
