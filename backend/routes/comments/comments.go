package comments

import (
	"mellow/services"
	"net/http"
)

func RegisterCommentRoutes(mux *http.ServeMux, postService services.PostService, commentService services.CommentService, authService services.AuthService) {
	// Ajouter un commentaire OU voir ceux d’un post
	mux.HandleFunc("/comments/", CommentRouter(postService, commentService, authService))
}
