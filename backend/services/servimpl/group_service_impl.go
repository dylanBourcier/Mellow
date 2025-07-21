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

func (s *groupServiceImpl) GetGroupByID(ctx context.Context, groupID string) (*models.Group, error) {
	// TODO: Appeler le repository pour récupérer un groupe par ID
	return nil, nil
}

func (s *groupServiceImpl) GetAllGroups(ctx context.Context) ([]*models.Group, error) {
	// TODO: Appeler le repository pour récupérer tous les groupes
	return nil, nil
}

func (s *groupServiceImpl) DeleteGroup(ctx context.Context, groupID, requesterID string) error {
	// TODO: Vérifier que le requester est créateur ou admin
	// TODO: Appeler le repository pour supprimer le groupe
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
	// TODO: Vérifier que le membre existe et peut être retiré
	// TODO: Appeler le repository pour supprimer l'entrée
	return nil
}

func (s *groupServiceImpl) GetGroupMembers(ctx context.Context, groupID string) ([]*models.User, error) {
	// TODO: Appeler le repository pour récupérer les membres du groupe
	return nil, nil
}

func (s *groupServiceImpl) IsMember(ctx context.Context, groupID, userID string) (bool, error) {
	// TODO: Appeler le repository pour vérifier la relation d’appartenance
	return s.groupRepo.IsMember(ctx, groupID, userID)
}

func (s *groupServiceImpl) GetGroupsJoinedByUser(ctx context.Context, userID string) ([]*models.Group, error) {
	// Call the repository function to get groups joined by the user
	groups, err := s.groupRepo.GetGroupsJoinedByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
