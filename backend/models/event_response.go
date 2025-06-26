package models

import "github.com/google/uuid"

type EventResponse struct {
	EventID uuid.UUID `json:"event_id"`
	UserID  uuid.UUID `json:"user_id"`
	Status  *string   `json:"status,omitempty"` // going, not_going
	Vote    *string   `json:"vote,omitempty"`
}
