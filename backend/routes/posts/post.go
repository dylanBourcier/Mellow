package posts

import "net/http"

func RegisterPostRoutes(mux *http.ServeMux) {
	// Créer un post OU voir tous les posts
	mux.HandleFunc("/posts", PostRootRouter)

	// Voir, éditer, supprimer un post spécifique
	mux.HandleFunc("/posts/", PostRouter)

	// Liker ou unliker un post
	//mux.HandleFunc("/posts/like/", LikeRouter)
}
