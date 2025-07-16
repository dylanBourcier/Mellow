package posts

import (
	"mellow/services"
	"net/http"
)

func RegisterPostRoutes(mux *http.ServeMux, PostService services.PostService, authService services.AuthService) {
	// Créer un post OU voir tous les posts
	mux.Handle("/posts", http.HandlerFunc(PostRootRouter(PostService, authService)))

	// Voir, éditer, supprimer un post spécifique
	mux.HandleFunc("/posts/", PostRouter)

	// Liker ou unliker un post
	//mux.HandleFunc("/posts/like/", LikeRouter)
}
