package groups

import (
	ctrl "mellow/controllers/groups"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

func RegisterGroupRoutes(mux *http.ServeMux, groupSvc services.GroupService, postSvc services.PostService, authSvc services.AuthService, notifSvc services.NotificationService, userSvc services.UserService, gjrSvc services.GroupJoinRequestService) {
	// Créer un groupe / voir tous les groupes
	mux.Handle("/groups", GroupRootRouter(groupSvc, authSvc))

	// Inviter un utilisateur dans un groupe
	mux.Handle("/groups/invite/", utils.ChainHTTP(ctrl.InviteUserToGroup(groupSvc, notifSvc), middlewares.RequireAuthMiddleware(authSvc)))

	//Repondre à une invitation de groupe
	mux.Handle("/groups/invite/answer/", utils.ChainHTTP(ctrl.AnswerGroupInvite(groupSvc, userSvc, notifSvc), middlewares.RequireAuthMiddleware(authSvc)))

	//Inviter un utilisateur dans un groupe

	// Voir les posts ou poster dans un groupe
	mux.HandleFunc("/groups/posts/", GroupPostsRouter(groupSvc, postSvc, authSvc))

	// Rejoindre un groupe
	mux.Handle("/groups/join/", utils.ChainHTTP(ctrl.JoinGroup(groupSvc), middlewares.RequireAuthMiddleware(authSvc)))

	// Quitter un groupe
	mux.Handle("/groups/leave/", utils.ChainHTTP(ctrl.LeaveGroupHandler(groupSvc), middlewares.RequireAuthMiddleware(authSvc)))

	// Voir les événements du groupe
	mux.Handle("/groups/events/", utils.ChainHTTP(GroupEventRouter(groupSvc, authSvc), middlewares.RequireAuthMiddleware(authSvc)))

	// Voir le chat de groupe
	mux.HandleFunc("/groups/chat/", ctrl.GroupChatHandler)

	// Voir les groupes auxquels l'utilisateur a adhéré
	mux.Handle("/groups/joined", utils.ChainHTTP(ctrl.GetGroupsJoinedByUser(groupSvc), middlewares.RequireAuthMiddleware(authSvc)))

	// Voir les groupes auxquels l'utilisateur n'est pas membre
	mux.Handle("/groups/not-joined", utils.ChainHTTP(ctrl.GetAllGroupsWithoutUser(groupSvc), middlewares.RequireAuthMiddleware(authSvc)))

	mux.Handle("/groups/events/vote/", utils.ChainHTTP(ctrl.InsertEventResponse(groupSvc), middlewares.RequireAuthMiddleware(authSvc)))

	// Join requests feature
	// POST /groups/{id}/join-requests
	mux.Handle("/groups/", utils.ChainHTTP(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/join-requests") && r.Method == http.MethodPost {
			ctrl.CreateJoinRequest(gjrSvc)(w, r)
			return
		}
		if strings.HasSuffix(r.URL.Path, "/join-requests") && r.Method == http.MethodGet {
			ctrl.ListJoinRequests(gjrSvc)(w, r)
			return
		}
		if strings.HasSuffix(r.URL.Path, "/join-requests/self") {
			ctrl.SelfJoinRequestRouter(gjrSvc)(w, r)
			return
		}
		if strings.Contains(r.URL.Path, "/join-requests/") && strings.HasSuffix(r.URL.Path, "/accept") {
			ctrl.AcceptJoinRequest(gjrSvc)(w, r)
			return
		}
		if strings.Contains(r.URL.Path, "/join-requests/") && strings.HasSuffix(r.URL.Path, "/reject") {
			ctrl.RejectJoinRequest(gjrSvc)(w, r)
			return
		}
		// fallthrough to default group router (already mounted above for /groups/)
		GroupRouter(groupSvc, authSvc)(w, r)
	}), middlewares.OptionalAuthMiddleware(authSvc)))

}
