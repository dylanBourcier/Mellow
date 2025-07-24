package notifications

import (
	"errors"
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

func MarkAsRead(notificationService services.NotificationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/notifications/read/")
		if r.Method != http.MethodPatch || id == "" || strings.Contains(id, "/") {
			http.NotFound(w, r)
			return
		}

		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if err := notificationService.MarkAsRead(r.Context(), id, uid.String()); err != nil {
			switch {
			case errors.Is(err, utils.ErrNotificationNotFound):
				utils.RespondError(w, http.StatusNotFound, "Notification not found", err)
			case errors.Is(err, utils.ErrForbidden):
				utils.RespondError(w, http.StatusForbidden, "Forbidden", err)
			case errors.Is(err, utils.ErrInvalidPayload):
				utils.RespondError(w, http.StatusBadRequest, "Invalid notification", err)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to mark notification as read", err)
			}
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Notification marked as read", nil)
	}
}
