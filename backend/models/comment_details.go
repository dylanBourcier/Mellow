package models

import "time"

type CommentDetails struct {
	CommentID    string    `json:"comment_id"`
	PostID       string    `json:"post_id"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"creation_date"`
	ImageURL     *string   `json:"image_url,omitempty"`
	UserID       string    `json:"user_id"`
	Username     string    `json:"username,omitempty"`
	AvatarURL    *string   `json:"avatar_url,omitempty"`
}
