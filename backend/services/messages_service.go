package services

import (
	"context"
	"mellow/models"
)

type MessageService interface {
	// SendMessage envoie un message privé entre deux utilisateurs.
	SendMessage(ctx context.Context, msg *models.Message) error

	// GetConversation retourne l’historique paginé entre deux utilisateurs.
	GetConversation(ctx context.Context, user1ID, user2ID string, page, pageSize int) ([]*models.Message, error)

	// DeleteMessage permet à l’auteur ou à un modérateur de supprimer un message.
	DeleteMessage(ctx context.Context, messageID, requesterID string) error

	// GetRecentConversations retourne les dernières conversations d’un utilisateur (par interlocuteur).
	GetRecentConversations(ctx context.Context, userID string) ([]*models.ConversationPreview, error)

	// MarkAsRead marque un message comme lu.
	MarkAsRead(ctx context.Context, messageID, userID string) error

	// MarkAsReadConversation met à jour le statut de lecture des messages dans une conversation.
	MarkAsReadConversation(ctx context.Context, userId, otherId string) error
}
