package posts

import "net/http"

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	// TODO: retourner tous les posts (feed)
}

func GetPostByID(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: retourner post spécifique
}
