// File: mellow/repositories/admin_repository.go
package repositories

import (
	"context"
	"mellow/models"
)

type AdminRepository interface {
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUserRole(ctx context.Context, userID string, newRole string) error
	DeleteUser(ctx context.Context, userID string) error

	GetAllReports(ctx context.Context) ([]models.Report, error)
	DeleteReportedContent(ctx context.Context, reportID string) error
}
