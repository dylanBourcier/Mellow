package services

import (
	"context"
	"mellow/models"
)

// UserService définit les méthodes liées à la gestion des utilisateurs.
type UserService interface {
	// CreateUser enregistre un nouvel utilisateur dans la base de données.
	CreateUser(ctx context.Context, user *models.User) error

	// GetUserByID retourne un utilisateur à partir de son ID.
	GetUserByID(ctx context.Context, userID string) (*models.User, error)

	// GetUserByUsername retourne un utilisateur à partir de son nom d'utilisateur.
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)

	// UpdateUser modifie les informations d'un utilisateur.
	UpdateUser(ctx context.Context, user *models.User) error

	// DeleteUser supprime un utilisateur et ses données associées.
	DeleteUser(ctx context.Context, userID string) error

	// Authenticate vérifie les identifiants et retourne l'utilisateur si valide.
	Authenticate(ctx context.Context, username, password string) (*models.User, error)

	// FollowUser permet à un utilisateur d'en suivre un autre.
	FollowUser(ctx context.Context, followerID, targetID string) error

	// UnfollowUser permet à un utilisateur d'arrêter de suivre un autre.
	UnfollowUser(ctx context.Context, followerID, targetID string) error

	// GetFollowers retourne la liste des utilisateurs qui suivent un utilisateur donné.
	GetFollowers(ctx context.Context, userID string) ([]*models.User, error)

	// GetFollowing retourne la liste des utilisateurs suivis par un utilisateur donné.
	GetFollowing(ctx context.Context, userID string) ([]*models.User, error)

	// SearchUsers retourne une liste d'utilisateurs correspondant à un mot-clé.
	SearchUsers(ctx context.Context, query string) ([]*models.User, error)
}
