package servimpl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/services"
)

type messageServiceImpl struct {
	db *sql.DB
}

// NewMessageService crée une nouvelle instance de MessageService.
func NewMessageService(db *sql.DB) services.MessageService {
	return &messageServiceImpl{db: db}
}

func (s *messageServiceImpl) SendMessage(ctx context.Context, msg *models.Message) error {
	// TODO: Vérifier que les deux utilisateurs existent et peuvent se parler
	// TODO: Appeler le repository pour enregistrer le message
	return nil
}

func (s *messageServiceImpl) GetConversation(ctx context.Context, user1ID, user2ID string, page, pageSize int) ([]*models.Message, error) {
	// TODO: Vérifier droits d’accès
	// TODO: Appeler le repository pour récupérer les messages paginés entre deux utilisateurs
	return nil, nil
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
