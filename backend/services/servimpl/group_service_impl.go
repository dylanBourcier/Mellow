package servimpl

import (
	"context"
	"mellow/models"
	"mellow/repositories"
	"mellow/services"
	"mellow/utils"
	"time"

	"github.com/google/uuid"
)

type groupServiceImpl struct {
	groupRepo repositories.GroupRepository
}

// NewGroupService crée une nouvelle instance de GroupService.
func NewGroupService(groupRepo repositories.GroupRepository) services.GroupService {
	return &groupServiceImpl{groupRepo: groupRepo}
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
