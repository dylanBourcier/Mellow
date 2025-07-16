package repositories

import (
	"context"
	"mellow/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	FindUserByID(ctx context.Context, userID string) (*models.User, error)
	FindUserByUsername(ctx context.Context, username string) (*models.User, error)
	UserExistsByEmailOrUsername(ctx context.Context, email, username string) (bool, error)
	GetUserByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, userID string) error

	Follow(ctx context.Context, followerID, targetID string) error
	Unfollow(ctx context.Context, followerID, targetID string) error
	IsFollowing(ctx context.Context, followerID, targetID string) (bool, error)

	GetFollowers(ctx context.Context, userID string) ([]*models.User, error)
	GetFollowing(ctx context.Context, userID string) ([]*models.User, error)

	SearchUsers(ctx context.Context, query string) ([]*models.User, error)
}
