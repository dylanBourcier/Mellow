package routes

import (
	"net/http"

	"mellow/routes/admin"
	"mellow/routes/auth"
	"mellow/routes/groups"
	"mellow/routes/messages"
	"mellow/routes/notifications"
	"mellow/routes/posts"
	"mellow/routes/users"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Authentification
	auth.RegisterAuthRoutes(mux)

	// Utilisateurs
	users.RegisterUserRoutes(mux)

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
