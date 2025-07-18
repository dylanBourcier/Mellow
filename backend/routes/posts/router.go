package posts

import (
	"mellow/controllers/posts"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

// /posts → POST (create) ou GET (list)
func PostRootRouter(PostService services.PostService, authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler := utils.ChainHTTP(posts.CreatePost(PostService), middlewares.RequireAuthMiddleware(authService))
			handler.ServeHTTP(w, r)
		case http.MethodGet:
			handler := utils.ChainHTTP(posts.GetFeedPosts(PostService), middlewares.OptionalAuthMiddleware(authService))
			handler.ServeHTTP(w, r)
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
		}
	}
}

// /posts/:id → GET (view one), PUT (edit), DELETE (delete)
func PostRouter(w http.ResponseWriter, r *http.Request, postService services.PostService) {
	id := strings.TrimPrefix(r.URL.Path, "/posts/")
	if id == "" || strings.Contains(id, "/") {
		utils.RespondError(w, http.StatusNotFound, "Post introuvable", utils.ErrPostNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		posts.GetPostByID(postService)(w, r)
	case http.MethodPut:
		posts.UpdatePost(w, r, id)
	case http.MethodDelete:
		posts.DeletePost(w, r, id)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
	}
}

// /posts/like/:id → POST (like) ou DELETE (unlike)
func LikeRouter(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/posts/like/")
	if id == "" || strings.Contains(id, "/") {
		utils.RespondError(w, http.StatusNotFound, "Post introuvable", utils.ErrPostNotFound)
		return
	}

	switch r.Method {
	case http.MethodPost:
		posts.LikePost(w, r, id)
	case http.MethodDelete:
		posts.UnlikePost(w, r, id)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
	}
}
