package repositories

import (
	"context"
	"mellow/models"
)

type GroupJoinRequestRepository interface {
	CreatePending(ctx context.Context, groupID, requesterID string) (*models.GroupJoinRequest, error)
	GetPendingByGroup(ctx context.Context, groupID string) ([]*models.GroupJoinRequest, error)
	GetByID(ctx context.Context, requestID string) (*models.GroupJoinRequest, error)
	ExistsPending(ctx context.Context, groupID, requesterID string) (bool, error)
	GetPendingByUserAndGroup(ctx context.Context, groupID, requesterID string) (*models.GroupJoinRequest, error)
	SetStatus(ctx context.Context, requestID, status, decidedBy string) error
	IsMember(ctx context.Context, groupID, userID string) (bool, error)
}
