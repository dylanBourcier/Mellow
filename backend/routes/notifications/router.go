package notifications

import (
	"mellow/controllers/notifications"
	"mellow/middlewares"
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func NotificationRootRouter(notificationService services.NotificationService, authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler := utils.ChainHTTP(notifications.CreateNotification(notificationService), middlewares.RequireAuthMiddleware(authService))
			handler.ServeHTTP(w, r)
		case http.MethodGet:
			handler := utils.ChainHTTP(http.HandlerFunc(notifications.GetNotificationsHandler), middlewares.RequireAuthMiddleware(authService))
			handler.ServeHTTP(w, r)
		default:
			utils.RespondError(w, http.StatusMethodNotAllowed, "Méthode non autorisée", utils.ErrBadRequest)
		}
	}
}
