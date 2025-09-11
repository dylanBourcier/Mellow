package messages

import (
	"mellow/services"
	"mellow/utils"
	"net/http"
)

func GetGroupMessages(service services.MessageService, groupId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		page := 1
		pageSize := 20
		msgs, err := service.GetGroupConversation(r.Context(), uid.String(), groupId, page, pageSize)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get conversation", err)
			return
		}

		// Mark the entire conversation as read for the current user
		if err := service.MarkAsReadConversation(r.Context(), uid.String(), groupId); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to mark conversation as read", err)
			return
		}
		for _, msg := range msgs {
			// You can process each msg here if needed
			msg.SenderImageUrl = utils.GetFullImageURLAvatar(msg.SenderImageUrl)
		}

		utils.RespondJSON(w, http.StatusOK, "Conversation retrieved", msgs)
	}
}

func SendGroupMessage(w http.ResponseWriter, r *http.Request, groupId string) {
	// TODO: envoyer un message au groupe
}
