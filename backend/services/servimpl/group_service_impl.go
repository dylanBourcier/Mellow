package servimpl

import (
	"context"
	"fmt"
	"mellow/models"
	"mellow/repositories"
	"mellow/services"
	"mellow/utils"
	"time"

	"github.com/google/uuid"
)

type groupServiceImpl struct {
	groupRepo repositories.GroupRepository
	notifSvc  services.NotificationService
}

// NewGroupService crée une nouvelle instance de GroupService.
func NewGroupService(groupRepo repositories.GroupRepository, notifSvc services.NotificationService) services.GroupService {
	return &groupServiceImpl{groupRepo: groupRepo, notifSvc: notifSvc}
}

func (s *groupServiceImpl) CreateGroup(ctx context.Context, group *models.Group) error {
	if group == nil || group.Title == "" || group.UserID == uuid.Nil {
		return utils.ErrInvalidPayload
	}

	taken, err := s.groupRepo.IsTitleTaken(ctx, group.Title)
	if err != nil {
		return err
	}
	if taken {
		return utils.ErrGroupAlreadyExists
	}

	gid, err := uuid.NewRandom()
	if err != nil {
		return utils.ErrUUIDGeneration
	}
	group.GroupID = gid
	group.CreationDate = time.Now()

	if err := s.groupRepo.InsertGroup(ctx, group); err != nil {
		return err
	}

	if err := s.groupRepo.AddMember(ctx, gid.String(), group.UserID.String()); err != nil {
		return err
	}

	return nil
}

func (s *groupServiceImpl) UpdateGroup(ctx context.Context, groupID, requesterID, title string, description string) error {
	if groupID == "" || requesterID == "" || title == "" {
		return utils.ErrInvalidPayload
	}

	existing, err := s.groupRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return err
	}
	if existing == nil {
		return utils.ErrGroupNotFound
	}

	if existing.UserID.String() != requesterID {
		return utils.ErrForbidden
	}

	if len(title) > 100 {
		return utils.ErrContentTooLong
	}
	if len(title) < 1 {
		return utils.ErrContentTooShort
	}
	if description != "" && len(description) > 1000 {
		return utils.ErrContentTooLong
	}

	gid, err := uuid.Parse(groupID)
	if err != nil {
		return utils.ErrInvalidPayload
	}
	g := &models.Group{
		GroupID:     gid,
		Title:       title,
		Description: description,
	}

	return s.groupRepo.UpdateGroup(ctx, g)
}

func (s *groupServiceImpl) GetGroupByID(ctx context.Context, groupID string) (*models.Group, error) {
	if groupID == "" {
		return nil, utils.ErrInvalidPayload
	}
	group, err := s.groupRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	if group == nil {
		return nil, utils.ErrGroupNotFound
	}
	return group, nil
}

func (s *groupServiceImpl) GetAllGroups(ctx context.Context) ([]*models.Group, error) {
	groups, err := s.groupRepo.GetAllGroups(ctx)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (s *groupServiceImpl) GetAllGroupsWithoutUser(ctx context.Context, userID string) ([]*models.Group, error) {
	if userID == "" {
		return nil, utils.ErrInvalidPayload
	}
	groups, err := s.groupRepo.GetAllGroupsWithoutUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (s *groupServiceImpl) DeleteGroup(ctx context.Context, groupID, requesterID string) error {
	if groupID == "" || requesterID == "" {
		return utils.ErrInvalidPayload
	}

	group, err := s.groupRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return err
	}
	if group == nil {
		return utils.ErrGroupNotFound
	}
	if group.UserID.String() != requesterID {
		return utils.ErrUnauthorized
	}

	if err := s.groupRepo.DeleteGroup(ctx, groupID); err != nil {
		return err
	}
	return nil
}

func (s *groupServiceImpl) AddMember(ctx context.Context, groupID, userID string) error {
	if groupID == "" || userID == "" {
		return utils.ErrInvalidPayload
	}
	exists, err := s.groupRepo.IsMember(ctx, groupID, userID)
	if err != nil {
		return err
	}
	if exists {
		return utils.ErrResourceConflict
	}
	return s.groupRepo.AddMember(ctx, groupID, userID)
}

func (s *groupServiceImpl) RemoveMember(ctx context.Context, groupID, userID string) error {
	if groupID == "" || userID == "" {
		return utils.ErrInvalidPayload
	}

	group, err := s.groupRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return err
	}
	if group == nil {
		return utils.ErrGroupNotFound
	}

	isMember, err := s.groupRepo.IsMember(ctx, groupID, userID)
	if err != nil {
		return err
	}
	if !isMember {
		return utils.ErrForbidden
	}

	if group.UserID.String() == userID {
		return utils.ErrForbidden
	}

	return s.groupRepo.RemoveMember(ctx, groupID, userID)
}

func (s *groupServiceImpl) GetGroupMembers(ctx context.Context, groupID string) ([]*models.User, error) {
	if groupID == "" {
		return nil, utils.ErrInvalidPayload
	}

	group, err := s.groupRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, utils.ErrGroupNotFound
	}

	members, err := s.groupRepo.GetGroupMembers(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s *groupServiceImpl) IsMember(ctx context.Context, groupID, userID string) (bool, error) {
	return s.groupRepo.IsMember(ctx, groupID, userID)
}

