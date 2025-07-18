package posts

import (
	"mellow/controllers/posts"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterPostRoutes(mux *http.ServeMux, PostService services.PostService, authService services.AuthService, userService services.UserService, groupService services.GroupService) {
	// Créer un post OU voir tous les posts
	mux.Handle("/posts", http.HandlerFunc(PostRootRouter(PostService, authService)))

	// Voir, éditer, supprimer un post spécifique
	mux.Handle("/posts/", utils.ChainHTTP(posts.GetPostByID(PostService), middlewares.OptionalAuthMiddleware(authService)))

	// Liker ou unliker un post
	//mux.HandleFunc("/posts/like/", LikeRouter)
}
