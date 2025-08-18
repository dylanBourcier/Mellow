package messages

import (
	"mellow/services"
	"mellow/utils"
	"net/http"
	"strconv"
)

func GetConversation(messageService services.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		otherUserId := r.URL.Path[len("/messages/"):]
		userId, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", err)
			return
		}
		userIdStr := userId.String()
		if otherUserId == "" || userIdStr == "" {
			utils.RespondError(w, http.StatusBadRequest, "Invalid user IDs", utils.ErrInvalidPayload)
			return
		}
		if otherUserId == userIdStr {
			utils.RespondError(w, http.StatusBadRequest, "Cannot message yourself", utils.ErrInvalidPayload)
			return
		}

		limit := 10
		offset := 0
		query := r.URL.Query()
		if l := query.Get("limit"); l != "" {
			if limit2, err := strconv.Atoi(l); err == nil {
				limit = limit2
			}
		}
		if o := query.Get("offset"); o != "" {
			if offset2, err := strconv.Atoi(o); err == nil {
				offset = offset2
			}
		}

		// Appeler le service pour récupérer la conversation
		messages, err := messageService.GetConversation(r.Context(), userIdStr, otherUserId, limit, offset)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get conversation", err)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Message fetch successful", messages)
	}
}

func SendMessage(messageService services.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: envoyer un message privé
	}
}
