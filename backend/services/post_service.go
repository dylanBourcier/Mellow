package services

import (
	"context"
	"mellow/models"
)

type PostService interface {
	// CreatePost publie un nouveau post.
	CreatePost(ctx context.Context, post *models.Post) error

	// GetPostByID retourne un post par son ID.
	GetPostByID(ctx context.Context, postID string, groupService GroupService, userService UserService, requesterID string) (*models.PostDetails, error)

	// DeletePost supprime un post (par son auteur ou un modérateur).
	DeletePost(ctx context.Context, postID, requesterID string) error

	// GetFeed retourne le flux de posts visibles par l’utilisateur connecté.
	GetFeed(ctx context.Context, userID *string, limit, offset int) ([]*models.PostDetails, error)

	// GetUserPosts retourne les posts publics ou privés d’un utilisateur (en fonction du rôle).
	GetUserPosts(ctx context.Context, ownerID, requesterID string) ([]*models.Post, error)

	// ReportPost permet de signaler un post inapproprié.
	ReportPost(ctx context.Context, report *models.Report) error

	// IsPostExisting vérifie si un post existe déjà.
	IsPostExisting(ctx context.Context, postID string) (bool, error)
}
