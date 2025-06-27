package repositories

import (
	"context"
	"mellow/models"
)

type AuthRepository interface {
	// CreateSession enregistre une nouvelle session utilisateur.
	CreateSession(ctx context.Context, session *models.Session) error

	// DeleteSession supprime une session existante.
	DeleteSession(ctx context.Context, sessionID string) error

	// GetUserFromSession récupère un utilisateur à partir de son ID de session.
	GetUserFromSession(ctx context.Context, sessionID string) (*models.User, error)

	// IsSessionValid vérifie si une session existe et est encore valide.
	IsSessionValid(ctx context.Context, sessionID string) (bool, error)
}
