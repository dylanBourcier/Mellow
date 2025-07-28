package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	NotificationID uuid.UUID  `gorm:"type:char(36);primaryKey" json:"notification_id"`
	UserID         uuid.UUID  `gorm:"type:char(36);not null" json:"user_id"`
	RequestID      *uuid.UUID `gorm:"type:char(36)" json:"request_id,omitempty"` // Optional, for group invites, and follow requests
	Type           string     `gorm:"type:text;not null" json:"type"`            // follow, group_invite, event_created
	Seen           bool       `gorm:"default:false" json:"seen"`
	CreationDate   time.Time  `gorm:"autoCreateTime" json:"creation_date"`
}

type CreateNotificationPayload struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}
