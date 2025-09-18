package services

import (
	"context"
	"mellow/models"
)

type GroupJoinRequestService interface {
	RequestJoin(ctx context.Context, userID, groupID string) (*models.GroupJoinRequest, error)
	ListPending(ctx context.Context, ownerID, groupID string) ([]*models.GroupJoinRequest, error)
	Accept(ctx context.Context, ownerID, groupID, requestID string) (*models.GroupJoinRequest, error)
	Reject(ctx context.Context, ownerID, groupID, requestID string) (*models.GroupJoinRequest, error)
	Cancel(ctx context.Context, userID, groupID string) error
	GetSelfPending(ctx context.Context, userID, groupID string) (*models.GroupJoinRequest, error)
}
