package services

import (
	"context"
	"mellow/models"

	"github.com/google/uuid"
)

// UserService définit les méthodes liées à la gestion des utilisateurs.
type UserService interface {
	// CreateUser enregistre un nouvel utilisateur dans la base de données.
	CreateUser(ctx context.Context, user *models.User) error

	// GetUserByID retourne un utilisateur à partir de son ID.
	GetUserByID(ctx context.Context, userID string) (*models.User, error)

	// GetUserByUsername retourne un utilisateur à partir de son nom d'utilisateur.
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)

	// GetUserByUsernameOrEmail retourne un utilisateur à partir de son nom d'utilisateur ou de son email.
	GetUserByUsernameOrEmail(ctx context.Context, identifier string) (*models.User, error)

	// UpdateUser modifie les informations d'un utilisateur.
	UpdateUser(ctx context.Context, user *models.User) error

	// DeleteUser supprime un utilisateur et ses données associées.
	DeleteUser(ctx context.Context, userID string) error

	// Authenticate vérifie les identifiants et retourne l'utilisateur si valide.
	Authenticate(ctx context.Context, username, password string) (*models.User, error)

	//SendFollowRequest envoie une demande de suivi à un utilisateur.
	SendFollowRequest(ctx context.Context, senderID, receiverID string) (uuid.UUID, error)

	// UnfollowUser permet à un utilisateur d'arrêter de suivre un autre.
	UnfollowUser(ctx context.Context, followerID, targetID string) error

	// GetFollowers retourne la liste des utilisateurs qui suivent un utilisateur donné.
	GetFollowers(ctx context.Context, viewerID, userID string) ([]*models.UserProfileData, error)

	// GetFollowing retourne la liste des utilisateurs suivis par un utilisateur donné.
	GetFollowing(ctx context.Context, viewerID, userID string) ([]*models.UserProfileData, error)

	// SearchUsers retourne une liste d'utilisateurs correspondant à un mot-clé.
	SearchUsers(ctx context.Context, query string, groupId string, excludeGroupMembers bool) ([]*models.User, error)

	// IsFollowing vérifie si un utilisateur suit un autre.
	IsFollowing(ctx context.Context, followerID, targetID string) (bool, error)

	// GetUserProfileData retourne les données de profil d'un utilisateur.
	GetUserProfileData(ctx context.Context, viewerID, userID string) (*models.UserProfileData, error)

	// InsertFollow insère une relation de suivi entre deux utilisateurs.
	InsertFollow(ctx context.Context, followerID, followedID string) error

	// GetUserPrivacy retourne la confidentialité d'un utilisateur.
	GetUserPrivacy(ctx context.Context, userID string) (string, error)

	// GetFollowRequestById retourne une demande de suivi par son ID.
	GetFollowRequestByID(ctx context.Context, requestID string) (*models.FollowRequest, error)

	//AnswerFollowRequest permet à un utilisateur d'accepter ou de rejeter une demande de suivi.
	AnswerFollowRequest(ctx context.Context, request models.FollowRequest, userId, action string) error

	// IsFollowRequestExists vérifie si une demande de suivi existe déjà.
	IsFollowRequestExists(ctx context.Context, senderID, targetID string) (bool, error)
}
