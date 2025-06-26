package posts

import "net/http"

func LikePost(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: liker
}

func UnlikePost(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: retirer le like
}
