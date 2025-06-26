package models

import (
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	NotificationID uuid.UUID `gorm:"type:char(36);primaryKey" json:"notification_id"`
	UserID         uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Type           string    `gorm:"type:text;not null" json:"type"` // follow, group_invite, event_created
	Seen           bool      `gorm:"default:false" json:"seen"`
	CreationDate   time.Time `gorm:"autoCreateTime" json:"creation_date"`
}
