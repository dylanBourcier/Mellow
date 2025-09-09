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

	SendFollowRequest(ctx context.Context, followRequest models.FollowRequest) error
	GetFollowRequestByID(ctx context.Context, requestID string) (*models.FollowRequest, error)
	AnswerFollowRequest(ctx context.Context, request models.FollowRequest, userId, action string) error
	IsFollowRequestExists(ctx context.Context, senderID, receiverID string) (bool, error)

	InsertFollow(ctx context.Context, follower_id, followed_id string) error
	Unfollow(ctx context.Context, followerID, targetID string) error
	IsFollowing(ctx context.Context, followerID, targetID string) (bool, error)

	GetFollowers(ctx context.Context, viewerID, userID string) ([]*models.UserProfileData, error)
	GetFollowing(ctx context.Context, viewerID, userID string) ([]*models.UserProfileData, error)

	GetUserProfile(ctx context.Context, viewerID, userID string) (*models.UserProfileData, error)

	GetUserPrivacy(ctx context.Context, userID string) (string, error)
	SearchUsers(ctx context.Context, query string) ([]*models.User, error)
	SearchUsersExcludingGroupMembers(ctx context.Context, query, groupId string) ([]*models.User, error)
}
