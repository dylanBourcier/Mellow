package services

import (
	"context"
	"mellow/models"
)

type GroupService interface {
	// CreateGroup crée un nouveau groupe.
	CreateGroup(ctx context.Context, group *models.Group) error

	// UpdateGroup met à jour le titre ou la description d'un groupe.
	UpdateGroup(ctx context.Context, groupID, requesterID, title string, description string) error

	// GetGroupByID récupère un groupe par son ID.
	GetGroupByID(ctx context.Context, groupID string) (*models.Group, error)

	// GetAllGroups retourne tous les groupes existants.
	GetAllGroups(ctx context.Context) ([]*models.Group, error)

	// DeleteGroup supprime un groupe (par son créateur ou un admin).
	DeleteGroup(ctx context.Context, groupID, requesterID string) error

	// AddMember ajoute un membre dans un groupe.
	AddMember(ctx context.Context, groupID, userID string) error

	// RemoveMember retire un membre du groupe.
	RemoveMember(ctx context.Context, groupID, userID string) error

	// GetGroupMembers retourne les membres d’un groupe.
	GetGroupMembers(ctx context.Context, groupID string) ([]*models.User, error)

	// IsMember vérifie si un utilisateur est membre d’un groupe.
	IsMember(ctx context.Context, groupID, userID string) (bool, error)

	// GetGroupsJoinedByUser retourne les groupes auxquels un utilisateur a adhéré.
	GetGroupsJoinedByUser(ctx context.Context, userID string) ([]*models.Group, error)
}
