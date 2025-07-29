package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `json:"user_id"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Username     string    `json:"username"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	Birthdate    time.Time `json:"birthdate"`
	Role         string    `json:"role"`
	ImageURL     *string   `json:"image_url,omitempty"`
	CreationDate time.Time `json:"creation_date"`
	Description  *string   `json:"description,omitempty"`
	Privacy      string    `json:"privacy"` // "public", "private"
}

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"` // email ou username
	Password   string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Username    *string    `json:"username,omitempty"`
	Password    *string    `json:"password,omitempty"`
	Firstname   *string    `json:"firstname,omitempty"`
	Lastname    *string    `json:"lastname,omitempty"`
	Birthdate   *time.Time `json:"birthdate,omitempty"`
	ImageURL    *string    `json:"image_url,omitempty"`
	Description *string    `json:"description,omitempty"`
}

type ReportRequest struct {
	Type     string `json:"type"`
	TargetID string `json:"target_id"`
	Content  string `json:"content"`
}

type UserProfileData struct {
	UserID         string    `json:"user_id"`
	Username       string    `json:"username"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	Email          string    `json:"email"`
	Birthdate      time.Time `json:"birthdate"`
	ImageURL       *string   `json:"image_url,omitempty"`
	Privacy        string    `json:"privacy"` // "public", "private"
	FollowersCount int       `json:"followers_count"`
	FollowedCount  int       `json:"followed_count"`
	FollowStatus   string    `json:"follow_status"` // "not_follow", "follows", "requested"
	Description    *string   `json:"description,omitempty"`
	Posts          []Post    `json:"posts"`
}
