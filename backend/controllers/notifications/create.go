package notifications

import (
	"encoding/json"
	"errors"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"

	"github.com/google/uuid"
)

func CreateNotification(notificationService services.NotificationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed", utils.ErrMethodNotAllowed)
			return
		}

		var payload models.CreateNotificationPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid JSON", utils.ErrInvalidJSON)
			return
		}

		uid, err := uuid.Parse(payload.UserID)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid user ID", utils.ErrInvalidPayload)
			return
		}

		notif := models.Notification{
			UserID: uid,
			Type:   payload.Type,
		}

		if err := notificationService.CreateNotification(r.Context(), &notif); err != nil {
			switch {
			case errors.Is(err, utils.ErrInvalidPayload), errors.Is(err, utils.ErrUserNotFound):
				utils.RespondError(w, http.StatusBadRequest, "Invalid notification", err)
			default:
				utils.RespondError(w, http.StatusInternalServerError, "Failed to create notification", err)
			}
			return
		}

		utils.RespondJSON(w, http.StatusCreated, "Notification created successfully", map[string]any{"notification_id": notif.NotificationID})
	}
}
