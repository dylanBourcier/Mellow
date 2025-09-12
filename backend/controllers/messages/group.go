package messages

import (
	"encoding/json"
	"mellow/models"
	"mellow/services"
	"mellow/utils"
	"mellow/websocket"
	"net/http"
	"time"

	"github.com/google/uuid"
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

func SendGroupMessage(service services.MessageService, userSvc services.UserService, groupId string) http.HandlerFunc {
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
		receiverID, err := uuid.Parse(groupId)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid group", utils.ErrInvalidPayload)
			return
		}

		content := payload.Content
		msg := models.Message{
			SenderID:   uid,
			ReceiverID: receiverID,
			Content:    &content,
		}
		msgId, err := service.SendMessage(r.Context(), &msg)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to send message", err)
			return
		}
		msg.MessageID, err = uuid.Parse(msgId)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to parse message ID", err)
			return
		}
		connectedUser, err := userSvc.GetUserByID(r.Context(), uid.String())
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to get user info", err)
			return
		}
		// 👉 Broadcast WS après insertion
		wsMsg := websocket.WSMessage{
			ID:             msg.MessageID.String(),
			SenderID:       msg.SenderID.String(),
			SenderUsername: &connectedUser.Username,
			SenderImageUrl: utils.GetFullImageURLAvatar(connectedUser.ImageURL),
			Content:        *msg.Content,
			Timestamp:      msg.CreationDate.Format(time.RFC3339),
			Room:           "group:" + msg.ReceiverID.String(), // room = destinataire
			Type:           "group",
		}
		websocket.BroadcastMessage(wsMsg.Room, wsMsg)

		utils.RespondJSON(w, http.StatusCreated, "Message sent", msg)
	}
}
