package groups

import (
	"mellow/controllers/groups"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

// /groups → POST (create), GET (list)
func GroupRootRouter(groupSvc services.GroupService, authSvc services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler := utils.ChainHTTP(groups.CreateGroup(groupSvc), middlewares.RequireAuthMiddleware(authSvc))
			handler.ServeHTTP(w, r)
		case http.MethodGet:
			handler := utils.ChainHTTP(groups.CreateGroup(groupSvc), middlewares.OptionalAuthMiddleware(authSvc))
			handler.ServeHTTP(w, r)
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
		}
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

// /groups/:id → PUT (update) [Future: GET, DELETE]
func GroupRouter(groupService services.GroupService, authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/groups/")
		if id == "" || strings.Contains(id, "/") {
			utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			return
		}

		switch r.Method {
		case http.MethodPut:
			handler := utils.ChainHTTP(groups.UpdateGroup(groupService, id), middlewares.RequireAuthMiddleware(authService))
			handler.ServeHTTP(w, r)
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
		}
	}
}
