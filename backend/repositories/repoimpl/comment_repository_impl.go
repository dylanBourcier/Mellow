package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
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
	query := `INSERT INTO comments (comment_id, post_id, user_id, content, creation_date, image_url) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, comment.CommentID, comment.PostID, comment.UserID, comment.Content, comment.CreationDate, comment.ImageURL)
	if err != nil {
		return fmt.Errorf("failed to insert comment: %w", err)
	}
	return nil
}

func (r *commentRepositoryImpl) GetCommentsByPostID(ctx context.Context, postID string) ([]*models.CommentDetails, error) {
	query := `SELECT c.comment_id, c.post_id, c.content, c.creation_date, c.image_url, u.user_id, u.username, u.image_url
			  FROM comments c
			  JOIN users u ON c.user_id = u.user_id
			  WHERE c.post_id = ?
			  ORDER BY c.creation_date ASC`
	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments for post %s: %w", postID, err)
	}
	defer rows.Close()
	var comments []*models.CommentDetails
	for rows.Next() {
		var comment models.CommentDetails
		if err := rows.Scan(&comment.CommentID, &comment.PostID, &comment.Content, &comment.CreationDate, &comment.ImageURL, &comment.UserID, &comment.Username, &comment.AvatarURL); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, &comment)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over comments: %w", err)
	}

	return comments, nil

}

func (r *commentRepositoryImpl) DeleteComment(ctx context.Context, commentID string) error {
	// TODO: DELETE FROM comments WHERE id = ?
	return nil
}

func (r *commentRepositoryImpl) UpdateComment(ctx context.Context, commentID string, content string) error {
	query := `UPDATE comments SET content = ? WHERE comment_id = ?`
	_, err := r.db.ExecContext(ctx, query, content, commentID)
	if err != nil {
		return fmt.Errorf("failed to update comment: %w", err)
	}
	return nil
}

func (r *commentRepositoryImpl) GetCommentByID(ctx context.Context, commentID string) (*models.Comment, error) {
	query := `SELECT comment_id, user_id, post_id, content, creation_date, image_url FROM comments WHERE comment_id = ?`
	var c models.Comment
	err := r.db.QueryRowContext(ctx, query, commentID).Scan(&c.CommentID, &c.UserID, &c.PostID, &c.Content, &c.CreationDate, &c.ImageURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get comment: %w", err)
	}
	return &c, nil
}

func (r *commentRepositoryImpl) InsertCommentReport(ctx context.Context, report *models.Report) error {
	// TODO: INSERT INTO reports (id, reporter_id, comment_id, reason, created_at) VALUES (?, ?, ?, ?, ?)
	return nil
}
