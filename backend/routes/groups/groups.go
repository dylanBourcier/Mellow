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
	mux.Handle("/groups", GroupRootRouter(groupSvc, authSvc))

	// Voir les posts ou poster dans un groupe
	mux.HandleFunc("/groups/posts/", GroupPostsRouter)

	// Supprimer un groupe
	mux.Handle("/groups/", GroupRouter(groupSvc, authSvc))

	// Rejoindre un groupe
	mux.Handle("/groups/join/", utils.ChainHTTP(groups.JoinGroup(groupSvc), middlewares.RequireAuthMiddleware(authSvc)))

	// Quitter un groupe
	mux.Handle("/groups/leave/", utils.ChainHTTP(groups.LeaveGroupHandler(groupSvc), middlewares.RequireAuthMiddleware(authSvc)))

	// Voir les événements du groupe
	mux.HandleFunc("/groups/events/", groups.GroupEventsHandler)

	// Voir le chat de groupe
	mux.HandleFunc("/groups/chat/", groups.GroupChatHandler)

	// Voir les groupes auxquels l'utilisateur a adhéré
	mux.Handle("/groups/joined", utils.ChainHTTP(groups.GetGroupsJoinedByUser(groupSvc), middlewares.RequireAuthMiddleware(authSvc)))

	// Mettre à jour un groupe
	mux.HandleFunc("/groups/", GroupRouter(groupSvc, authSvc))
}
