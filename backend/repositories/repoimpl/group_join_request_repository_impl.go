package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"mellow/models"
	"mellow/repositories"
	"time"

	"github.com/google/uuid"
)

type groupJoinRequestRepositoryImpl struct {
	db *sql.DB
}

func NewGroupJoinRequestRepository(db *sql.DB) repositories.GroupJoinRequestRepository {
	return &groupJoinRequestRepositoryImpl{db: db}
}

func (r *groupJoinRequestRepositoryImpl) CreatePending(ctx context.Context, groupID, requesterID string) (*models.GroupJoinRequest, error) {
	id := uuid.New()
	_, err := r.db.ExecContext(ctx, `INSERT INTO group_join_requests (id, group_id, requester_id, status, created_at) VALUES (?,?,?,?,CURRENT_TIMESTAMP)`, id, groupID, requesterID, "pending")
	if err != nil {
		return nil, fmt.Errorf("failed to insert join request: %w", err)
	}
	// fetch row
	return r.GetByID(ctx, id.String())
}

func (r *groupJoinRequestRepositoryImpl) GetPendingByGroup(ctx context.Context, groupID string) ([]*models.GroupJoinRequest, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT gjr.id, gjr.group_id, gjr.requester_id, gjr.status, gjr.created_at, gjr.decided_at, gjr.decided_by,
               u.username, u.image_url
        FROM group_join_requests gjr
        JOIN users u ON u.user_id = gjr.requester_id
        WHERE gjr.group_id = ? AND gjr.status = 'pending'
        ORDER BY gjr.created_at ASC
    `, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to query pending join requests: %w", err)
	}
	defer rows.Close()
	var out []*models.GroupJoinRequest
	for rows.Next() {
		var item models.GroupJoinRequest
		var decidedBy sql.NullString
		var decidedAt sql.NullTime
		var username sql.NullString
		var avatar sql.NullString
		if err := rows.Scan(&item.ID, &item.GroupID, &item.RequesterID, &item.Status, &item.CreatedAt, &decidedAt, &decidedBy, &username, &avatar); err != nil {
			return nil, fmt.Errorf("failed to scan join request: %w", err)
		}
		if decidedAt.Valid {
			t := decidedAt.Time
			item.DecidedAt = &t
		}
		if decidedBy.Valid {
			uid, err := uuid.Parse(decidedBy.String)
			if err == nil {
				item.DecidedBy = &uid
			}
		}
		if username.Valid {
			s := username.String
			item.RequesterUsername = &s
		}
		if avatar.Valid {
			s := avatar.String
			item.RequesterAvatar = &s
		}
		out = append(out, &item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *groupJoinRequestRepositoryImpl) GetByID(ctx context.Context, requestID string) (*models.GroupJoinRequest, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, group_id, requester_id, status, created_at, decided_at, decided_by FROM group_join_requests WHERE id = ?`, requestID)
	var item models.GroupJoinRequest
	var decidedBy sql.NullString
	var decidedAt sql.NullTime
	if err := row.Scan(&item.ID, &item.GroupID, &item.RequesterID, &item.Status, &item.CreatedAt, &decidedAt, &decidedBy); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get join request: %w", err)
	}
	if decidedAt.Valid {
		t := decidedAt.Time
		item.DecidedAt = &t
	}
	if decidedBy.Valid {
		if uid, err := uuid.Parse(decidedBy.String); err == nil {
			item.DecidedBy = &uid
		}
	}
	return &item, nil
}

func (r *groupJoinRequestRepositoryImpl) ExistsPending(ctx context.Context, groupID, requesterID string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM group_join_requests WHERE group_id = ? AND requester_id = ? AND status = 'pending')`, groupID, requesterID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check pending existence: %w", err)
	}
	return exists, nil
}

func (r *groupJoinRequestRepositoryImpl) GetPendingByUserAndGroup(ctx context.Context, groupID, requesterID string) (*models.GroupJoinRequest, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, group_id, requester_id, status, created_at, decided_at, decided_by FROM group_join_requests WHERE group_id = ? AND requester_id = ? AND status = 'pending'`, groupID, requesterID)
	var item models.GroupJoinRequest
	var decidedBy sql.NullString
	var decidedAt sql.NullTime
	if err := row.Scan(&item.ID, &item.GroupID, &item.RequesterID, &item.Status, &item.CreatedAt, &decidedAt, &decidedBy); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get pending join request: %w", err)
	}
	if decidedAt.Valid {
		t := decidedAt.Time
		item.DecidedAt = &t
	}
	if decidedBy.Valid {
		if uid, err := uuid.Parse(decidedBy.String); err == nil {
			item.DecidedBy = &uid
		}
	}
	return &item, nil
}

func (r *groupJoinRequestRepositoryImpl) SetStatus(ctx context.Context, requestID, status, decidedBy string) error {
	var decidedByPtr interface{}
	if decidedBy == "" {
		decidedByPtr = nil
	} else {
		decidedByPtr = decidedBy
	}
	_, err := r.db.ExecContext(ctx, `UPDATE group_join_requests SET status = ?, decided_at = ?, decided_by = ? WHERE id = ?`, status, time.Now(), decidedByPtr, requestID)
	if err != nil {
		return fmt.Errorf("failed to update join request status: %w", err)
	}
	return nil
}

func (r *groupJoinRequestRepositoryImpl) IsMember(ctx context.Context, groupID, userID string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM groups_member WHERE group_id = ? AND user_id = ?)`, groupID, userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check membership: %w", err)
	}
	return exists, nil
}
