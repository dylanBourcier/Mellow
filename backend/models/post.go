package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	PostID       uuid.UUID  `json:"post_id"`
	GroupID      *uuid.UUID `json:"group_id,omitempty"` // nullable
	UserID       uuid.UUID  `json:"user_id"`
	Title        string     `json:"title"`
	Content      string     `json:"content"`
	CreationDate time.Time  `json:"creation_date"`
	Visibility   string     `json:"visibility"`
	ImageURL     *string    `json:"image_url,omitempty"` // nullable
}

type UpdatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
