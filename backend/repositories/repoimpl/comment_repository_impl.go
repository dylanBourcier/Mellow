package repoimpl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/repositories"
)

type commentRepositoryImpl struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) repositories.CommentRepository {
	return &commentRepositoryImpl{db: db}
}

func (r *commentRepositoryImpl) InsertComment(ctx context.Context, comment *models.Comment) error {
	// TODO: INSERT INTO comments (id, post_id, author_id, content, created_at) VALUES (?, ?, ?, ?, ?)
	return nil
}

func (r *commentRepositoryImpl) GetCommentsByPostID(ctx context.Context, postID string) ([]*models.Comment, error) {
	// TODO: SELECT * FROM comments WHERE post_id = ? ORDER BY created_at ASC
	return nil, nil
}

func (r *commentRepositoryImpl) DeleteComment(ctx context.Context, commentID string) error {
	// TODO: DELETE FROM comments WHERE id = ?
	return nil
}

func (r *commentRepositoryImpl) InsertCommentReport(ctx context.Context, report *models.Report) error {
	// TODO: INSERT INTO reports (id, reporter_id, comment_id, reason, created_at) VALUES (?, ?, ?, ?, ?)
	return nil
}
