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

type commentServiceImpl struct {
	commentRepo repositories.CommentRepository
	userRepo    repositories.UserRepository
	postService services.PostService
}

// NewCommentService crée une nouvelle instance de CommentService.
func NewCommentService(commentRepo repositories.CommentRepository, userRepo repositories.UserRepository, postService services.PostService) services.CommentService {
	return &commentServiceImpl{commentRepo: commentRepo, userRepo: userRepo, postService: postService}
}

func (s *commentServiceImpl) CreateComment(ctx context.Context, comment *models.Comment) error {
	// TODO: Vérifier que le post existe et que l'utilisateur a le droit de commenter
	if comment.PostID == uuid.Nil || comment.Content == nil || *comment.Content == "" || comment.UserID == uuid.Nil {
		return utils.ErrInvalidPayload
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return utils.ErrUUIDGeneration
	}
	if len(*comment.Content) > 500 {
		return utils.ErrContentTooLong
	}
	if len(*comment.Content) < 1 {
		return utils.ErrContentTooShort
	}

	comment.CommentID = uuid
	comment.CreationDate = time.Now()

	// TODO: Appeler le repository pour insérer le commentaire

	return s.commentRepo.InsertComment(ctx, comment)
}

func (s *commentServiceImpl) GetCommentsByPostID(ctx context.Context, postID string) ([]*models.CommentDetails, error) {
	// TODO: Vérifier l'accès au post (visibilité)
	if postID == "" {
		return nil, utils.ErrInvalidPayload
	}
	// Vérifier que le post existe
	exists, err := s.postService.IsPostExisting(ctx, postID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, utils.ErrPostNotFound
	}

	canSee, err := s.postService.CanUserSeePost(ctx, postID, &models.PostDetails{})
	if err != nil {
		return nil, err
	}
	if !canSee {
		return nil, utils.ErrUnauthorized
	}
	comments, err := s.commentRepo.GetCommentsByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}

	for _, comment := range comments {
		if comment.ImageURL != nil && *comment.ImageURL != "" {
			comment.ImageURL = utils.GetFullImageURL(comment.ImageURL)
		}
		if comment.AvatarURL != nil && *comment.AvatarURL != "" {
			comment.AvatarURL = utils.GetFullImageURLAvatar(comment.AvatarURL)
		}
	}
	return comments, nil
}

func (s *commentServiceImpl) DeleteComment(ctx context.Context, commentID, requesterID string) error {
	// TODO: Vérifier que le requester est l'auteur ou a les droits de modération
	// TODO: Appeler le repository pour supprimer le commentaire
	return nil
}

func (s *commentServiceImpl) UpdateComment(ctx context.Context, commentID, requesterID, content string) error {
	if commentID == "" || requesterID == "" {
		return utils.ErrInvalidPayload
	}
	if len(content) > 500 {
		return utils.ErrContentTooLong
	}
	if len(content) < 1 {
		return utils.ErrContentTooShort
	}

	comment, err := s.commentRepo.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}
	if comment == nil {
		return utils.ErrCommentNotFound
	}
	if comment.UserID.String() != requesterID {
		return utils.ErrForbidden
	}

	return s.commentRepo.UpdateComment(ctx, commentID, content)
}

func (s *commentServiceImpl) ReportComment(ctx context.Context, report *models.Report) error {
	// TODO: Vérifier que le commentaire existe
	// TODO: Appeler le repository pour enregistrer le signalement
	return nil
}
