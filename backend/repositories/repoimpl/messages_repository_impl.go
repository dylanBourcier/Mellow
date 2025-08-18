package repoimpl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/repositories"
)

type messageRepositoryImpl struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) repositories.MessageRepository {
	return &messageRepositoryImpl{db: db}
}

func (r *messageRepositoryImpl) InsertMessage(ctx context.Context, message *models.Message) error {
	// TODO: INSERT INTO messages (id, sender_id, receiver_id, content, creation_date, is_read) VALUES (?, ?, ?, ?, ?, ?)
	return nil
}

func (r *messageRepositoryImpl) GetConversation(ctx context.Context, user1ID, user2ID string, offset, limit int) ([]*models.Message, error) {
	// TODO: SELECT * FROM messages WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
	// ORDER BY creation_date ASC LIMIT ? OFFSET ?
	query := `
		SELECT * FROM messages
		WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
		ORDER BY creation_date ASC
		LIMIT ? OFFSET ?
	`
	rows, err := r.db.QueryContext(ctx, query, user1ID, user2ID, user2ID, user1ID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.MessageID, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.CreationDate); err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}
	return messages, nil
}

func (r *messageRepositoryImpl) DeleteMessage(ctx context.Context, messageID string) error {
	// TODO: DELETE FROM messages WHERE id = ?
	return nil
}

func (r *messageRepositoryImpl) GetRecentConversations(ctx context.Context, userID string) ([]*models.ConversationPreview, error) {
	// TODO: Requête pour récupérer les dernières conversations (dernier message par interlocuteur)
	return nil, nil
}

func (r *messageRepositoryImpl) MarkAsRead(ctx context.Context, messageID, userID string) error {
	// TODO: UPDATE messages SET is_read = true WHERE id = ? AND receiver_id = ?
	return nil
}
