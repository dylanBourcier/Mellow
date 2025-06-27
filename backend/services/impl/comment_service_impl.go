package impl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/services"
)

type commentServiceImpl struct {
	db *sql.DB
}

// NewCommentService crée une nouvelle instance de CommentService.
func NewCommentService(db *sql.DB) services.CommentService {
	return &commentServiceImpl{db: db}
}

func (s *commentServiceImpl) CreateComment(ctx context.Context, comment *models.Comment) error {
	// TODO: Vérifier que le post existe et que l'utilisateur a le droit de commenter
	// TODO: Appeler le repository pour insérer le commentaire
	return nil
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