func (s *groupServiceImpl) GetGroupsJoinedByUser(ctx context.Context, userID string) ([]*models.Group, error) {
	groups, err := s.groupRepo.GetGroupsJoinedByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (s *groupServiceImpl) InsertEvent(ctx context.Context, event *models.Event) error {
	fmt.Println("InsertEvent called with event:", event)
	if event == nil || event.UserID == uuid.Nil || event.GroupID == uuid.Nil || event.EventDate.IsZero() || event.Title == "" {
		return utils.ErrInvalidPayload
	}
	//vérifier si le groupe existe
	group, err := s.groupRepo.GetGroupByID(ctx, event.GroupID.String())
	if err != nil {
		if err == utils.ErrGroupNotFound {
			return utils.ErrGroupNotFound
		}
		return fmt.Errorf("failed to get group: %w", err)
	}
	// vérifier si l'utilisateur est membre du groupe
	isMember, err := s.groupRepo.IsMember(ctx, event.GroupID.String(), event.UserID.String())
	if err != nil {
		return fmt.Errorf("failed to check group membership: %w", err)
	}
	if !isMember {
		return utils.ErrForbidden
	}

	eventID, err := uuid.NewRandom()
	if err != nil {
		return utils.ErrUUIDGeneration
	}
	event.EventID = eventID
	event.CreationDate = time.Now()

	if err := s.groupRepo.InsertEvent(ctx, event); err != nil {
		return err
	}

	// Envoyer des notifications à tous les membres du groupe (sauf le créateur de l'événement)
	members, err := s.groupRepo.GetGroupMembers(ctx, event.GroupID.String())
	if err != nil {
		fmt.Printf("Warning: failed to get group members for notifications: %v\n", err)
		// On continue même si on ne peut pas envoyer les notifications
	} else {
		for _, member := range members {
			// Ne pas envoyer de notification au créateur de l'événement
			if member.UserID == event.UserID {
				continue
			}

			notification := &models.Notification{
				UserID:    member.UserID,
				Type:      models.NotificationTypeEventCreated,
				SenderID:  event.UserID.String(),
				GroupID:   &event.GroupID,
				GroupName: &group.Title, // Ajouter le nom du groupe
			}

			if err := s.notifSvc.CreateNotification(ctx, notification); err != nil {
				fmt.Printf("Warning: failed to create notification for user %s: %v\n", member.UserID.String(), err)
				// On continue même si une notification individuelle échoue
			}
		}
	}

	return nil
}

func (s *groupServiceImpl) InsertEventResponse(ctx context.Context, eventResponse *models.EventResponse) error {
	if eventResponse == nil || eventResponse.UserID == uuid.Nil || eventResponse.EventID == uuid.Nil {
		return utils.ErrInvalidPayload
	}
	// Vérifier si l'événement existe
	event, err := s.groupRepo.GetEventById(ctx, eventResponse.EventID.String())
	if err != nil {
		return fmt.Errorf("failed to get event: %w", err)
	}
	if event == nil {
		return utils.ErrEventNotFound
	}

	// Vérifier si l'utilisateur est membre du groupe
	isMember, err := s.groupRepo.IsMember(ctx, event.GroupID.String(), eventResponse.UserID.String())
	if err != nil {
		return fmt.Errorf("failed to check group membership: %w", err)
	}
	if !isMember {
		return utils.ErrForbidden
	}
	if err := s.groupRepo.InsertEventResponse(ctx, eventResponse); err != nil {
		return fmt.Errorf("failed to insert event response: %w", err)
	}
	// Si l'insertion réussit, on peut retourner nil

	return nil
}

func (s *groupServiceImpl) GetGroupEvents(ctx context.Context, groupID string) ([]*models.EventDetails, error) {

	// Vérifier si le groupe existe
	group, err := s.groupRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group: %w", err)
	}
	if group == nil {
		return nil, utils.ErrGroupNotFound
	}

	// Récupérer les événements du groupe
	events, err := s.groupRepo.GetGroupEvents(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group events: %w", err)
	}
	for _, event := range events {
		if event.AvatarURL != nil {
			event.AvatarURL = utils.GetFullImageURLAvatar(event.AvatarURL)
		}
	}

	return events, nil
}

func (s *groupServiceImpl) InviteUser(ctx context.Context, groupID, senderId, userID string) (uuid.UUID, error) {
	if groupID == "" || senderId == "" || userID == "" {
		return uuid.Nil, utils.ErrInvalidPayload
	}

	group, err := s.groupRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get group: %w", err)
	}
	if group == nil {
		return uuid.Nil, utils.ErrGroupNotFound
	}
	if group.UserID.String() == userID {
		return uuid.Nil, utils.ErrInvalidPayload
	}
	isMember, err := s.groupRepo.IsMember(ctx, groupID, userID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to check group membership: %w", err)
	}
	if isMember {
		return uuid.Nil, utils.ErrResourceConflict
	}
	// Créer une requête de suivi pour inviter l'utilisateur
	requestID, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to generate request ID: %w", err)
	}
	groupIDParsed := uuid.MustParse(groupID)
	request := models.FollowRequest{
		RequestID:  requestID,
		SenderID:   uuid.MustParse(senderId),
		ReceiverID: uuid.MustParse(userID),
		GroupID:    &groupIDParsed,
	}
	if err := s.groupRepo.InviteUser(ctx, request); err != nil {
		return uuid.Nil, fmt.Errorf("failed to invite user: %w", err)
	}

	return requestID, nil
}

func (s *groupServiceImpl) AnswerGroupInvite(ctx context.Context, request models.FollowRequest, userId, action string) error {
	if request.RequestID == uuid.Nil || userId == "" || action == "" {
		return utils.ErrInvalidPayload
	}

	if request.ReceiverID.String() != userId {
		return utils.ErrForbidden
	}

	if action != "accept" && action != "reject" {
		return utils.ErrInvalidUserData
	}

	if err := s.groupRepo.AnswerGroupInvite(ctx, request, userId, action); err != nil {
		return fmt.Errorf("failed to answer group invite: %w", err)
	}

	return nil
}
