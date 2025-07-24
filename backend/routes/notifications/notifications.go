package notifications

import (
	"mellow/controllers/notifications"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func RegisterNotificationRoutes(mux *http.ServeMux, notificationService services.NotificationService, authService services.AuthService) {
	mux.Handle("/notifications", utils.ChainHTTP(NotificationRootRouter(notificationService, authService), middlewares.RequireAuthMiddleware(authService)))
	mux.Handle("/notifications/read/", utils.ChainHTTP(notifications.MarkAsRead(notificationService), middlewares.RequireAuthMiddleware(authService)))
}
