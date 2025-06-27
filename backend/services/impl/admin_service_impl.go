package impl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/services"
)

type adminServiceImpl struct {
	db *sql.DB
}

// NewAdminService crée une nouvelle instance de AdminService.
func NewAdminService(db *sql.DB) services.AdminService {
	return &adminServiceImpl{db: db}
}

func (s *adminServiceImpl) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	// TODO: Appeler le repository pour récupérer tous les utilisateurs
	return nil, nil
}

func (s *adminServiceImpl) PromoteToModerator(ctx context.Context, userID string) error {
	// TODO: Vérifier que l'utilisateur peut être promu
	// TODO: Appeler le repository pour modifier le rôle
	return nil
}

func (s *adminServiceImpl) DemoteToUser(ctx context.Context, userID string) error {
	// TODO: Vérifier que le modérateur peut être rétrogradé
	// TODO: Appeler le repository pour modifier le rôle
	return nil
}

func (s *adminServiceImpl) DeleteAnyUser(ctx context.Context, userID string) error {
	// TODO: Vérifier les contraintes (ex: ne pas supprimer un admin)
	// TODO: Supprimer toutes les données liées, puis l’utilisateur
	return nil
}

func (s *adminServiceImpl) GetReportedContent(ctx context.Context) ([]models.Report, error) {
	// TODO: Appeler le repository pour récupérer les contenus signalés
	return nil, nil
}

func (s *adminServiceImpl) DeleteReportedContent(ctx context.Context, reportID string) error {
	// TODO: Appliquer les règles (ex: contenu supprimable ?)
	// TODO: Supprimer le contenu concerné via le repository
	return nil
}
