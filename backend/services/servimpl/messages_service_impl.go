package servimpl

import (
	"context"
	"mellow/models"
	"mellow/repositories"
	"mellow/services"
)

type messageServiceImpl struct {
	messageRepository repositories.MessageRepository
}

// NewMessageService crée une nouvelle instance de MessageService.
func NewMessageService(messageRepository repositories.MessageRepository) services.MessageService {
	return &messageServiceImpl{messageRepository: messageRepository}
}

func (s *messageServiceImpl) SendMessage(ctx context.Context, msg *models.Message) error {
	// TODO: Vérifier que les deux utilisateurs existent et peuvent se parler
	// TODO: Appeler le repository pour enregistrer le message
	return nil
}

func (s *messageServiceImpl) GetConversation(ctx context.Context, user1ID, user2ID string, limit, offset int) ([]*models.Message, error) {
	if limit == 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	// TODO: Appeler le repository pour récupérer les messages paginés entre deux utilisateurs
	return s.messageRepository.GetConversation(ctx, user1ID, user2ID, limit, offset)
}

func (s *messageServiceImpl) DeleteMessage(ctx context.Context, messageID, requesterID string) error {
	// TODO: Vérifier si le requester est l’auteur ou un modérateur
	// TODO: Supprimer le message via le repository
	return nil
}

func (s *messageServiceImpl) GetRecentConversations(ctx context.Context, userID string) ([]*models.ConversationPreview, error) {
	// TODO: Récupérer les dernières conversations par utilisateur unique
	return nil, nil
}

func (s *messageServiceImpl) MarkAsRead(ctx context.Context, messageID, userID string) error {
	// TODO: Vérifier que l’utilisateur est le destinataire
	// TODO: Mettre à jour le statut via le repository
	return nil
}
