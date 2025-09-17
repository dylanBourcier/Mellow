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
	postRepo     repositories.PostRepository
	userService  services.UserService
	groupService services.GroupService
}

// NewPostService crée une nouvelle instance de PostService.
func NewPostService(postRepo repositories.PostRepository, userService services.UserService, groupService services.GroupService) services.PostService {
	return &postServiceImpl{postRepo: postRepo, userService: userService, groupService: groupService}
}
func (s *postServiceImpl) CreatePost(ctx context.Context, post *models.Post) error {
	// Validate the post (content, visibility)
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

	// Insert the post into the repository
	err = s.postRepo.InsertPost(ctx, post)
	if err != nil {
		return err
	}

	// If the visibility is "private", insert viewers into the post_viewer table
	if post.Visibility == "private" {
		for _, viewerID := range post.Viewers {
			err := s.postRepo.AddPostViewer(ctx, post.PostID.String(), viewerID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *postServiceImpl) UpdatePost(ctx context.Context, postID, requesterID, title, content string) error {
	if postID == "" || requesterID == "" || title == "" || content == "" {
		return utils.ErrInvalidPayload
	}

	existing, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return utils.ErrPostNotFound
	}
	if existing == nil {
		return utils.ErrPostNotFound
	}
	if existing.UserID.String() != requesterID {
		return utils.ErrForbidden
	}

	if len(content) > 5000 {
		return utils.ErrContentTooLong
	}
	if len(content) < 1 {
		return utils.ErrContentTooShort
	}

	pid, err := uuid.Parse(postID)
	if err != nil {
		return utils.ErrInvalidPayload
	}

	post := &models.Post{
		PostID:  pid,
		Title:   title,
		Content: content,
	}

	return s.postRepo.UpdatePost(ctx, post)
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
	if postID == "" || requesterID == "" {
		return utils.ErrInvalidPayload
	}

	post, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return err
	}
	if post == nil {
		return utils.ErrPostNotFound
	}

	if post.UserID.String() != requesterID {
		return utils.ErrUnauthorized
	}

	if err := s.postRepo.DeletePost(ctx, postID); err != nil {
		return err
	}
	return nil
}

func (s *postServiceImpl) GetFeed(ctx context.Context, userID string, limit, offset int) ([]*models.PostDetails, error) {
	// Validate limit and offset
	if limit <= 0 || offset < 0 {
		return nil, utils.ErrInvalidPayload
	}
	// Additional validation can be added here if needed
	posts, err := s.postRepo.GetFeed(ctx, userID, "", limit, offset)
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

func (s *postServiceImpl) GetGroupPosts(ctx context.Context, groupID string, limit, offset int) ([]*models.PostDetails, error) {
	if groupID == "" || limit <= 0 || offset < 0 {
		return nil, utils.ErrInvalidPayload
	}
	posts, err := s.postRepo.GetGroupPosts(ctx, groupID, limit, offset)
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

func (s *postServiceImpl) GetUserPosts(ctx context.Context, userID, requesterID string, limit, offset int) ([]*models.PostDetails, error) {
	// Validate limit and offset
	if limit <= 0 || offset < 0 {
		return nil, utils.ErrInvalidPayload
	}
	// Additional validation can be added here if needed
	posts, err := s.postRepo.GetFeed(ctx, userID, requesterID, limit, offset)
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
			return false, err
		}
		if postDetails == nil {
			return false, utils.ErrPostNotFound
		}
	}

	if postDetails.GroupID != nil && postDetails.GroupID.String() != "" {
		if userID.String() == "" {
			return false, nil // User is not logged in, cannot see the post
		}
		// Vérifier si l'utilisateur est membre du groupe
		isMember, err := s.groupService.IsMember(ctx, postDetails.GroupID.String(), userID.String())
		if err != nil {
			return false, err
		}
		if !isMember {
			return false, nil // User is not a member of the group, cannot see the post
		}
		return true, nil // User is a member of the group, can see the post
	}
	switch postDetails.Visibility {
	case "public":
		return true, nil
	case "followers":
		if userID.String() == "" {
			return false, nil
		}
		if userID == postDetails.UserID {
			return true, nil
		}
		return s.userService.IsFollowing(ctx, userID.String(), postDetails.UserID.String())
	case "private":
		if userID.String() == "" {
			return false, nil
		}
		if userID == postDetails.UserID {
			return true, nil
		}
		return s.postRepo.IsUserAllowed(ctx, postDetails.PostID.String(), userID.String())
	default:
		return false, nil
	}
}
