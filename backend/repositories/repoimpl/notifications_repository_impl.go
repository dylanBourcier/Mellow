package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
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
	query := `INSERT INTO notifications (notification_id, user_id, type, seen, creation_date) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, notif.NotificationID, notif.UserID, notif.Type, notif.Seen, notif.CreationDate)
	if err != nil {
		return fmt.Errorf("failed to insert notification: %w", err)
	}
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
