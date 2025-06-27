package impl

import (
	"context"
	"database/sql"
	"mellow/models"
	"mellow/services"
)

type userServiceImpl struct {
	db *sql.DB
}

// NewUserService crée une nouvelle instance de UserService.
func NewUserService(db *sql.DB) services.UserService {
	return &userServiceImpl{db: db}
}

func (s *userServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	// TODO: Appliquer la logique métier (ex: validation, vérification unicité)
	// TODO: Appeler le repository pour insérer l'utilisateur en base
	return nil
}

func (s *userServiceImpl) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	// TODO: Vérifier la validité de l'ID si nécessaire
	// TODO: Appeler le repository pour récupérer l'utilisateur
	return nil, nil
}

func (s *userServiceImpl) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	// TODO: Appliquer les éventuelles règles métier (ex: autorisations)
	// TODO: Appeler le repository pour récupérer l'utilisateur par username
	return nil, nil
}

func (s *userServiceImpl) UpdateUser(ctx context.Context, user *models.User) error {
	// TODO: Appliquer la logique métier (ex: contrôle d'accès, validation)
	// TODO: Appeler le repository pour mettre à jour les données en base
	return nil
}

func (s *userServiceImpl) DeleteUser(ctx context.Context, userID string) error {
	// TODO: Vérifier les droits de suppression, effets de bord éventuels
	// TODO: Appeler le repository pour supprimer l'utilisateur
	return nil
}

func (s *userServiceImpl) Authenticate(ctx context.Context, username, password string) (*models.User, error) {
	// TODO: Vérifier les règles métier liées à l'authentification
	// TODO: Appeler le repository pour récupérer l'utilisateur et comparer le mot de passe
	return nil, nil
}

func (s *userServiceImpl) FollowUser(ctx context.Context, followerID, targetID string) error {
	// TODO: Vérifier la logique métier (ex: ne pas suivre soi-même, blocage, etc.)
	// TODO: Appeler le repository pour créer la relation de suivi
	return nil
}

func (s *userServiceImpl) UnfollowUser(ctx context.Context, followerID, targetID string) error {
	// TODO: Appliquer les règles de désabonnement (ex: vérifications)
	// TODO: Appeler le repository pour supprimer la relation de suivi
	return nil
}

func (s *userServiceImpl) GetFollowers(ctx context.Context, userID string) ([]*models.User, error) {
	// TODO: Appliquer les règles d'accès (ex: profil privé)
	// TODO: Appeler le repository pour récupérer les followers
	return nil, nil
}

func (s *userServiceImpl) GetFollowing(ctx context.Context, userID string) ([]*models.User, error) {
	// TODO: Appliquer les règles d'accès si nécessaire
	// TODO: Appeler le repository pour récupérer les utilisateurs suivis
	return nil, nil
}

func (s *userServiceImpl) SearchUsers(ctx context.Context, query string) ([]*models.User, error) {
	// TODO: Nettoyer/valider le terme de recherche
	// TODO: Appeler le repository pour exécuter la recherche
	return nil, nil
}
