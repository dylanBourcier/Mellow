package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	EventID      uuid.UUID `json:"event_id"`
	UserID       uuid.UUID `json:"user_id"`
	GroupID      uuid.UUID `json:"group_id"`
	CreationDate time.Time `json:"creation_date"`
	EventDate    time.Time `json:"event_date"`
	Title        string    `json:"title"`
}
