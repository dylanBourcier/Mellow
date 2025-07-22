package models

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	GroupID      uuid.UUID `json:"group_id"`
	UserID       uuid.UUID `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description,omitempty"`
	CreationDate time.Time `json:"creation_date"`
}

type GroupEditPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
