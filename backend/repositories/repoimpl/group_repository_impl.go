package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
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
	query := `INSERT INTO groups (group_id, user_id, title, description, creation_date) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, group.GroupID, group.UserID, group.Title, group.Description, group.CreationDate)
	if err != nil {
		return fmt.Errorf("failed to insert group: %w", err)
	}
	return nil
}

func (r *groupRepositoryImpl) GetGroupByID(ctx context.Context, groupID string) (*models.Group, error) {
	query := `SELECT group_id, user_id, title, description, creation_date FROM groups WHERE group_id = ?`
	var g models.Group
	err := r.db.QueryRowContext(ctx, query, groupID).Scan(&g.GroupID, &g.UserID, &g.Title, &g.Description, &g.CreationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &g, nil
}

func (r *groupRepositoryImpl) GetAllGroups(ctx context.Context) ([]*models.Group, error) {
	query := `SELECT group_id, user_id, title, description, creation_date FROM groups ORDER BY creation_date DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*models.Group
	for rows.Next() {
		var g models.Group
		if err := rows.Scan(&g.GroupID, &g.UserID, &g.Title, &g.Description, &g.CreationDate); err != nil {
			return nil, err
		}
		groups = append(groups, &g)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *groupRepositoryImpl) DeleteGroup(ctx context.Context, groupID string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM groups WHERE group_id = ?`, groupID)
	if err != nil {
		return fmt.Errorf("failed to delete group: %w", err)
	}
	return nil
}

func (r *groupRepositoryImpl) UpdateGroup(ctx context.Context, group *models.Group) error {
	query := `UPDATE groups SET title = ?, description = ? WHERE group_id = ?`
	_, err := r.db.ExecContext(ctx, query, group.Title, group.Description, group.GroupID)
	if err != nil {
		return err
	}
	return nil
}

func (r *groupRepositoryImpl) AddMember(ctx context.Context, groupID, userID string) error {
	// TODO: INSERT INTO groups_member (group_id, user_id) VALUES (?, ?)
	query := `INSERT INTO groups_member (group_id, user_id, role, join_date) VALUES (?, ?, 'member', CURRENT_TIMESTAMP)`
	_, err := r.db.ExecContext(ctx, query, groupID, userID)
	if err != nil {
		return fmt.Errorf("failed to add member: %w", err)
	}
	return nil
}

func (r *groupRepositoryImpl) RemoveMember(ctx context.Context, groupID, userID string) error {
	// TODO: DELETE FROM groups_member WHERE group_id = ? AND user_id = ?
	return nil
}

func (r *groupRepositoryImpl) GetGroupMembers(ctx context.Context, groupID string) ([]*models.User, error) {
	query := `SELECT u.user_id, u.email, u.password, u.username, u.firstname, u.lastname, u.birthdate, u.role, u.image_url, u.creation_date, u.description
                        FROM users u
                        JOIN groups_member gm ON u.user_id = gm.user_id
                        WHERE gm.group_id = ?`
	rows, err := r.db.QueryContext(ctx, query, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to get group members: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UserID, &u.Email, &u.Password, &u.Username, &u.Firstname, &u.Lastname, &u.Birthdate, &u.Role, &u.ImageURL, &u.CreationDate, &u.Description); err != nil {
			return nil, fmt.Errorf("failed to scan group member: %w", err)
		}
		users = append(users, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating group members: %w", err)
	}
	return users, nil
}

func (r *groupRepositoryImpl) IsMember(ctx context.Context, groupID, userID string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM groups_member WHERE group_id = ? AND user_id = ?)`
	err := r.db.QueryRowContext(ctx, query, groupID, userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *groupRepositoryImpl) GetGroupsJoinedByUser(ctx context.Context, userID string) ([]*models.Group, error) {
	query := `SELECT g.* FROM groups g
			JOIN groups_member gm ON g.group_id = gm.group_id
			WHERE gm.user_id = ?`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var groups []*models.Group
	for rows.Next() {
		var group models.Group
		if err := rows.Scan(&group.GroupID, &group.Title, &group.Description, &group.UserID, &group.CreationDate); err != nil {
			return nil, err
		}
		groups = append(groups, &group)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *groupRepositoryImpl) IsTitleTaken(ctx context.Context, title string) (bool, error) {
	var count int
	err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM groups WHERE title = ?`, title).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check title: %w", err)
	}
	return count > 0, nil
}
