package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	NotificationID  uuid.UUID  `gorm:"type:char(36);primaryKey" json:"notification_id"`
	UserID          uuid.UUID  `gorm:"type:char(36);not null" json:"user_id"`
	RequestID       *uuid.UUID `gorm:"type:char(36)" json:"request_id,omitempty"` // Optional, for group invites, and follow requests
	Type            string     `gorm:"type:text;not null" json:"type"`            // follow, group_invite, event_created
	Seen            bool       `gorm:"default:false" json:"seen"`
	CreationDate    time.Time  `gorm:"autoCreateTime" json:"creation_date"`
	SenderID        string     `gorm:"not null" json:"sender_id"`               // Required
	SenderUsername  *string    `gorm:"-" json:"sender_username,omitempty"`      // Optional, for follow and group requests
	SenderAvatarURL *string    `gorm:"-" json:"sender_avatar_url,omitempty"`    // Optional, for follow and group requests
	GroupID         *uuid.UUID `gorm:"type:char(36)" json:"group_id,omitempty"` // Optional, for group invites
	GroupName       *string    `gorm:"-" json:"group_name,omitempty"`           // Optional, for group invites
}

type CreateNotificationPayload struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}

const (
	NotificationTypeNewFollower  = "new_follower"
	NotificationTypeEventCreated = "event_created"

	NotificationTypeFollowRequest         = "follow_request"
	NotificationTypeAcceptedFollowRequest = "accepted_follow_request"
	NotificationTypeRejectedFollowRequest = "rejected_follow_request"

	NotificationTypeGroupInvite         = "group_invite"
	NotificationTypeAcceptedGroupInvite = "accepted_group_invite"
	NotificationTypeRejectedGroupInvite = "rejected_group_invite"

	NotificationTypeGroupRequest         = "group_request"
	NotificationTypeAcceptedGroupRequest = "accepted_group_request"
	NotificationTypeRejectedGroupRequest = "rejected_group_request"
)
