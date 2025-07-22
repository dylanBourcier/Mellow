package models

import (
	"time"

	"github.com/google/uuid"
)

type PostDetails struct {
	PostID        uuid.UUID  `json:"post_id"`
	GroupID       *uuid.UUID `json:"group_id,omitempty"` // nullable
	UserID        uuid.UUID  `json:"user_id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	CreationDate  time.Time  `json:"creation_date"`
	Visibility    string     `json:"visibility"`
	ImageURL      *string    `json:"image_url,omitempty"` // nullable
	Username      string     `json:"username"`
	AvatarURL     *string    `json:"avatar_url,omitempty"` // nullable
	CommentsCount int        `json:"comments_count"`
	GroupName     *string    `json:"group_name,omitempty"` // nullable
}
