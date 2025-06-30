package repositories

import (
	"context"
	"mellow/models"
)

type MessageRepository interface {
	// InsertMessage enregistre un nouveau message.
	InsertMessage(ctx context.Context, message *models.Message) error

	// GetConversation retourne l’historique paginé entre deux utilisateurs.
	GetConversation(ctx context.Context, user1ID, user2ID string, offset, limit int) ([]*models.Message, error)

	// DeleteMessage supprime un message par son ID.
	DeleteMessage(ctx context.Context, messageID string) error

	// GetRecentConversations retourne les dernières conversations de l’utilisateur.
	GetRecentConversations(ctx context.Context, userID string) ([]*models.ConversationPreview, error)

	// MarkAsRead met à jour le statut de lecture d’un message.
	MarkAsRead(ctx context.Context, messageID, userID string) error
}
