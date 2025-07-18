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
}

// NewCommentService crée une nouvelle instance de CommentService.
func NewCommentService(commentRepo repositories.CommentRepository) services.CommentService {
	return &commentServiceImpl{commentRepo: commentRepo}
}

func (s *commentServiceImpl) CreateComment(ctx context.Context, comment *models.Comment) error {
	// TODO: Vérifier que le post existe et que l'utilisateur a le droit de commenter
	if comment.PostID == uuid.Nil || comment.Content == nil || comment.UserID == uuid.Nil {
		return utils.ErrInvalidPayload
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return utils.ErrUUIDGeneration
	}
	if comment.Content == nil || *comment.Content == "" {
		return utils.ErrInvalidPayload
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

func (s *commentServiceImpl) GetCommentsByPostID(ctx context.Context, postID string) ([]*models.Comment, error) {
	// TODO: Vérifier l'accès au post (visibilité)
	// TODO: Appeler le repository pour récupérer les commentaires liés au post
	return nil, nil
}

func (s *commentServiceImpl) DeleteComment(ctx context.Context, commentID, requesterID string) error {
	// TODO: Vérifier que le requester est l'auteur ou a les droits de modération
	// TODO: Appeler le repository pour supprimer le commentaire
	return nil
}

func (s *commentServiceImpl) ReportComment(ctx context.Context, report *models.Report) error {
	// TODO: Vérifier que le commentaire existe
	// TODO: Appeler le repository pour enregistrer le signalement
	return nil
}
