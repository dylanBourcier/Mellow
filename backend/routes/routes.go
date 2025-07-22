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
)

func SetupRoutes(services *bootstrap.Services) *http.ServeMux {
	mux := http.NewServeMux()

	// Authentification
	auth.RegisterAuthRoutes(mux, services.UserService, services.AuthService)

	// Utilisateurs
	users.RegisterUserRoutes(mux, services.UserService, services.AuthService)
	// Publications (posts)
	posts.RegisterPostRoutes(mux, services.PostService, services.AuthService, services.UserService, services.GroupService)

	// Groupes
	groups.RegisterGroupRoutes(mux, services.GroupService, services.AuthService)

	// Notifications
	notifications.RegisterNotificationRoutes(mux)

	// Messages privés + groupes
	messages.RegisterMessageRoutes(mux)

	// Modération (admin)
	admin.RegisterAdminRoutes(mux)

	// comments
	comments.RegisterCommentRoutes(mux, services.PostService, services.CommentService, services.AuthService)

	return mux
}
