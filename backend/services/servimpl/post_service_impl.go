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

func (s *postServiceImpl) GetPostByID(ctx context.Context, postID string, groupService services.GroupService, userService services.UserService, requesterID string) (*models.PostDetails, error) {
	// TODO: Vérifier la visibilité du post pour le requester
	// TODO: Récupérer le post depuis le repository
	post, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return nil, utils.ErrPostNotFound
	}
	if post == nil {
		return nil, utils.ErrPostNotFound
	}
	if post.GroupID != nil && post.GroupID.String() != "" && requesterID != "" {

		//verifier si l'user est membre du groupe
		isMember, err := groupService.IsMember(ctx, post.GroupID.String(), requesterID)
		if err != nil {
			return nil, utils.ErrInternalServerError
		}
		if !isMember {
			return nil, utils.ErrUnauthorized
		}
	}
	if post.ImageURL != nil {
		post.ImageURL = utils.GetFullImageURL(post.ImageURL)
	}
	if post.AvatarURL != nil {
		post.AvatarURL = utils.GetFullImageURLAvatar(post.AvatarURL)
	}

	switch post.Visibility {
	case "public":
		return post, nil
	case "followers":
		//Vérifier si l'user follow l'auteur
		isFollowing, err := userService.IsFollowing(ctx, requesterID, post.UserID.String())
		if err != nil {
			return nil, utils.ErrInternalServerError
		}
		if !isFollowing {
			return nil, utils.ErrUnauthorized
		}

		return post, nil
	case "private":
		//TODO: Vérifier si le user est autorisé à voir le post
	default:
		return nil, utils.ErrBadRequest

	}
	return post, nil
}

func (s *postServiceImpl) DeletePost(ctx context.Context, postID, requesterID string) error {
	// TODO: Vérifier que le requester est l’auteur ou un modérateur
	// TODO: Supprimer le post via le repository
	return nil
}

func (s *postServiceImpl) GetFeed(ctx context.Context, userID *string, limit, offset int) ([]*models.PostDetails, error) {
	// Validate limit and offset
	if limit <= 0 || offset < 0 {
		return nil, utils.ErrInvalidPayload
	}
	// Additional validation can be added here if needed
	posts, err := s.postRepo.GetFeed(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		if post.ImageURL != nil {
			post.ImageURL = utils.GetFullImageURL(post.ImageURL)
		}
		if post.AvatarURL != nil {
			post.AvatarURL = utils.GetFullImageURLAvatar(post.AvatarURL)
		}
	}
	return posts, nil
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
