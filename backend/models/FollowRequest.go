package models

import (
	"github.com/google/uuid"
	"time"
)

type FollowRequest struct {
	SenderID     uuid.UUID  `json:"sender_id"`
	ReceiverID   uuid.UUID  `json:"receiver_id"`
	GroupID      *uuid.UUID `json:"group_id,omitempty"`
	Status       *bool      `json:"status,omitempty"`
	CreationDate time.Time  `json:"creation_date"`
	Type         string     `json:"type"` // user ou group
}
