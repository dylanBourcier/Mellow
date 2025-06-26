package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CommentID    uuid.UUID `json:"comment_id"`
	UserID       uuid.UUID `json:"user_id"`
	PostID       uuid.UUID `json:"post_id"`
	Content      *string   `json:"content,omitempty"`
	CreationDate time.Time `json:"creation_date"`
}
