package servimpl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/services"
)

type notificationServiceImpl struct {
	db *sql.DB
}

// NewNotificationService crée une nouvelle instance de NotificationService.
func NewNotificationService(db *sql.DB) services.NotificationService {
	return &notificationServiceImpl{db: db}
}

func (s *notificationServiceImpl) CreateNotification(ctx context.Context, notif *models.Notification) error {
	// TODO: Vérifier que le destinataire existe
	// TODO: Appeler le repository pour insérer la notification
	return nil
}

func (s *notificationServiceImpl) GetUserNotifications(ctx context.Context, userID string) ([]*models.Notification, error) {
	// TODO: Appeler le repository pour récupérer les notifications d’un utilisateur
	return nil, nil
}

func (s *notificationServiceImpl) MarkAsRead(ctx context.Context, notificationID string) error {
	// TODO: Appeler le repository pour mettre à jour le statut de lecture
	return nil
}

func (s *notificationServiceImpl) DeleteNotification(ctx context.Context, notificationID string) error {
	// TODO: Vérifier les droits (self ou admin)
	// TODO: Appeler le repository pour supprimer la notification
	return nil
}
