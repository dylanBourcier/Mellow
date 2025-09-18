package servimpl

import (
	"context"
	"fmt"
	"mellow/models"
	"mellow/repositories"
	"mellow/services"
	"mellow/utils"
	"mellow/websocket"
	"time"
)

type groupJoinRequestServiceImpl struct {
	repo       repositories.GroupJoinRequestRepository
	groupsRepo repositories.GroupRepository
	notifSvc   services.NotificationService
}

func NewGroupJoinRequestService(repo repositories.GroupJoinRequestRepository, groupsRepo repositories.GroupRepository, notifSvc services.NotificationService) services.GroupJoinRequestService {
	return &groupJoinRequestServiceImpl{repo: repo, groupsRepo: groupsRepo, notifSvc: notifSvc}
}

func (s *groupJoinRequestServiceImpl) RequestJoin(ctx context.Context, userID, groupID string) (*models.GroupJoinRequest, error) {
	if userID == "" || groupID == "" {
		return nil, utils.ErrInvalidPayload
	}
	group, err := s.groupsRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, utils.ErrGroupNotFound
	}
	if group.UserID.String() == userID {
		return nil, utils.ErrInvalidPayload
	}
	isMember, err := s.repo.IsMember(ctx, groupID, userID)
	if err != nil {
		return nil, err
	}
	if isMember {
		return nil, utils.ErrResourceConflict
	}
	exists, err := s.repo.ExistsPending(ctx, groupID, userID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, utils.ErrResourceConflict
	}
	req, err := s.repo.CreatePending(ctx, groupID, userID)
	if err != nil {
		return nil, err
	}
	// realtime event
	websocket.BroadcastMessage("group:"+groupID, websocket.WSMessage{Type: "system", Room: "group:" + groupID, Content: "join_request.created", SenderID: userID, Timestamp: time.Now().Format(time.RFC3339)})
	// notification to owner using requestID link
	rid := req.ID
	_ = s.notifSvc.CreateNotification(ctx, &models.Notification{
		UserID:       group.UserID,
		SenderID:     userID,
		RequestID:    &rid,
		Type:         models.NotificationTypeGroupRequest,
		Seen:         false,
		CreationDate: time.Now(),
	})
	// notification to owner
	_ = s.notifSvc.CreateNotification(ctx, &models.Notification{
		UserID:   group.UserID,
		SenderID: userID,
		Type:     models.NotificationTypeGroupRequest,
		Seen:     false,
		// Group info is not persisted in notifications table; UI should query the group panel
		CreationDate: time.Now(),
	})
	return req, nil
}

func (s *groupJoinRequestServiceImpl) ListPending(ctx context.Context, ownerID, groupID string) ([]*models.GroupJoinRequest, error) {
	if ownerID == "" || groupID == "" {
		return nil, utils.ErrInvalidPayload
	}
	group, err := s.groupsRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, utils.ErrGroupNotFound
	}
	if group.UserID.String() != ownerID {
		return nil, utils.ErrForbidden
	}
	list, err := s.repo.GetPendingByGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	// Add full avatar URLs
	for _, it := range list {
		if it.RequesterAvatar != nil {
			it.RequesterAvatar = utils.GetFullImageURLAvatar(it.RequesterAvatar)
		}
	}
	return list, nil
}

func (s *groupJoinRequestServiceImpl) Accept(ctx context.Context, ownerID, groupID, requestID string) (*models.GroupJoinRequest, error) {
	if ownerID == "" || groupID == "" || requestID == "" {
		return nil, utils.ErrInvalidPayload
	}
	group, err := s.groupsRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, utils.ErrGroupNotFound
	}
	if group.UserID.String() != ownerID {
		return nil, utils.ErrForbidden
	}
	req, err := s.repo.GetByID(ctx, requestID)
	if err != nil {
		return nil, err
	}
	if req == nil || req.GroupID.String() != groupID {
		return nil, utils.ErrBadRequest
	}
	if req.Status != "pending" {
		return nil, utils.ErrResourceConflict
	}
	// idempotent add member
	if err := s.groupsRepo.AddMember(ctx, groupID, req.RequesterID.String()); err != nil {
		// if unique violation exists, keep going to set status to accepted
		fmt.Println("AddMember error:", err)
	}
	if err := s.repo.SetStatus(ctx, requestID, "accepted", ownerID); err != nil {
		return nil, err
	}
	websocket.BroadcastMessage("group:"+groupID, websocket.WSMessage{Type: "system", Room: "group:" + groupID, Content: "join_request.accepted", SenderID: ownerID, Timestamp: time.Now().Format(time.RFC3339)})
	// notify requester
	rid := req.ID
	_ = s.notifSvc.CreateNotification(ctx, &models.Notification{
		UserID:       req.RequesterID,
		SenderID:     ownerID,
		RequestID:    &rid,
		Type:         models.NotificationTypeAcceptedGroupRequest,
		CreationDate: time.Now(),
		Seen:         false,
	})
	return s.repo.GetByID(ctx, requestID)
}

func (s *groupJoinRequestServiceImpl) Reject(ctx context.Context, ownerID, groupID, requestID string) (*models.GroupJoinRequest, error) {
	if ownerID == "" || groupID == "" || requestID == "" {
		return nil, utils.ErrInvalidPayload
	}
	group, err := s.groupsRepo.GetGroupByID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, utils.ErrGroupNotFound
	}
	if group.UserID.String() != ownerID {
		return nil, utils.ErrForbidden
	}
	req, err := s.repo.GetByID(ctx, requestID)
	if err != nil {
		return nil, err
	}
	if req == nil || req.GroupID.String() != groupID {
		return nil, utils.ErrBadRequest
	}
	if req.Status != "pending" {
		return nil, utils.ErrResourceConflict
	}
	if err := s.repo.SetStatus(ctx, requestID, "rejected", ownerID); err != nil {
		return nil, err
	}
	websocket.BroadcastMessage("group:"+groupID, websocket.WSMessage{Type: "system", Room: "group:" + groupID, Content: "join_request.rejected", SenderID: ownerID, Timestamp: time.Now().Format(time.RFC3339)})
	rid := req.ID
	_ = s.notifSvc.CreateNotification(ctx, &models.Notification{
		UserID:       req.RequesterID,
		SenderID:     ownerID,
		RequestID:    &rid,
		Type:         models.NotificationTypeRejectedGroupRequest,
		CreationDate: time.Now(),
		Seen:         false,
	})
	return s.repo.GetByID(ctx, requestID)
}

func (s *groupJoinRequestServiceImpl) Cancel(ctx context.Context, userID, groupID string) error {
	if userID == "" || groupID == "" {
		return utils.ErrInvalidPayload
	}
	req, err := s.repo.GetPendingByUserAndGroup(ctx, groupID, userID)
	if err != nil {
		return err
	}
	if req == nil {
		return utils.ErrBadRequest
	}
	if err := s.repo.SetStatus(ctx, req.ID.String(), "cancelled", ""); err != nil {
		return err
	}
	websocket.BroadcastMessage("group:"+groupID, websocket.WSMessage{Type: "system", Room: "group:" + groupID, Content: "join_request.cancelled", SenderID: userID, Timestamp: time.Now().Format(time.RFC3339)})
	return nil
}

func (s *groupJoinRequestServiceImpl) GetSelfPending(ctx context.Context, userID, groupID string) (*models.GroupJoinRequest, error) {
	if userID == "" || groupID == "" {
		return nil, utils.ErrInvalidPayload
	}
	return s.repo.GetPendingByUserAndGroup(ctx, groupID, userID)
}
