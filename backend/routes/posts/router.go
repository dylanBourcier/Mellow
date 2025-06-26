package posts

import (
	"mellow/controllers/posts"
	"net/http"
	"strings"
)

// /posts → POST (create) ou GET (list)
func PostRootRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		posts.CreatePost(w, r)
	case http.MethodGet:
		posts.GetAllPosts(w, r)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

// /posts/:id → GET (view one), PUT (edit), DELETE (delete)
func PostRouter(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/posts/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		posts.GetPostByID(w, r, id)
	case http.MethodPut:
		posts.UpdatePost(w, r, id)
	case http.MethodDelete:
		posts.DeletePost(w, r, id)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

// /posts/like/:id → POST (like) ou DELETE (unlike)
func LikeRouter(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/posts/like/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodPost:
		posts.LikePost(w, r, id)
	case http.MethodDelete:
		posts.UnlikePost(w, r, id)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
