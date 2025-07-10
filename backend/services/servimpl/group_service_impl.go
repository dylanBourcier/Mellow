package servimpl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/services"
)

type groupServiceImpl struct {
	db *sql.DB
}

// NewGroupService crée une nouvelle instance de GroupService.
func NewGroupService(db *sql.DB) services.GroupService {
	return &groupServiceImpl{db: db}
}

func (s *groupServiceImpl) CreateGroup(ctx context.Context, group *models.Group) error {
	// TODO: Vérifier que le nom est unique, valider les données
	// TODO: Appeler le repository pour insérer le groupe
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
	// TODO: Vérifier que l'utilisateur n'est pas déjà membre
	// TODO: Appeler le repository pour insérer la relation dans groups_member
	return nil
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
	return false, nil
}
