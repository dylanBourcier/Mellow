package notifications

import (
	"mellow/controllers/notifications"
	"net/http"
)

func RegisterNotificationRoutes(mux *http.ServeMux) {
	// Voir toutes ses notifications
	mux.HandleFunc("/notifications", notifications.GetNotificationsHandler)

	// Marquer comme lue
	mux.HandleFunc("/notifications/read/", notifications.MarkAsReadHandler)
}
