package services

import (
	"context"
	"mellow/models"
)

type CommentService interface {
	// CreateComment ajoute un nouveau commentaire à un post.
	CreateComment(ctx context.Context, comment *models.Comment) error

	// GetCommentsByPostID récupère tous les commentaires liés à un post.
	GetCommentsByPostID(ctx context.Context, postID string) ([]*models.CommentDetails, error)

	// DeleteComment supprime un commentaire spécifique (par son auteur ou un modérateur).
	DeleteComment(ctx context.Context, commentID, requesterID string) error

	// UpdateComment modifie le contenu d'un commentaire existant.
	UpdateComment(ctx context.Context, commentID, requesterID, content string) error

	// ReportComment permet de signaler un commentaire inapproprié.
	ReportComment(ctx context.Context, report *models.Report) error
}
