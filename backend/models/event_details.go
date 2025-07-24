package models

type EventDetails struct {
	EventID        string           `json:"event_id"`
	UserID         string           `json:"user_id"`
	GroupID        string           `json:"group_id"`
	CreationDate   string           `json:"creation_date"`
	EventDate      string           `json:"event_date"`
	Title          string           `json:"title"`
	Username       string           `json:"username"`
	AvatarURL      *string          `json:"avatar_url,omitempty"`
	EventResponses *[]EventResponse `json:"event_responses"`
}
