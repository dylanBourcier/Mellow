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
	query := `SELECT notification_id, user_id, type, seen, creation_date
                  FROM notifications
                  WHERE user_id = ?
                  ORDER BY creation_date DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query notifications: %w", err)
	}
	defer rows.Close()

	var notifs []*models.Notification
	for rows.Next() {
		var n models.Notification
		if err := rows.Scan(&n.NotificationID, &n.UserID, &n.Type, &n.Seen, &n.CreationDate); err != nil {
			return nil, fmt.Errorf("failed to scan notification: %w", err)
		}
		notifs = append(notifs, &n)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return notifs, nil
}

func (r *notificationRepositoryImpl) GetNotificationByID(ctx context.Context, notificationID string) (*models.Notification, error) {
	query := `SELECT notification_id, user_id, type, seen, creation_date FROM notifications WHERE notification_id = ?`
	var n models.Notification
	err := r.db.QueryRowContext(ctx, query, notificationID).Scan(&n.NotificationID, &n.UserID, &n.Type, &n.Seen, &n.CreationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get notification: %w", err)
	}
	return &n, nil
}

func (r *notificationRepositoryImpl) MarkAsRead(ctx context.Context, notificationID string) error {
	query := `UPDATE notifications SET seen = 1 WHERE notification_id = ?`
	_, err := r.db.ExecContext(ctx, query, notificationID)
	if err != nil {
		return fmt.Errorf("failed to mark notification as read: %w", err)
	}
	return nil
}
func (r *notificationRepositoryImpl) DeleteNotification(ctx context.Context, notificationID string) error {
	// TODO: DELETE FROM notifications WHERE id = ?
	return nil
}
