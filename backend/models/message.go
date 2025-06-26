package models

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	MessageID    uuid.UUID `json:"message_id"`
	SenderID     uuid.UUID `json:"sender_id"`
	ReceiverID   uuid.UUID `json:"receiver_id"`
	Content      *string   `json:"content,omitempty"`
	CreationDate time.Time `json:"creation_date"`
}
