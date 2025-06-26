package messages

import (
	"mellow/controllers/messages"
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
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		messages.GetConversation(w, r, userId)
	case http.MethodPost:
		messages.SendMessage(w, r, userId)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

// /messages/group/:groupId
func MessageGroupRouter(w http.ResponseWriter, r *http.Request) {
	groupId := strings.TrimPrefix(r.URL.Path, "/messages/group/")
	if groupId == "" || strings.Contains(groupId, "/") {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		messages.GetGroupMessages(w, r, groupId)
	case http.MethodPost:
		messages.SendGroupMessage(w, r, groupId)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
