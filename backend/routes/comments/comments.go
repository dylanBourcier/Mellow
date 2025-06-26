package comments

import "net/http"

func RegisterCommentRoutes(mux *http.ServeMux) {
	// Ajouter un commentaire OU voir ceux d’un post
	mux.HandleFunc("/comments/", CommentRouter)
}
