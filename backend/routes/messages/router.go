package messages

import (
	msg "mellow/controllers/messages"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

// /messages/:userId
func MessageUserRouter(msgService services.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.URL.Path, "/messages/group/") {
			return // évite conflit avec groupe
		}

		userId := strings.TrimPrefix(r.URL.Path, "/messages/")
		if userId == "" || strings.Contains(userId, "/") {
			utils.RespondError(w, http.StatusNotFound, "Utilisateur introuvable", utils.ErrUserNotFound)
			return
		}

		switch r.Method {
		case http.MethodGet:
			msg.GetConversation(msgService, userId)(w, r)
		case http.MethodPost:
			msg.SendMessage(msgService, userId)(w, r)
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
		msg.GetGroupMessages(w, r, groupId)
	case http.MethodPost:
		msg.SendGroupMessage(w, r, groupId)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
	}
}
