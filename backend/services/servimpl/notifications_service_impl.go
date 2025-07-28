package servimpl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"mellow/models"
	"mellow/repositories"
	"mellow/services"
	"mellow/utils"
	"time"
)

type notificationServiceImpl struct {
	notifRepo repositories.NotificationRepository
	userRepo  repositories.UserRepository
}

// NewNotificationService crée une nouvelle instance de NotificationService.
func NewNotificationService(notifRepo repositories.NotificationRepository, userRepo repositories.UserRepository) services.NotificationService {
	return &notificationServiceImpl{notifRepo: notifRepo, userRepo: userRepo}
}

func (s *notificationServiceImpl) CreateNotification(ctx context.Context, notif *models.Notification) error {
	if notif == nil || notif.UserID == uuid.Nil || notif.Type == "" {
		return utils.ErrInvalidPayload
	}

	switch notif.Type {
	case "follow", "group_invite", "event_created":
	default:
		return utils.ErrInvalidPayload
	}

	// verify user exists
	user, err := s.userRepo.FindUserByID(ctx, notif.UserID.String())
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if user == nil {
		return utils.ErrUserNotFound
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return utils.ErrUUIDGeneration
	}
	notif.NotificationID = id
	notif.Seen = false
	notif.CreationDate = time.Now()

	if err := s.notifRepo.InsertNotification(ctx, notif); err != nil {
		return err
	}
	return nil
}

func (s *notificationServiceImpl) GetUserNotifications(ctx context.Context, userID string) ([]*models.Notification, error) {
	if userID == "" {
		return nil, utils.ErrInvalidPayload
	}

	notifs, err := s.notifRepo.GetUserNotifications(ctx, userID)
	if err != nil {
		return nil, err
	}

	return notifs, nil
}

func (s *notificationServiceImpl) MarkAsRead(ctx context.Context, notificationID, userID string) error {
	if notificationID == "" || userID == "" {
		return utils.ErrInvalidPayload
	}

	notif, err := s.notifRepo.GetNotificationByID(ctx, notificationID)
	if err != nil {
		return err
	}
	if notif == nil {
		return utils.ErrNotificationNotFound
	}
	if notif.UserID.String() != userID {
		return utils.ErrForbidden
	}

	return s.notifRepo.MarkAsRead(ctx, notificationID)
}

func (s *notificationServiceImpl) DeleteNotification(ctx context.Context, notificationID string) error {
	// TODO: Vérifier les droits (self ou admin)
	// TODO: Appeler le repository pour supprimer la notification
	return nil
}
