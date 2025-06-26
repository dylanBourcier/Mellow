package models

import "github.com/google/uuid"

type PostViewer struct {
	PostID uuid.UUID `json:"post_id"`
	UserID uuid.UUID `json:"user_id"`
}
