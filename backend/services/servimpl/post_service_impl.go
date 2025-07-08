package servimpl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/services"
)

type postServiceImpl struct {
	db *sql.DB
}

// NewPostService crée une nouvelle instance de PostService.
func NewPostService(db *sql.DB) services.PostService {
	return &postServiceImpl{db: db}
}

func (s *postServiceImpl) CreatePost(ctx context.Context, post *models.Post) error {
	// TODO: Vérifier la validité du post (contenu, visibilité)
	// TODO: Appeler le repository pour insérer le post
	return nil
}

func (s *postServiceImpl) GetPostByID(ctx context.Context, postID string, requesterID string) (*models.Post, error) {
	// TODO: Vérifier la visibilité du post pour le requester
	// TODO: Récupérer le post depuis le repository
	return nil, nil
}

func (s *postServiceImpl) DeletePost(ctx context.Context, postID, requesterID string) error {
	// TODO: Vérifier que le requester est l’auteur ou un modérateur
	// TODO: Supprimer le post via le repository
	return nil
}

func (s *postServiceImpl) GetFeed(ctx context.Context, userID string) ([]*models.Post, error) {
	// TODO: Récupérer les posts visibles par l’utilisateur (publics + privés de ses abonnements)
	return nil, nil
}

func (s *postServiceImpl) GetUserPosts(ctx context.Context, ownerID, requesterID string) ([]*models.Post, error) {
	// TODO: Vérifier si le requester peut voir les posts privés
	// TODO: Retourner les posts via le repository
	return nil, nil
}

func (s *postServiceImpl) ReportPost(ctx context.Context, report *models.Report) error {
	// TODO: Vérifier que le post existe
	// TODO: Enregistrer le signalement via le repository
	return nil
}
