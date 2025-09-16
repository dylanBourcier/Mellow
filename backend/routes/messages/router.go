package messages

import (
	"net/http"
	"strings"

	msg "mellow/controllers/messages"
	"mellow/services"
	"mellow/utils"
)

// /messages
func MessageRouter(msgService services.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
			return
		}
		msg.GetRecentConversations(msgService)(w, r)
	}
}

// /messages/:userId
func MessageUserRouter(msgService services.MessageService, userSvc services.UserService) http.HandlerFunc {
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
			msg.GetConversation(msgService, userSvc, userId)(w, r)
		case http.MethodPost:
			msg.SendMessage(msgService, userSvc, userId)(w, r)
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
		}
	}
}

// /messages/group/:groupId
func MessageGroupRouter(msgService services.MessageService, userSvc services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupId := strings.TrimPrefix(r.URL.Path, "/messages/group/")
		if groupId == "" || strings.Contains(groupId, "/") {
			utils.RespondError(w, http.StatusNotFound, "Groupe introuvable", utils.ErrGroupNotFound)
			return
		}

		switch r.Method {
		case http.MethodGet:
			msg.GetGroupMessages(msgService, groupId)(w, r)
		case http.MethodPost:
			msg.SendGroupMessage(msgService, userSvc, groupId)(w, r)
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
		}
	}
}
