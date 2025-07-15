package posts

import (
	"mellow/controllers/posts"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterPostRoutes(mux *http.ServeMux, PostService services.PostService, authService services.AuthService) {
	// Créer un post OU voir tous les posts
	mux.Handle("/posts", utils.ChainHTTP(posts.CreatePost(PostService), middlewares.AuthMiddleware(authService)))

	// Voir, éditer, supprimer un post spécifique
	mux.HandleFunc("/posts/", PostRouter)

	// Liker ou unliker un post
	//mux.HandleFunc("/posts/like/", LikeRouter)
}
