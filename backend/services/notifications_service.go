package services

import (
	"context"
	"mellow/models"
)

type NotificationService interface {
	// CreateNotification enregistre une nouvelle notification à destination d'un utilisateur.
	CreateNotification(ctx context.Context, notif *models.Notification) error

	// GetUserNotifications récupère toutes les notifications d’un utilisateur.
	GetUserNotifications(ctx context.Context, userID string) ([]*models.Notification, error)

	// MarkAsRead marque une notification comme lue.
	MarkAsRead(ctx context.Context, notificationID, userID string) error

	// DeleteNotification supprime une notification (par l’utilisateur ou le système).
	DeleteNotification(ctx context.Context, notificationID string) error
}
