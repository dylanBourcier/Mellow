package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	MessageID      uuid.UUID `json:"message_id"`
	SenderID       uuid.UUID `json:"sender_id"`
	ReceiverID     uuid.UUID `json:"receiver_id"`
	Content        *string   `json:"content,omitempty"`
	CreationDate   time.Time `json:"creation_date"`
	IsRead         bool      `json:"is_read"`
	SenderUsername *string   `json:"username,omitempty"`
	SenderImageUrl *string   `json:"image_url,omitempty"`
}

type ConversationPreview struct {
	UserID      uuid.UUID `json:"user_id"`      // L’autre utilisateur
	Avatar      *string   `json:"avatar"`       // Avatar de l’utilisateur
	Username    string    `json:"username"`     // Pour affichage dans l’UI
	LastMessage string    `json:"last_message"` // Dernier message échangé
	LastSentAt  time.Time `json:"last_sent_at"` // Date d’envoi du dernier message
	UnreadCount int       `json:"unread_count"` // Nb de messages non lus
}
