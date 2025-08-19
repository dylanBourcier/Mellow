package models

import (
	"time"

	"github.com/google/uuid"
)

type FollowRequest struct {
	RequestID    uuid.UUID  `json:"request_id"`
	SenderID     uuid.UUID  `json:"sender_id"`
	ReceiverID   uuid.UUID  `json:"receiver_id"`
	GroupID      *uuid.UUID `json:"group_id,omitempty"`
	Status       *bool      `json:"status,omitempty"`
	CreationDate time.Time  `json:"creation_date"`
	Type         string     `json:"type"` // user ou group
}
