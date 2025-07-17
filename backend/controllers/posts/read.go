package posts

import (
	"database/sql"
	"fmt"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	// TODO: retourner tous les posts (feed)
}

func GetPostByID(groupService services.GroupService, userService services.UserService, postService services.PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: retourner post spécifique
		//récupérer le postID depuis l'URL
		id := r.URL.Path[len("/posts/"):]
		fmt.Println("Post ID:", id)
		if id == "" || len(id) < 36 { // Assuming UUID length
			utils.RespondError(w, http.StatusNotFound, "Post not found", utils.ErrPostNotFound)
			return
		}
		userID, ok := r.Context().Value(utils.CtxKeyUserID).(string)
		if !ok {
			userID = "" // No user ID in context, maybe unauthenticated request
		}
		post, err := postService.GetPostByID(r.Context(), id, groupService, userService, userID)
		fmt.Println("Post retrieved:", post)
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
