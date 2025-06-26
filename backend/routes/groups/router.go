package groups

import (
	"mellow/controllers/groups"
	"mellow/utils"
	"net/http"
	"strings"
)

// /groups → POST (create), GET (list)
func GroupRootRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		groups.CreateGroup(w, r)
	case http.MethodGet:
		groups.GetAllGroups(w, r)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
	}
}

// /groups/posts/:id → GET (voir posts groupe), POST (ajouter post)
func GroupPostsRouter(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/groups/posts/")
	if id == "" || strings.Contains(id, "/") {
		utils.RespondError(w, http.StatusNotFound, "Ressource introuvable", utils.ErrGroupNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		groups.GetGroupPosts(w, r, id)
	case http.MethodPost:
		groups.AddGroupPost(w, r, id)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
	}
}
