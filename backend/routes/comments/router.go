package comments

import (
	"mellow/controllers/comments"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

// Routeur général : redirige selon la méthode + nature de l'ID (postId vs commentId)
func CommentRouter(postService services.PostService, commentService services.CommentService, authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/comments/")
		if id == "" || strings.Contains(id, "/") {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case http.MethodPost:
			handler := utils.ChainHTTP(comments.AddComment(commentService, postService, id), middlewares.RequireAuthMiddleware(authService)) // id = postId
			handler.ServeHTTP(w, r)
		case http.MethodGet:
			handler := utils.ChainHTTP(comments.GetComments(commentService, postService, id), middlewares.OptionalAuthMiddleware(authService)) // id = postId
			handler.ServeHTTP(w, r)
		case http.MethodPut:
			handler := utils.ChainHTTP(comments.UpdateComment(commentService, id), middlewares.RequireAuthMiddleware(authService)) // id = commentId
			handler.ServeHTTP(w, r)
		case http.MethodDelete:
			comments.DeleteComment(w, r, id) // id = commentId
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
		}
	}
}
