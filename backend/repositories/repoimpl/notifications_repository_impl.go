package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"mellow/models"
	"mellow/repositories"

	"github.com/google/uuid"
)

type notificationRepositoryImpl struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) repositories.NotificationRepository {
	return &notificationRepositoryImpl{db: db}
}

func (r *notificationRepositoryImpl) InsertNotification(ctx context.Context, notif *models.Notification) error {
	query := `INSERT INTO notifications (notification_id, user_id, sender_id,request_id, type, seen, creation_date) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, notif.NotificationID, notif.UserID, notif.SenderID, notif.RequestID, notif.Type, notif.Seen, notif.CreationDate)
	if err != nil {
		return fmt.Errorf("failed to insert notification: %w", err)
	}
	return nil
}

func (r *notificationRepositoryImpl) GetUserNotifications(ctx context.Context, userID string) ([]*models.Notification, error) {
	query := `SELECT n.notification_id, n.request_id, n.user_id, n.type, n.seen, n.creation_date, 
					 COALESCE(u.username, '') AS sender_username,
					 COALESCE(u.image_url, '') AS sender_avatar_url,
					 n.sender_id,
					 COALESCE(fr.group_id, '') AS group_id,
					 COALESCE(g.title, '') AS group_title
			  FROM notifications n
			  LEFT JOIN users u ON n.sender_id = u.user_id
			  LEFT JOIN follow_requests fr ON n.request_id = fr.request_id
			  LEFT JOIN groups g ON fr.group_id = g.group_id
			  WHERE n.user_id = ?
			  ORDER BY n.creation_date ASC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query notifications: %w", err)
	}
	defer rows.Close()

	var notifs []*models.Notification
	for rows.Next() {
		var n models.Notification
		var senderUsername, senderAvatarURL, groupID, groupName string
		if err := rows.Scan(&n.NotificationID, &n.RequestID, &n.UserID, &n.Type, &n.Seen, &n.CreationDate, &senderUsername, &senderAvatarURL, &n.SenderID, &groupID, &groupName); err != nil {
			return nil, fmt.Errorf("failed to scan notification: %w", err)
		}
		n.SenderUsername = &senderUsername   // Assuming Notification model has a SenderUsername field
		n.SenderAvatarURL = &senderAvatarURL // Assuming Notification model has a SenderAvatarURL field
		if groupID != "" {
			parsedGroupID, err := uuid.Parse(groupID)
			if err != nil {
				return nil, fmt.Errorf("failed to parse groupID as UUID: %w", err)
			}
			n.GroupID = &parsedGroupID // Assuming Notification model has a GroupID field
			n.GroupName = &groupName   // Assuming Notification model has a GroupName field
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

func (r *notificationRepositoryImpl) GetNotificationByTypeSenderReceiver(ctx context.Context, notifType, senderID, receiverID string) (*models.Notification, error) {
	query := `SELECT notification_id, user_id, type, seen, creation_date FROM notifications 
			  WHERE type = ? AND sender_id = ? AND user_id = ?`
	var n models.Notification
	err := r.db.QueryRowContext(ctx, query, notifType, senderID, receiverID).Scan(&n.NotificationID, &n.UserID, &n.Type, &n.Seen, &n.CreationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get notification by type and sender: %w", err)
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
	query := `DELETE FROM notifications WHERE notification_id = ?`
	_, err := r.db.ExecContext(ctx, query, notificationID)
	if err != nil {
		return fmt.Errorf("failed to delete notification: %w", err)
	}
	// Successfully deleted the notification
	return nil
}
