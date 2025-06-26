package models

import "github.com/google/uuid"

type Report struct {
	PostID  uuid.UUID `json:"post_id"`
	UserID  uuid.UUID `json:"user_id"`
	Content *string   `json:"content,omitempty"`
	Type    string    `json:"type"` // spam, abuse, other
}
