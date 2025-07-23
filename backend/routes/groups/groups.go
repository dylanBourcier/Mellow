package groups

import (
	"mellow/controllers/groups"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterGroupRoutes(mux *http.ServeMux, groupSvc services.GroupService, postSvc services.PostService, authSvc services.AuthService) {
	// Créer un groupe / voir tous les groupes
	mux.Handle("/groups", GroupRootRouter(groupSvc, authSvc))

	// Voir un groupe spécifique
	mux.Handle("/groups/", utils.ChainHTTP(GroupRouter(groupSvc, authSvc), middlewares.OptionalAuthMiddleware(authSvc)))

	// Voir les posts ou poster dans un groupe
	mux.HandleFunc("/groups/posts/", GroupPostsRouter(groupSvc, postSvc, authSvc))

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

	// Voir les groupes auxquels l'utilisateur n'est pas membre
	mux.Handle("/groups/not-joined", utils.ChainHTTP(groups.GetAllGroupsWithoutUser(groupSvc), middlewares.RequireAuthMiddleware(authSvc)))

}
