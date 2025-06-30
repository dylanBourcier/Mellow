package impl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/repositories"
)

type adminRepositoryImpl struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) repositories.AdminRepository {
	return &adminRepositoryImpl{db: db}
}

func (r *adminRepositoryImpl) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	// TODO: SELECT * FROM users
	return nil, nil
}

func (r *adminRepositoryImpl) UpdateUserRole(ctx context.Context, userID string, newRole string) error {
	// TODO: UPDATE users SET role = ? WHERE id = ?
	return nil
}

func (r *adminRepositoryImpl) DeleteUser(ctx context.Context, userID string) error {
	// TODO: Supprimer l'utilisateur et éventuellement ses relations
	return nil
}

func (r *adminRepositoryImpl) GetAllReports(ctx context.Context) ([]models.Report, error) {
	// TODO: SELECT * FROM reports
	return nil, nil
}

func (r *adminRepositoryImpl) DeleteReportedContent(ctx context.Context, reportID string) error {
	// TODO: Supprimer un contenu signalé via son ID (et/ou supprimer l'entrée de report)
	return nil
}
