package messages

import (
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strings"
)

func MarkAsRead(service services.MessageService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		msgID := strings.TrimPrefix(r.URL.Path, "/messages/read/")
		if msgID == "" || strings.Contains(msgID, "/") {
			http.NotFound(w, r)
			return
		}

		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if err := service.MarkAsRead(r.Context(), msgID, uid.String()); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to mark message as read", err)
			return
		}
		utils.RespondJSON(w, http.StatusOK, "Message marked as read", nil)
	}
}
func MarkAsReadConversation(service services.MessageService, otherId string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		if err := service.MarkAsReadConversation(r.Context(), uid.String(), otherId); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to mark conversation as read", err)
			return
		}
		utils.RespondJSON(w, http.StatusOK, "Conversation marked as read", nil)
	}
}
