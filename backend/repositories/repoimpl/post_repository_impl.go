package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"mellow/models"
	"mellow/repositories"
)

type postRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) repositories.PostRepository {
	return &postRepositoryImpl{db: db}
}

func (r *postRepositoryImpl) InsertPost(ctx context.Context, post *models.Post) error {
	query := `INSERT INTO posts (post_id, group_id, user_id, title, content, creation_date, visibility, image_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, post.PostID, post.GroupID, post.UserID, post.Title, post.Content, post.CreationDate, post.Visibility, post.ImageURL)
	if err != nil {
		return fmt.Errorf("failed to insert post: %w", err)
	}
	return nil
}
func (r *postRepositoryImpl) GetPostByID(ctx context.Context, postID string) (*models.PostDetails, error) {
	query := `
		SELECT 
			p.post_id, p.group_id, p.user_id, p.title, p.content, 
			p.creation_date, p.visibility, p.image_url, 
			u.username, u.image_url 
		FROM 
			posts p
		JOIN 
			users u ON p.user_id = u.user_id
		WHERE 
			p.post_id = ?`
	var post models.PostDetails
	err := r.db.QueryRowContext(ctx, query, postID).Scan(
		&post.PostID, &post.GroupID, &post.UserID, &post.Title,
		&post.Content, &post.CreationDate, &post.Visibility, &post.ImageURL,
		&post.Username,&post.AvatarURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Post not found
		}
		return nil, fmt.Errorf("error retrieving post: %w", err)
	}
	return &post, nil
}

func (r *postRepositoryImpl) DeletePost(ctx context.Context, postID string) error {
	// TODO: DELETE FROM posts WHERE id = ?
	return nil
}

func (r *postRepositoryImpl) GetFeed(ctx context.Context, userID string) ([]*models.Post, error) {
	// TODO: Requête pour récupérer les posts visibles pour userID (publics + suivis)
	return nil, nil
}

func (r *postRepositoryImpl) GetUserPosts(ctx context.Context, ownerID string) ([]*models.Post, error) {
	// TODO: SELECT * FROM posts WHERE author_id = ?
	return nil, nil
}

func (r *postRepositoryImpl) InsertPostReport(ctx context.Context, report *models.Report) error {
	// TODO: INSERT INTO reports (id, reporter_id, post_id, reason, created_at) VALUES (?, ?, ?, ?, ?)
	return nil
}
