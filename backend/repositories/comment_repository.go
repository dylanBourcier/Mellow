package repositories

import (
	"context"
	"mellow/models"
)

type CommentRepository interface {
	// InsertComment ajoute un nouveau commentaire.
	InsertComment(ctx context.Context, comment *models.Comment) error

	// GetCommentsByPostID récupère tous les commentaires liés à un post.
	GetCommentsByPostID(ctx context.Context, postID string) ([]*models.CommentDetails, error)

	// DeleteComment supprime un commentaire par son ID.
	DeleteComment(ctx context.Context, commentID string) error

	// UpdateComment met à jour le contenu d'un commentaire.
	UpdateComment(ctx context.Context, commentID string, content string) error

	// GetCommentByID récupère un commentaire par son ID.
	GetCommentByID(ctx context.Context, commentID string) (*models.Comment, error)

	// InsertCommentReport enregistre un signalement de commentaire.
	InsertCommentReport(ctx context.Context, report *models.Report) error
}
