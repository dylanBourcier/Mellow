package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"mellow/models"
	"mellow/repositories"
)

type messageRepositoryImpl struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) repositories.MessageRepository {
	return &messageRepositoryImpl{db: db}
}

func (r *messageRepositoryImpl) InsertMessage(ctx context.Context, message *models.Message) (string, error) {
	query := `insert into messages (message_id, sender_id, receiver_id, content, creation_date, is_read) values (?,?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, message.MessageID, message.SenderID, message.ReceiverID, message.Content, message.CreationDate, message.IsRead)
	if err != nil {
		return "", fmt.Errorf("failed to insert message: %w", err)
	}
	return message.MessageID.String(), nil
}

func (r *messageRepositoryImpl) GetConversation(ctx context.Context, user1ID, user2ID string, offset, limit int) ([]*models.Message, error) {
	query := `select message_id, sender_id, receiver_id, content, creation_date, is_read
			from messages
			where (sender_id = ? and receiver_id = ?) or (sender_id = ? and receiver_id = ?)
			order by creation_date asc
			limit ? offset ?`

	rows, err := r.db.QueryContext(ctx, query, user1ID, user2ID, user2ID, user1ID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get conversation: %w", err)
	}
	defer rows.Close()
	var messages []*models.Message
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.MessageID, &m.SenderID, &m.ReceiverID, &m.Content, &m.CreationDate, &m.IsRead); err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		messages = append(messages, &m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over messages: %w", err)
	}
	return messages, nil
}

func (r *messageRepositoryImpl) GetGroupConversation(ctx context.Context, user1ID, groupID string, offset, limit int) ([]*models.Message, error) {
	query := `select m.message_id, m.sender_id, m.receiver_id, m.content, m.creation_date, m.is_read,u.username,u.image_Url
			from messages m
			join users u on m.sender_id = u.user_id
			where m.receiver_id = ? 
			order by m.creation_date asc
			limit ? offset ?`

	rows, err := r.db.QueryContext(ctx, query, groupID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get conversation: %w", err)
	}
	defer rows.Close()
	var messages []*models.Message
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.MessageID, &m.SenderID, &m.ReceiverID, &m.Content, &m.CreationDate, &m.IsRead, &m.SenderUsername, &m.SenderImageUrl); err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		messages = append(messages, &m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over messages: %w", err)
	}
	return messages, nil
}

func (r *messageRepositoryImpl) DeleteMessage(ctx context.Context, messageID string) error {
	_, err := r.db.ExecContext(ctx, `delete from messages where message_id = ?`, messageID)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}
	return nil
}

func (r *messageRepositoryImpl) GetRecentConversations(ctx context.Context, userID string) ([]*models.ConversationPreview, error) {
	query := `SELECT u.user_id, u.image_url, u.username, m.content, m.creation_date,
                        (SELECT COUNT(*) FROM messages m2 WHERE m2.sender_id = u.user_id AND m2.receiver_id = ? AND m2.is_read = 0) AS unread
                FROM (
                        SELECT CASE WHEN sender_id = ? THEN receiver_id ELSE sender_id END AS other_id,
                               MAX(creation_date) AS last_date
                        FROM messages
                        WHERE sender_id = ? OR receiver_id = ?
                        GROUP BY other_id
                ) conv
                JOIN users u ON u.user_id = conv.other_id
                JOIN messages m ON ((m.sender_id = ? AND m.receiver_id = u.user_id) OR (m.receiver_id = ? AND m.sender_id = u.user_id)) AND m.creation_date = conv.last_date
                ORDER BY conv.last_date DESC`
	rows, err := r.db.QueryContext(ctx, query, userID, userID, userID, userID, userID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query recent conversations: %w", err)
	}
	defer rows.Close()
	var convs []*models.ConversationPreview
	for rows.Next() {
		var cp models.ConversationPreview
		if err := rows.Scan(&cp.UserID, &cp.Avatar, &cp.Username, &cp.LastMessage, &cp.LastSentAt, &cp.UnreadCount); err != nil {
			return nil, fmt.Errorf("failed to scan conversation preview: %w", err)
		}
		convs = append(convs, &cp)
		fmt.Println(cp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return convs, nil
}

func (r *messageRepositoryImpl) MarkAsRead(ctx context.Context, messageID, userID string) error {
	_, err := r.db.ExecContext(ctx, `update messages set is_read = 1 where message_id = ? and receiver_id = ?`, messageID, userID)
	if err != nil {
		return fmt.Errorf("failed to mark message as read: %w", err)
	}
	return nil
}
func (r *messageRepositoryImpl) MarkAsReadConversation(ctx context.Context, userId, otherId string) error {
	_, err := r.db.ExecContext(ctx, `update messages set is_read = 1 where sender_id = ? and receiver_id = ? and is_read = 0`, otherId, userId)
	if err != nil {
		return fmt.Errorf("failed to mark messages as read: %w", err)
	}
	return nil
}
