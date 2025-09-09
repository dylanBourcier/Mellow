package servimpl

import (
	"context"
	"github.com/google/uuid"
	"mellow/models"
	"mellow/repositories"
	"mellow/services"
	"mellow/utils"
	"time"
)

type messageServiceImpl struct {
	repo repositories.MessageRepository
}

// NewMessageService crée une nouvelle instance de MessageService.
func NewMessageService(repo repositories.MessageRepository) services.MessageService {
	return &messageServiceImpl{repo: repo}
}

// SendMessage envoie un messgae depuis un "sender" vers un "receiver" et le persistent dans le repository.
func (s *messageServiceImpl) SendMessage(ctx context.Context, msg *models.Message) error {
	if msg == nil || msg.Content == nil || *msg.Content == "" {
		return utils.ErrInvalidPayload
	}
	if msg.SenderID == uuid.Nil || msg.ReceiverID == uuid.Nil {
		return utils.ErrInvalidUserData
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return utils.ErrUUIDGeneration
	}

	msg.MessageID = id
	msg.CreationDate = time.Now()
	msg.IsRead = false
	return s.repo.InsertMessage(ctx, msg)
}

func (s *messageServiceImpl) GetConversation(ctx context.Context, user1ID, user2ID string, page, pageSize int) ([]*models.Message, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 30
	}
	if user1ID == "" || user2ID == "" {
		return nil, utils.ErrInvalidPayload
	}

	offset := (page - 1) * pageSize
	msgs, err := s.repo.GetConversation(ctx, user1ID, user2ID, offset, pageSize)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (s *messageServiceImpl) DeleteMessage(ctx context.Context, messageID, requesterID string) error {
	if messageID == "" || requesterID == "" {
		return utils.ErrInvalidPayload
	}
	return s.repo.DeleteMessage(ctx, messageID)
}

func (s *messageServiceImpl) GetRecentConversations(ctx context.Context, userID string) ([]*models.ConversationPreview, error) {
	if userID == "" {
		return nil, utils.ErrInvalidPayload
	}

	conversations, err := s.repo.GetRecentConversations(ctx, userID)
	if err != nil {
		return nil, err
	}
	return conversations, nil
}

func (s *messageServiceImpl) MarkAsRead(ctx context.Context, messageID, userID string) error {
	if messageID == "" || userID == "" {
		return utils.ErrInvalidPayload
	}

	return s.repo.MarkAsRead(ctx, messageID, userID)
}
