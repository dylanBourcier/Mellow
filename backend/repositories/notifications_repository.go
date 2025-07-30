package repositories

import (
	"context"
	"mellow/models"
)

type NotificationRepository interface {
	// InsertNotification ajoute une nouvelle notification.
	InsertNotification(ctx context.Context, notif *models.Notification) error

	// GetUserNotifications récupère toutes les notifications d’un utilisateur.
	GetUserNotifications(ctx context.Context, userID string) ([]*models.Notification, error)

	// GetNotificationByID récupère une notification par son ID.
	GetNotificationByID(ctx context.Context, notificationID string) (*models.Notification, error)

	// GetNotificationByTypeSenderReceiver récupère une notification par son type, l'ID de l'expéditeur et l'ID du destinataire.
	GetNotificationByTypeSenderReceiver(ctx context.Context, notifType, senderID, receiverID string) (*models.Notification, error)

	// MarkAsRead met à jour le statut d’une notification.
	MarkAsRead(ctx context.Context, notificationID string) error

	// DeleteNotification supprime une notification.
	DeleteNotification(ctx context.Context, notificationID string) error
}
