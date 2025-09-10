package messages

import (
	"encoding/json"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"net/http"

	"github.com/google/uuid"
)

func GetConversation(service services.MessageService, userId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		page := 1
		pageSize := 20
		msgs, err := service.GetConversation(r.Context(), uid.String(), userId, page, pageSize)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get conversation", err)
			return
		}

		// Mark the entire conversation as read for the current user
		if err := service.MarkAsReadConversation(r.Context(), uid.String(), userId); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to mark conversation as read", err)
			return
		}

		utils.RespondJSON(w, http.StatusOK, "Conversation retrieved", msgs)
	}
}

func SendMessage(service services.MessageService, userId string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		var payload struct {
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid JSON", utils.ErrInvalidJSON)
			return
		}

		if payload.Content == "" {
			utils.RespondError(w, http.StatusBadRequest, "Empty content", utils.ErrInvalidPayload)
			return
		}

		receiverID, err := uuid.Parse(userId)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid user", utils.ErrInvalidPayload)
			return
		}

		content := payload.Content
		msg := models.Message{
			SenderID:   uid,
			ReceiverID: receiverID,
			Content:    &content,
		}
		if err := service.SendMessage(r.Context(), &msg); err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to send message", err)
			return
		}
		utils.RespondJSON(w, http.StatusCreated, "Message sent", msg)
	}
}

func GetRecentConversations(service services.MessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uid, err := utils.GetUserIDFromContext(r.Context())
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Unauthorized", utils.ErrUnauthorized)
			return
		}

		msgs, err := service.GetRecentConversations(r.Context(), uid.String())
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get recent conversations", err)
			return
		}
		for _, conversation := range msgs {
			if conversation.Avatar != nil {
				conversation.Avatar = utils.GetFullImageURLAvatar(conversation.Avatar)
			}
		}
		utils.RespondJSON(w, http.StatusOK, "Recent conversations retrieved", msgs)
	}
}
