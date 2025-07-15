package groups

import (
	"mellow/controllers/groups"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterGroupRoutes(mux *http.ServeMux, groupSvc services.GroupService, authSvc services.AuthService) {
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

	// Voir les groupes auxquels l'utilisateur a adhéré
	mux.Handle("/groups/joined", utils.ChainHTTP(groups.GetGroupsJoinedByUser(groupSvc), middlewares.AuthMiddleware(authSvc)))
}
