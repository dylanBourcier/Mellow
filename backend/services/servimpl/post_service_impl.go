package servimpl

import (
	"context"
	"fmt"
	"mellow/models"
	"mellow/repositories"
	"mellow/services"
	"mellow/utils"
	"time"

	"github.com/google/uuid"
)

type postServiceImpl struct {
	postRepo     repositories.PostRepository
	userService  services.UserService
	groupService services.GroupService
}

// NewPostService crée une nouvelle instance de PostService.
func NewPostService(postRepo repositories.PostRepository, userService services.UserService, groupService services.GroupService) services.PostService {
	return &postServiceImpl{postRepo: postRepo, userService: userService, groupService: groupService}
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

func (s *postServiceImpl) GetPostByID(ctx context.Context, postID string, requesterID string) (*models.PostDetails, error) {
	// Récupérer le post depuis le repository
	post, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return nil, utils.ErrPostNotFound
	}
	if post == nil {
		return nil, utils.ErrPostNotFound
	}

	// Vérifier la visibilité du post pour le requester
	canSee, err := s.CanUserSeePost(ctx, postID, post)
	if err != nil {
		return nil, utils.ErrInternalServerError
	}
	if !canSee {
		return nil, utils.ErrUnauthorized
	}

	if post.ImageURL != nil {
		post.ImageURL = utils.GetFullImageURL(post.ImageURL)
	}
	if post.AvatarURL != nil {
		post.AvatarURL = utils.GetFullImageURLAvatar(post.AvatarURL)
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

func (s *postServiceImpl) IsPostExisting(ctx context.Context, postID string) (bool, error) {
	exists, err := s.postRepo.IsPostExisting(ctx, postID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// postService.go

func (s *postServiceImpl) CanUserSeePost(ctx context.Context, postId string, postDetails *models.PostDetails) (bool, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil && err != utils.ErrNoUserInContext {
		return false, err
	}
	if postDetails.PostID == uuid.Nil {
		//If post is nil, get it from the repository
		postDetails, err = s.postRepo.GetPostByID(ctx, postId)
		if err != nil {
			fmt.Println("2", err)

			return false, err
		}
		if postDetails == nil {
			fmt.Println("3", err)
			return false, utils.ErrPostNotFound
		}
	}

	if postDetails.GroupID != nil && postDetails.GroupID.String() != "" {
		if userID.String() == "" {
			fmt.Println("4", err)
			return false, nil // User is not logged in, cannot see the post
		}
		// Vérifier si l'utilisateur est membre du groupe
		isMember, err := s.groupService.IsMember(ctx, postDetails.GroupID.String(), userID.String())
		if err != nil {
			fmt.Println("5", err)
			return false, err
		}
		if !isMember {
			fmt.Println("6", err)
			return false, nil // User is not a member of the group, cannot see the post
		}
		return true, nil // User is a member of the group, can see the post
	}
	fmt.Println("post", postDetails)
	switch postDetails.Visibility {
	case "public":
		return true, nil
	case "followers":
		if userID.String() == "" {
			fmt.Println("7", err)
			return false, nil
		}
		if userID == postDetails.UserID {
			return true, nil
		}
		return s.userService.IsFollowing(ctx, userID.String(), postDetails.UserID.String())
	case "private":
		if userID.String() == "" {
			fmt.Println("8", err)
			return false, nil
		}
		if userID == postDetails.UserID {
			return true, nil
		}
		return s.postRepo.IsUserAllowed(ctx, postDetails.PostID.String(), userID.String())
	default:
		fmt.Println("9", err)
		return false, nil
	}
}
