package messages

import (
	"mellow/controllers/messages"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

// /messages/:userId
func MessageUserRouter(messageService services.MessageService, authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/messages/group/") {
			return // évite conflit avec groupe
		}

		switch r.Method {
		case http.MethodGet:
			handler := utils.ChainHTTP(messages.GetConversation(messageService), middlewares.RequireAuthMiddleware(authService))
			handler.ServeHTTP(w, r)
		case http.MethodPost:
			handler := utils.ChainHTTP(messages.SendMessage(messageService), middlewares.RequireAuthMiddleware(authService))
			handler.ServeHTTP(w, r)
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
		}
	}
}

// /messages/group/:groupId
func MessageGroupRouter(w http.ResponseWriter, r *http.Request) {
	groupId := strings.TrimPrefix(r.URL.Path, "/messages/group/")
	if groupId == "" || strings.Contains(groupId, "/") {
		utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		messages.GetGroupMessages(w, r, groupId)
	case http.MethodPost:
		messages.SendGroupMessage(w, r, groupId)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
	}
}
