package repositories

import (
	"context"
	"mellow/models"

	"github.com/google/uuid"
)

type AuthRepository interface {
	// CreateSession enregistre une nouvelle session utilisateur.
	CreateSession(ctx context.Context, session *models.Session) error

	// DeleteSession supprime une session existante.
	DeleteSession(ctx context.Context, sessionID string) error

	// GetUserFromSession récupère un utilisateur à partir de son ID de session.
	GetUserFromSession(ctx context.Context, sessionID string) (*models.User, error)

	// GetUserIDFromSession récupère l'ID utilisateur à partir de l'ID de session.
	GetUserIDFromSession(ctx context.Context, sessionID string) (uuid.UUID, error)

	// IsSessionValid vérifie si une session existe et est encore valide.
	IsSessionValid(ctx context.Context, sessionID string) (bool, error)

	// DeleteExpiredSessions supprime les sessions expirées (par exemple, celles inactives depuis plus de 7 jours).
	DeleteExpiredSessions(ctx context.Context) error

	UpdateLastActivity(ctx context.Context, sessionID string) error
}
