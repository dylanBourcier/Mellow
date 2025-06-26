package messages

import (
	"mellow/controllers/messages"
	"mellow/utils"
	"net/http"
	"strings"
)

// /messages/:userId
func MessageUserRouter(w http.ResponseWriter, r *http.Request) {
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
		messages.GetConversation(w, r, userId)
	case http.MethodPost:
		messages.SendMessage(w, r, userId)
	default:
		utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
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
