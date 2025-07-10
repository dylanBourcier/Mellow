package routes

import (
	"net/http"

	"mellow/bootstrap"
	"mellow/routes/admin"
	"mellow/routes/auth"
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
	users.RegisterUserRoutes(mux, services.UserService)

	// Publications (posts)
	posts.RegisterPostRoutes(mux)

	// Groupes
	groups.RegisterGroupRoutes(mux)

	// Notifications
	notifications.RegisterNotificationRoutes(mux)

	// Messages privés + groupes
	messages.RegisterMessageRoutes(mux)

	// Modération (admin)
	admin.RegisterAdminRoutes(mux)

	return mux
}
