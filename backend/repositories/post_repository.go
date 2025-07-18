package repositories

import (
	"context"
	"mellow/models"
)

type PostRepository interface {
	// InsertPost crée un nouveau post.
	InsertPost(ctx context.Context, post *models.Post) error

	// GetPostByID retourne un post par son ID.
	GetPostByID(ctx context.Context, postID string) (*models.PostDetails, error)

	// DeletePost supprime un post par son ID.
	DeletePost(ctx context.Context, postID string) error

	// GetFeed récupère les posts visibles par un utilisateur (ex : publics ou des gens suivis).
	GetFeed(ctx context.Context, userID *string, limit, offset int) ([]*models.PostDetails, error)

	// GetUserPosts retourne tous les posts d’un utilisateur.
	GetUserPosts(ctx context.Context, ownerID string) ([]*models.Post, error)

	// InsertPostReport signale un post via un report.
	InsertPostReport(ctx context.Context, report *models.Report) error
}
