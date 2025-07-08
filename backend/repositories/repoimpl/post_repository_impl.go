package repoimpl

import (
	"context"
	"database/sql"
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
	// TODO: INSERT INTO posts (id, author_id, content, visibility, created_at) VALUES (?, ?, ?, ?, ?)
	return nil
}

func (r *postRepositoryImpl) GetPostByID(ctx context.Context, postID string) (*models.Post, error) {
	// TODO: SELECT * FROM posts WHERE id = ?
	return nil, nil
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
