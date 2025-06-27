package services

import (
	"context"
	"mellow/models"
)

type AdminService interface {
	// GetAllUsers retourne la liste de tous les utilisateurs.
	GetAllUsers(ctx context.Context) ([]*models.User, error)

	// PromoteToModerator promeut un utilisateur au rôle de modérateur.
	PromoteToModerator(ctx context.Context, userID string) error

	// DemoteToUser rétrograde un modérateur au rôle d'utilisateur simple.
	DemoteToUser(ctx context.Context, userID string) error

	// DeleteAnyUser supprime n’importe quel utilisateur du réseau.
	DeleteAnyUser(ctx context.Context, userID string) error

	// GetReportedContent retourne la liste du contenu signalé.
	GetReportedContent(ctx context.Context) ([]models.Report, error)

	// DeleteReportedContent supprime un contenu signalé.
	DeleteReportedContent(ctx context.Context, reportID string) error
}
