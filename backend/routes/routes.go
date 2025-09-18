package routes

import (
	"net/http"

	"mellow/bootstrap"
	"mellow/routes/admin"
	"mellow/routes/auth"
	"mellow/routes/comments"
	"mellow/routes/groups"
	"mellow/routes/messages"
	"mellow/routes/notifications"
	"mellow/routes/posts"
	"mellow/routes/users"
	"mellow/websocket"
)

func SetupRoutes(services *bootstrap.Services) *http.ServeMux {
	mux := http.NewServeMux()

	// Authentification
	auth.RegisterAuthRoutes(mux, services.UserService, services.AuthService)

	// Utilisateurs
	users.RegisterUserRoutes(mux, services.UserService, services.AuthService, services.PostService, services.NotificationService)
	// Publications (posts)
	posts.RegisterPostRoutes(mux, services.PostService, services.AuthService, services.UserService, services.GroupService)

	// Groupes
	groups.RegisterGroupRoutes(mux, services.GroupService, services.PostService, services.AuthService, services.NotificationService, services.UserService, services.GroupJoinRequestService)

	// Notifications
	notifications.RegisterNotificationRoutes(mux, services.NotificationService, services.AuthService)

	// Messages privés + groupes
	messages.RegisterMessageRoutes(mux, services.MessageService, services.UserService, services.AuthService)
	websocket.RegisterDebugRoutes(mux)
	mux.HandleFunc("/ws/chat", websocket.WsHandler)

	// Modération (admin)
	admin.RegisterAdminRoutes(mux)

	// comments
	comments.RegisterCommentRoutes(mux, services.PostService, services.CommentService, services.AuthService)

	return mux
}
