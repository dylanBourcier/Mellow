package groups

import (
	"mellow/controllers/groups"
	"net/http"
)

func RegisterGroupRoutes(mux *http.ServeMux) {
	// Créer un groupe / voir tous les groupes
	mux.HandleFunc("/groups", GroupRootRouter)

	// Voir les posts ou poster dans un groupe
	mux.HandleFunc("/groups/posts/", GroupPostsRouter)

	// Rejoindre un groupe
	mux.HandleFunc("/groups/join/", groups.JoinGroupHandler)

	// Quitter un groupe
	mux.HandleFunc("/groups/leave/", groups.LeaveGroupHandler)

	// Voir les événements du groupe
	mux.HandleFunc("/groups/events/", groups.GroupEventsHandler)

	// Voir le chat de groupe
	mux.HandleFunc("/groups/chat/", groups.GroupChatHandler)
}
