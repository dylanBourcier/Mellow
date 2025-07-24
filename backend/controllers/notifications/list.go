package notifications

import (
	"mellow/services"
	"mellow/utils"
	"net/http"
)

// GetNotifications retourne la liste des notifications de l'utilisateur connecté.
func GetNotifications(notificationService services.NotificationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		notifs, err := notificationService.GetUserNotifications(r.Context(), uid.String())
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get notifications", err)
			return
		}

		if len(notifs) == 0 {
			utils.RespondJSON(w, http.StatusOK, "No notifications", nil)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Notifications retrieved successfully", notifs)
	}
}
