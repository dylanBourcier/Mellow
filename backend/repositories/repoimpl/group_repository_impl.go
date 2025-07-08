package repoimpl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/repositories"
)

type groupRepositoryImpl struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) repositories.GroupRepository {
	return &groupRepositoryImpl{db: db}
}

func (r *groupRepositoryImpl) InsertGroup(ctx context.Context, group *models.Group) error {
	// TODO: INSERT INTO groups (id, name, description, creator_id, created_at) VALUES (?, ?, ?, ?, ?)
	return nil
}

func (r *groupRepositoryImpl) GetGroupByID(ctx context.Context, groupID string) (*models.Group, error) {
	// TODO: SELECT * FROM groups WHERE id = ?
	return nil, nil
}

func (r *groupRepositoryImpl) GetAllGroups(ctx context.Context) ([]*models.Group, error) {
	// TODO: SELECT * FROM groups ORDER BY created_at DESC
	return nil, nil
}

func (r *groupRepositoryImpl) DeleteGroup(ctx context.Context, groupID string) error {
	// TODO: DELETE FROM groups WHERE id = ?
	return nil
}

func (r *groupRepositoryImpl) AddMember(ctx context.Context, groupID, userID string) error {
	// TODO: INSERT INTO groups_member (group_id, user_id) VALUES (?, ?)
	return nil
}

func (r *groupRepositoryImpl) RemoveMember(ctx context.Context, groupID, userID string) error {
	// TODO: DELETE FROM groups_member WHERE group_id = ? AND user_id = ?
	return nil
}

func (r *groupRepositoryImpl) GetGroupMembers(ctx context.Context, groupID string) ([]*models.User, error) {
	// TODO: SELECT u.* FROM users u JOIN groups_member gm ON u.id = gm.user_id WHERE gm.group_id = ?
	return nil, nil
}

func (r *groupRepositoryImpl) IsMember(ctx context.Context, groupID, userID string) (bool, error) {
	// TODO: SELECT 1 FROM groups_member WHERE group_id = ? AND user_id = ?
	return false, nil
}
