package models

import (
	"github.com/google/uuid"
	"time"
)

type GroupMember struct {
	GroupID  uuid.UUID  `json:"group_id"`
	UserID   uuid.UUID  `json:"user_id"`
	Role     *string    `json:"role,omitempty"`
	JoinDate *time.Time `json:"join_date,omitempty"`
}
