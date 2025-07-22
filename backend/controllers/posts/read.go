package posts

import (
	"database/sql"
	"fmt"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func GetFeedPosts(postSvc services.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Recuperer les paramètres de la requête, limit et l'offset
		limit := 10 // Default limit
		offset := 0 // Default offset
		query := r.URL.Query()
		if l := query.Get("limit"); l != "" {
			fmt.Sscanf(l, "%d", &limit)
		}
		if o := query.Get("offset"); o != "" {
			fmt.Sscanf(o, "%d", &offset)
		}
		if limit <= 0 || offset < 0 {
			utils.RespondError(w, http.StatusBadRequest, "Invalid limit or offset", utils.ErrInvalidPayload)
			return
		}
		// Format de la requete : /posts?limit=10&offset=0
		// Appeler le service pour récupérer les posts
		posts, err := postSvc.GetFeed(r.Context(), nil, limit, offset)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.RespondError(w, http.StatusNotFound, "No posts found", utils.ErrPostNotFound)
				return
			}
			utils.RespondError(w, http.StatusInternalServerError, "Internal Server Error", err)
			return
		}
		if len(posts) == 0 {
			utils.RespondJSON(w, http.StatusOK, "No posts found", nil)
			return
		}
		utils.RespondJSON(w, http.StatusOK, "Posts retrieved successfully", posts)
	}
}

func GetPostByID(postService services.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: retourner post spécifique
		//récupérer le postID depuis l'URL
		id := r.URL.Path[len("/posts/"):]
		if id == "" || len(id) < 36 { // Assuming UUID length
			utils.RespondError(w, http.StatusNotFound, "Post not found", utils.ErrPostNotFound)
			return
		}
		userID, ok := r.Context().Value(utils.CtxKeyUserID).(string)
		if !ok {
			userID = "" // No user ID in context, maybe unauthenticated request
		}
		post, err := postService.GetPostByID(r.Context(), id, userID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Post not found", http.StatusNotFound)
				return
			}
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		utils.RespondJSON(w, http.StatusOK, "post retrieved successfully", post)
	}

}
