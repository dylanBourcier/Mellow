package models

import (
	"github.com/google/uuid"
	"time"
)

type GroupJoinRequest struct {
	ID          uuid.UUID  `json:"id"`
	GroupID     uuid.UUID  `json:"group_id"`
	RequesterID uuid.UUID  `json:"requester_id"`
	Status      string     `json:"status"` // pending, accepted, rejected, cancelled
	CreatedAt   time.Time  `json:"created_at"`
	DecidedAt   *time.Time `json:"decided_at,omitempty"`
	DecidedBy   *uuid.UUID `json:"decided_by,omitempty"`

	// Optional denormalized presenter fields for listing
	RequesterUsername *string `json:"requester_username,omitempty"`
	RequesterAvatar   *string `json:"requester_avatar_url,omitempty"`
}
