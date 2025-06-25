package posts

import "net/http"

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// TODO: créer un post
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	// TODO: retourner tous les posts (feed)
}

func GetPostByID(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: retourner post spécifique
}

func UpdatePost(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: modifier post
}

func DeletePost(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: supprimer post
}

func LikePost(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: liker
}

func UnlikePost(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: retirer le like
}
