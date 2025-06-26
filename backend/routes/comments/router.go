package comments

import (
	"mellow/controllers/comments"
	"mellow/utils"
	"net/http"
	"strings"
)

// Routeur général : redirige selon la méthode + nature de l'ID (postId vs commentId)
func CommentRouter(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/comments/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodPost:
		comments.AddComment(w, r, id) // id = postId
	case http.MethodGet:
		comments.GetComments(w, r, id) // id = postId
	case http.MethodPut:
		comments.UpdateComment(w, r, id) // id = commentId
	case http.MethodDelete:
		comments.DeleteComment(w, r, id) // id = commentId
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
	}
}
