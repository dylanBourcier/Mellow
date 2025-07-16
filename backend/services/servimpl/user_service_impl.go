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

type userServiceImpl struct {
	userRepo repositories.UserRepository // Référence au repository utilisateur
}

// NewUserService crée une nouvelle instance de UserService.
func NewUserService(userRepo repositories.UserRepository) services.UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (s *userServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	// Validation simple
	if user.Email == "" || user.Username == "" || user.Password == "" || user.Firstname == "" || user.Lastname == "" || user.Birthdate.IsZero() {
		return fmt.Errorf("%s: missing required fields", utils.ErrInvalidUserData)
	}
	// Vérifier si l'utilisateur existe déjà par email ou nom d'utilisateur
	exists, err := s.userRepo.UserExistsByEmailOrUsername(ctx, user.Email, user.Username)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}
	if exists {
		return utils.ErrUserAlreadyExists
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("failed to generate user ID: %w", err)
	}
	user.UserID = uuid

	user.CreationDate = time.Now() // Assigner la date de création actuelle

	if user.Role == "" {
		user.Role = "user" // Assigner un rôle par défaut si non spécifié
	}

	//Hash the password before storing it
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	return s.userRepo.InsertUser(ctx, user)
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
func (s *userServiceImpl) GetUserByUsernameOrEmail(ctx context.Context, identifier string) (*models.User, error) {

	var user *models.User
	user, err := s.userRepo.GetUserByUsernameOrEmail(ctx, identifier)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	if user == nil {
		return nil, utils.ErrUserNotFound
	}
	return user, nil
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

func (s *userServiceImpl) IsFollowing(ctx context.Context, followerID, targetID string) (bool, error) {
	return s.userRepo.IsFollowing(ctx, followerID, targetID)
}