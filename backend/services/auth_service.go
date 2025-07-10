package services

import (
	"context"
	"mellow/models"
)

type AuthService interface {
	// Login vérifie les identifiants et crée une session.
	Login(ctx context.Context, username, password string) (*models.User, string, error)

	// Logout détruit la session utilisateur.
	Logout(ctx context.Context, sessionID string) error

	// IsAuthenticated vérifie si une session est valide.
	IsAuthenticated(ctx context.Context, sessionID string) (bool, error)

	// GetUserFromSession retourne l'utilisateur lié à une session.
	GetUserFromSession(ctx context.Context, sessionID string) (*models.User, error)
}
