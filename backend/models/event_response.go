package models

import "github.com/google/uuid"

type EventResponse struct {
	EventID uuid.UUID `json:"event_id"`
	UserID  uuid.UUID `json:"user_id"`
	GroupID uuid.UUID `json:"group_id"`
	Vote    bool      `json:"vote"`
}
