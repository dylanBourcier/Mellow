package servimpl

import (
	"context"
	"mellow/models"
	"mellow/repositories"
	"mellow/services"
	"mellow/utils"
	"time"

	"github.com/google/uuid"
)

type postServiceImpl struct {
	postRepo repositories.PostRepository
}

// NewPostService crée une nouvelle instance de PostService.
func NewPostService(postRepo repositories.PostRepository) services.PostService {
	return &postServiceImpl{postRepo: postRepo}
}

func (s *postServiceImpl) CreatePost(ctx context.Context, post *models.Post) error {
	// TODO: Vérifier la validité du post (contenu, visibilité)
	// TODO: Appeler le repository pour insérer le post
	if post.Title == "" || post.Content == "" || post.Visibility == "" {
		return utils.ErrInvalidPayload
	}

	if post.Visibility != "public" && post.Visibility != "private" && post.Visibility != "followers" {
		return utils.ErrInvalidPayload
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return utils.ErrUUIDGeneration
	}
	post.PostID = uuid
	post.CreationDate = time.Now()
	return s.postRepo.InsertPost(ctx, post)
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
