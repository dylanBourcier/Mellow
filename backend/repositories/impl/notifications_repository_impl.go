package impl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/repositories"
)

type notificationRepositoryImpl struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) repositories.NotificationRepository {
	return &notificationRepositoryImpl{db: db}
}

func (r *notificationRepositoryImpl) InsertNotification(ctx context.Context, notif *models.Notification) error {
	// TODO: INSERT INTO notifications (id, user_id, message, is_read, created_at) VALUES (?, ?, ?, ?, ?)
	return nil
}

func (r *notificationRepositoryImpl) GetUserNotifications(ctx context.Context, userID string) ([]*models.Notification, error) {
	// TODO: SELECT * FROM notifications WHERE user_id = ? ORDER BY created_at DESC
	return nil, nil
}

func (r *notificationRepositoryImpl) MarkAsRead(ctx context.Context, notificationID string) error {
	// TODO: UPDATE notifications SET is_read = true WHERE id = ?
	return nil
}

func (r *notificationRepositoryImpl) DeleteNotification(ctx context.Context, notificationID string) error {
	// TODO: DELETE FROM notifications WHERE id = ?
	return nil
}
