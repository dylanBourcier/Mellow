package groups

import (
	"net/http"
	"strings"
)

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	// TODO: liste des groupes
}

func GetGroupPosts(w http.ResponseWriter, r *http.Request, id string) {
	// TODO: posts du groupe
}

func GroupEventsHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/groups/events/")
	if r.Method != http.MethodGet || id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	// TODO: événements du groupe
}

func GroupChatHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/groups/chat/")
	if r.Method != http.MethodGet || id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	// TODO: chat du groupe, plutôt à faire dans la partie message ??
}
